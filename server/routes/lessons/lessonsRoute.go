package lessons

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	communicationJson "server/internal/http_server/network/communication/json"
	"server/internal/http_server/permissions"
	"server/internal/lib/logger/sl"
	"server/internal/storage/postgres"
	"server/models"
)

func AddLesson(logger *slog.Logger, storage *postgres.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		group, err := permissions.ValidateTeacherGroup(logger, storage, w, r)
		if err != nil {
			return
		}

		var lessonData communicationJson.AddLessonJson
		if err := json.NewDecoder(r.Body).Decode(&lessonData); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			logger.Error("Can't unmarshal JSON", sl.Err(err))
			return
		}

		lesson, err := storage.AddLesson(lessonData.Date, group.ID)
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
		lesson, done := permissions.ValidatePermissionsInLesson(w, r, logger, storage)
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
		tx := storage.Db.Begin()
		tx.Save(&lesson)
		tx.Commit()
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(lesson); err != nil {
			logger.Error("Can't marshall lesson json", sl.Err(err))
		}
	}
}
