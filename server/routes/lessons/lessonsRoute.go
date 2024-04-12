package lessons

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
	"log/slog"
	"net/http"
	"server/internal/http_server/middlewares"
	communicationJson "server/internal/http_server/network/communication/json"
	"server/internal/lib/logger/sl"
	"server/internal/storage/postgres"
	"server/models"
	"strconv"
)

func getGroup(logger *slog.Logger, storage *postgres.Storage, w http.ResponseWriter, r *http.Request) (*models.Group, error) {
	groupId, err := strconv.Atoi(chi.URLParam(r, "group_id"))
	if err != nil {
		http.Error(w, "You must send groupId as URL part like /groups/{group_id}/lessons", http.StatusBadRequest)
		logger.Error("Can't parse groupId", sl.Err(err))
		return nil, err
	}

	group, err := storage.GetGroupById(uint(groupId))
	if err != nil {
		http.Error(w, "Group not found", http.StatusNotFound)
		logger.Error("Group not found", sl.Err(err))
		return nil, err
	}
	return &group, nil
}

func validateTeacherGroup(logger *slog.Logger, storage *postgres.Storage, w http.ResponseWriter, r *http.Request) (*models.Group, error) {
	group, err := getGroup(logger, storage, w, r)
	if err != nil {
		return nil, err
	}

	teacherId, err := middlewares.GetTeacherIdFromContext(r.Context())

	if err != nil {
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		logger.Error("Can't get teacher_id from context", sl.Err(err))
		return nil, err
	}

	if group.TeacherID != teacherId {
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		logger.Error("Forbidden", sl.Err(err))
		return nil, err
	}
	return group, nil
}

func validatePermissionsInGroup(w http.ResponseWriter, r *http.Request, logger *slog.Logger, storage *postgres.Storage) (*models.Group, bool) {
	group, err := getGroup(logger, storage, w, r)
	if err != nil {
		return nil, true
	}

	_, claims, _ := jwtauth.FromContext(r.Context())
	userId := uint(claims["id"].(float64))
	if claims["table"] == "teacher" && group.TeacherID != userId || claims["table"] == "student" && !storage.IsStudentInGroup(userId, group.ID) {
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		logger.Error("Forbidden", sl.Err(err))
		return nil, true
	}
	return group, false
}

func validatePermissionsInLesson(w http.ResponseWriter, r *http.Request, logger *slog.Logger, storage *postgres.Storage) (*models.Lesson, bool) {
	group, forbidden := validatePermissionsInGroup(w, r, logger, storage)
	if forbidden {
		return nil, true
	}

	lessonId, err := strconv.Atoi(chi.URLParam(r, "lesson_id"))
	if err != nil {
		http.Error(w, "You must send groupId as URL part like /groups/{group_id}/lessons/{lesson_id}", http.StatusBadRequest)
		logger.Error("Can't parse groupId", sl.Err(err))
		return nil, true
	}

	lesson, err := storage.GetLessonById(uint(lessonId))
	if err != nil {
		http.Error(w, "Lesson not found", http.StatusNotFound)
		return nil, true
	}

	if lesson.GroupID != group.ID {
		http.Error(w, "Lesson not found", http.StatusNotFound)
		return nil, true
	}
	return lesson, false
}

func AddLesson(logger *slog.Logger, storage *postgres.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		group, err := validateTeacherGroup(logger, storage, w, r)
		if err != nil {
			return
		}

		var lessonData communicationJson.AddLessonJson
		if err := json.NewDecoder(r.Body).Decode(&lessonData); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			logger.Error("Can't unmarshal JSON", sl.Err(err))
			return
		}

		lesson := storage.AddLesson(lessonData.Date, group.ID)
		if err = json.NewEncoder(w).Encode(lesson); err != nil {
			logger.Error("Can't marshall lesson json", sl.Err(err))
		}
	}
}

func GetLessons(logger *slog.Logger, storage *postgres.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		group, forbidden := validatePermissionsInGroup(w, r, logger, storage)
		if forbidden {
			return
		}

		lessons := group.Lessons
		if err := json.NewEncoder(w).Encode(lessons); err != nil {
			logger.Error("Can't marshall lessons json", sl.Err(err))
		}
	}
}

func GetLessonByGroup(logger *slog.Logger, storage *postgres.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lesson, done := validatePermissionsInLesson(w, r, logger, storage)
		if done {
			return
		}

		if err := json.NewEncoder(w).Encode(lesson); err != nil {
			logger.Error("Can't marshall lesson json", sl.Err(err))
		}
	}
}

func MarkStudentAttendance(logger *slog.Logger, storage *postgres.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lesson, done := validatePermissionsInLesson(w, r, logger, storage)
		if done {
			return
		}

		var attendanceData communicationJson.MarkStudentAttendanceJson
		if err := json.NewDecoder(r.Body).Decode(&attendanceData); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			logger.Error("Can't unmarshal JSON", sl.Err(err))
			return
		}

		students := make([]*models.Student, 0, len(attendanceData.StudentIDs))
		for _, studentId := range attendanceData.StudentIDs {
			student, err := storage.GetStudentByID(studentId)
			if err != nil {
				http.Error(w, fmt.Sprintf("Student not found with id=%d is not found", studentId),
					http.StatusBadRequest)
				return
			}
			students = append(students, student)
		}
		lesson.Students = append(lesson.Students, students...)
		storage.Db.Commit()
		if err := json.NewEncoder(w).Encode(lesson); err != nil {
			logger.Error("Can't marshall lesson json", sl.Err(err))
		}
	}
}
