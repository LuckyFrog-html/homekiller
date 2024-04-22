package lessons

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"server/internal/http_server/middlewares"
	communicationJson "server/internal/http_server/network/communication/json"
	"server/internal/http_server/permissions"
	"server/internal/lib/logger/sl"
	"server/internal/storage/postgres"
	"server/models"
)

func AddLesson(logger *slog.Logger, storage *postgres.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var lessonData communicationJson.AddLessonJson
		if err := json.NewDecoder(r.Body).Decode(&lessonData); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			logger.Error("Can't unmarshal JSON", sl.Err(err))
			return
		}

		teacherId, err := middlewares.GetTeacherIdFromContext(r.Context())
		if err != nil {
			http.Error(w, "Can't get teacherId", http.StatusNotFound)
			logger.Error("Can't get teacherId", sl.Err(err))
			return

		}

		if !storage.IsTeacherInGroup(lessonData.GroupId, teacherId) {
			http.Error(w, "Teacher is not owner of this group", http.StatusForbidden)
			return
		}

		lesson, err := storage.AddLesson(lessonData.Date, lessonData.GroupId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err = json.NewEncoder(w).Encode(lesson); err != nil {
			logger.Error("Can't marshall lesson json", sl.Err(err))
		}
	}
}

func GetLessons(logger *slog.Logger, storage *postgres.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		group, forbidden := permissions.ValidatePermissionsInGroup(w, r, logger, storage)
		if forbidden {
			return
		}

		lessons := group.Lessons
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(lessons); err != nil {
			logger.Error("Can't marshall lessons json", sl.Err(err))
		}
	}
}

func GetLessonByGroup(logger *slog.Logger, storage *postgres.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lesson, done := permissions.ValidatePermissionsInLesson(w, r, logger, storage)
		if done {
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(lesson); err != nil {
			logger.Error("Can't marshall lesson json", sl.Err(err))
		}
	}
}

func MarkStudentAttendance(logger *slog.Logger, storage *postgres.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var attendanceData communicationJson.MarkStudentAttendanceJson

		if err := json.NewDecoder(r.Body).Decode(&attendanceData); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			logger.Error("Can't unmarshal JSON", sl.Err(err))
			return
		}

		teacherId, err := middlewares.GetTeacherIdFromContext(r.Context())
		if err != nil {
			http.Error(w, "Can't get teacherId", http.StatusNotFound)
			logger.Error("Can't get teacherId", sl.Err(err))
			return
		}

		lesson, err := storage.GetLessonById(attendanceData.LessonID)
		if lesson.Group.TeacherID != teacherId {
			http.Error(w, "Teacher is not owner of this group", http.StatusForbidden)
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
		tx := storage.Db.Begin()
		tx.Save(&lesson)
		tx.Commit()
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(lesson); err != nil {
			logger.Error("Can't marshall lesson json", sl.Err(err))
		}
	}
}
