package permissions

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
	"log/slog"
	"net/http"
	"server/internal/http_server/middlewares"
	"server/internal/lib/logger/sl"
	"server/internal/storage/postgres"
	"server/models"
	"strconv"
)

func GetGroup(logger *slog.Logger, storage *postgres.Storage, w http.ResponseWriter, r *http.Request) (*models.Group, error) {
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

func ValidateTeacherGroup(logger *slog.Logger, storage *postgres.Storage, w http.ResponseWriter, r *http.Request) (*models.Group, error) {
	group, err := GetGroup(logger, storage, w, r)
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

func ValidatePermissionsInGroup(w http.ResponseWriter, r *http.Request, logger *slog.Logger, storage *postgres.Storage) (*models.Group, bool) {
	group, err := GetGroup(logger, storage, w, r)
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

func ValidatePermissionsInLesson(w http.ResponseWriter, r *http.Request, logger *slog.Logger, storage *postgres.Storage) (*models.Lesson, bool) {
	group, forbidden := ValidatePermissionsInGroup(w, r, logger, storage)
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

func ValidateTeachersPermissionInLesson(w http.ResponseWriter, r *http.Request, logger *slog.Logger, storage *postgres.Storage) (*models.Lesson, bool) {
	lesson, ok := ValidatePermissionsInLesson(w, r, logger, storage)
	if ok {
		return nil, true
	}

	teacherId, err := middlewares.GetTeacherIdFromContext(r.Context())
	if err != nil {
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		logger.Error("Can't get teacher_id from context", sl.Err(err))
		return nil, true
	}

	if lesson.Group.TeacherID != teacherId {
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return nil, true
	}
	return lesson, false
}

func GetHomeworkIdFromRequest(r *http.Request) (uint, error) {
	homeworkId, err := strconv.Atoi(chi.URLParam(r, "homework_id"))
	return uint(homeworkId), err
}

func GetSolutionIdFromRequest(r *http.Request) (uint, error) {
	solutionId, err := strconv.Atoi(chi.URLParam(r, "solution_id"))
	return uint(solutionId), err
}

func GetStudentIdFromContext(r *http.Request) uint {
	_, claims, _ := jwtauth.FromContext(r.Context())
	return uint(claims["id"].(float64))
}

func GetStudentIdFromRequest(r *http.Request) (uint, error) {
	studentId, err := strconv.Atoi(chi.URLParam(r, "student_id"))
	return uint(studentId), err
}

func GetSolveIdFromRequest(r *http.Request) (uint, error) {
	solveId, err := strconv.Atoi(chi.URLParam(r, "solve_id"))
	return uint(solveId), err
}

func GetLessonIdFromRequest(r *http.Request) (uint, error) {
	lessonId, err := strconv.Atoi(chi.URLParam(r, "lesson_id"))
	return uint(lessonId), err
}

func GetGroupIdFromRequest(r *http.Request) (uint, error) {
	groupId, err := strconv.Atoi(chi.URLParam(r, "group_id"))
	return uint(groupId), err
}
