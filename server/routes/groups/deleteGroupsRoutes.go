package groups

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"server/internal/http_server/middlewares"
	communicationJson "server/internal/http_server/network/communication/json"
	"server/internal/lib/logger/sl"
	"server/internal/storage/postgres"
)

func DeleteGroup(logger *slog.Logger, storage *postgres.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		teacherId, err := middlewares.GetTeacherIdFromContext(r.Context())
		if err != nil {
			http.Error(w, "Can't get teacher id", http.StatusInternalServerError)
			logger.Error("Can't get teacher id", sl.Err(err))
			return
		}

		var groupIdJson communicationJson.DeleteGroupJson
		if err := json.NewDecoder(r.Body).Decode(&groupIdJson); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			logger.Error("Can't unmarshal JSON", sl.Err(err))
			return
		}

		if !storage.IsTeacherInGroup(groupIdJson.GroupId, teacherId) {
			http.Error(w, "Teacher is not owner of this group", http.StatusForbidden)
			return
		}

		if err := storage.DeleteGroup(groupIdJson.GroupId); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func DeleteStudentFromGroup(logger *slog.Logger, storage *postgres.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		teacherId, err := middlewares.GetTeacherIdFromContext(r.Context())
		if err != nil {
			http.Error(w, "Can't get teacher id", http.StatusInternalServerError)
			logger.Error("Can't get teacher id", sl.Err(err))
			return
		}

		var groupData communicationJson.DeleteStudentFromGroupJson
		if err := json.NewDecoder(r.Body).Decode(&groupData); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			logger.Error("Can't unmarshal JSON", sl.Err(err))
			return
		}

		if !storage.IsTeacherInGroup(groupData.GroupId, teacherId) {
			http.Error(w, "Teacher is not owner of this group", http.StatusForbidden)
			return
		}

		if err := storage.DeleteStudentFromGroup(groupData.GroupId, groupData.StudentId); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
