package homeworks

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"server/internal/http_server/middlewares"
	communicationJson "server/internal/http_server/network/communication/json"
	"server/internal/lib/logger/sl"
	"server/internal/storage/postgres"
)

func DeleteHomework(logger *slog.Logger, storage *postgres.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var homeworkData communicationJson.DeleteHomeworkJson
		if err := json.NewDecoder(r.Body).Decode(&homeworkData); err != nil {
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

		homework, err := storage.GetHomeworkById(homeworkData.HomeworkId)
		if err != nil {
			http.Error(w, "Can't get homework", http.StatusInternalServerError)
			logger.Error("Can't get homework", sl.Err(err))
			return
		}

		if !storage.IsTeacherInGroup(homework.Lesson.GroupID, teacherId) {
			http.Error(w, "Teacher is not owner of this group", http.StatusForbidden)
			return
		}

		err = storage.DeleteHomework(homeworkData.HomeworkId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func DeleteHomeworkSolve(logger *slog.Logger, storage *postgres.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var homeworkData communicationJson.DeleteHomeworkSolveJson
		if err := json.NewDecoder(r.Body).Decode(&homeworkData); err != nil {
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

		homeworkSolve, err := storage.GetHomeworkSolveById(homeworkData.HomeworkSolveId)
		if err != nil {
			http.Error(w, "Can't get homework solve", http.StatusInternalServerError)
			logger.Error("Can't get homework solve", sl.Err(err))
			return
		}

		if !storage.IsTeacherInSolve(teacherId, homeworkSolve.ID) {
			http.Error(w, "Teacher is not owner of this solve", http.StatusForbidden)
			return
		}

		err = storage.DeleteHomeworkSolve(homeworkData.HomeworkSolveId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
