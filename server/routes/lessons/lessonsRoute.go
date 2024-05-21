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
		studentId := permissions.GetStudentIdFromContext(r)
		lessons, err := storage.GetLessonsByStudentId(studentId)

		if err != nil {
			http.Error(w, "Can't get lessons", http.StatusInternalServerError)
			logger.Error("Can't get lessons", sl.Err(err))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(map[string]any{"lessons": lessons}); err != nil {
			logger.Error("Can't marshall lessons json", sl.Err(err))
		}
	}
}

func GetLessonById(logger *slog.Logger, storage *postgres.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		studentId := permissions.GetStudentIdFromContext(r)
		lessonId, err := permissions.GetLessonIdFromRequest(r)

		if err != nil {
			http.Error(w, "Can't get lessonId", http.StatusBadRequest)
			logger.Error("Can't get lessonId", sl.Err(err))
			return
		}

		lesson, err := storage.GetLessonById(lessonId)
		if err != nil {
			http.Error(w, "Can't get lesson", http.StatusInternalServerError)
			logger.Error("Can't get lesson", sl.Err(err))
			return
		}

		if isStudentInLesson, err := storage.IsStudentInLesson(lesson.ID, studentId); err != nil {
			http.Error(w, "Can't get lesson", http.StatusInternalServerError)
			logger.Error("Can't get lesson", sl.Err(err))
			return
		} else if !isStudentInLesson {
			http.Error(w, "Student is not in lesson", http.StatusForbidden)
			logger.Error("Student is not in lesson", sl.Err(err))
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

func GetHomeworksByLessonId(logger *slog.Logger, storage *postgres.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lessonId, err := permissions.GetLessonIdFromRequest(r)
		if err != nil {
			http.Error(w, "Can't get lessonId", http.StatusBadRequest)
			logger.Error("Can't get lessonId", sl.Err(err))
			return
		}

		teacherId, err := middlewares.GetTeacherIdFromContext(r.Context())
		if err != nil {
			http.Error(w, "Can't get teacherId", http.StatusNotFound)
			logger.Error("Can't get teacherId", sl.Err(err))
			return
		}

		lesson, err := storage.GetLessonById(lessonId)
		if err != nil {
			http.Error(w, "Can't get lesson", http.StatusInternalServerError)
			logger.Error("Can't get lesson", sl.Err(err))
			return
		}

		if lesson.Group.TeacherID != teacherId {
			http.Error(w, "Teacher is not owner of this group", http.StatusForbidden)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(map[string]any{"homeworks": lesson.Homeworks}); err != nil {
			logger.Error("Can't marshall homeworks json", sl.Err(err))
		}
	}
}

func GetLessonsByTeacher(logger *slog.Logger, storage *postgres.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		teacherId, err := middlewares.GetTeacherIdFromContext(r.Context())
		if err != nil {
			http.Error(w, "Can't get teacherId", http.StatusNotFound)
			logger.Error("Can't get teacherId", sl.Err(err))
			return
		}

		lessons, err := storage.GetLessonsByTeacher(teacherId)
		if err != nil {
			http.Error(w, "Can't get lessons", http.StatusInternalServerError)
			logger.Error("Can't get lessons", sl.Err(err))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(map[string]any{"lessons": lessons}); err != nil {
			logger.Error("Can't marshall lessons json", sl.Err(err))
		}
	}
}

func GetLessonsByGroup(logger *slog.Logger, storage *postgres.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		groupId, err := permissions.GetGroupIdFromRequest(r)
		if err != nil {
			http.Error(w, "Can't get groupId", http.StatusBadRequest)
			logger.Error("Can't get groupId", sl.Err(err))
			return
		}

		teacherId, err := middlewares.GetTeacherIdFromContext(r.Context())
		if err != nil {
			http.Error(w, "Can't get teacherId", http.StatusNotFound)
			logger.Error("Can't get teacherId", sl.Err(err))
			return
		}

		if !storage.IsTeacherInGroup(groupId, teacherId) {
			http.Error(w, "Teacher is not owner of this group", http.StatusForbidden)
			return
		}

		group, err := storage.GetGroupById(groupId)
		if err != nil {
			http.Error(w, "Can't get group", http.StatusInternalServerError)
			logger.Error("Can't get group", sl.Err(err))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(map[string]any{"lessons": group.Lessons}); err != nil {
			logger.Error("Can't marshall lessons json", sl.Err(err))
		}
	}
}

func DeleteLesson(logger *slog.Logger, storage *postgres.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var lessonData communicationJson.DeleteLessonJson
		if err := json.NewDecoder(r.Body).Decode(&lessonData); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			logger.Error("Can't unmarshal JSON", sl.Err(err))
			return
		}

		teacherId, err := middlewares.GetTeacherIdFromContext(r.Context())
		if err != nil {
			http.Error(w, "Can't get teacherId", http.StatusNotFound)
			logger.Error("Can't get teacherId", sl.Err(err))
			return
		}

		lesson, err := storage.GetLessonById(lessonData.LessonID)
		if err != nil {
			http.Error(w, "Can't get lesson", http.StatusInternalServerError)
			logger.Error("Can't get lesson", sl.Err(err))
			return
		}

		if lesson.Group.TeacherID != teacherId {
			http.Error(w, "Teacher is not owner of this group", http.StatusForbidden)
			return
		}

		if err := storage.DeleteLesson(lessonData.LessonID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
