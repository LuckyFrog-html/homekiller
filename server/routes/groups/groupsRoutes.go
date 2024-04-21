package groups

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
	"log/slog"
	"net/http"
	communicationJson "server/internal/http_server/network/communication/json"
	"server/internal/http_server/permissions"
	"server/internal/lib/logger/sl"
	"server/internal/storage/postgres"
	"strconv"
)

func AddGroup(logger *slog.Logger, storage *postgres.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, claims, _ := jwtauth.FromContext(r.Context())

		var groupData communicationJson.AddGroupJson

		if err := json.NewDecoder(r.Body).Decode(&groupData); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			logger.Error("Can't unmarshal JSON", sl.Err(err))
			return
		}

		group, err := storage.AddGroup(groupData.Title, uint(claims["id"].(float64)))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(group); err != nil {
			logger.Error("Can't marshall group json", sl.Err(err))
		}
	}
}

func AddStudentsToGroup(logger *slog.Logger, storage *postgres.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, claims, err := jwtauth.FromContext(r.Context())

		var groupData communicationJson.AddStudentToGroupJson

		groupId, err := strconv.Atoi(chi.URLParam(r, "group_id"))

		if err != nil {
			http.Error(w, "You must send groupId as URL part like /groups/{group_id}/students", http.StatusBadRequest)
			logger.Error("Can't parse groupId", sl.Err(err))
			return
		}

		group, err := storage.GetGroupById(uint(groupId))
		if err != nil {
			http.Error(w, "Group not found", http.StatusNotFound)
			return
		}

		if group.TeacherID != uint(claims["id"].(float64)) {
			http.Error(w, "You are not the owner of this group", http.StatusForbidden)
			return
		}

		if err := json.NewDecoder(r.Body).Decode(&groupData); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			logger.Error("Can't unmarshal JSON", sl.Err(err))
			return
		}

		err = storage.AddStudentsToGroup(uint(groupId), groupData.StudentsIds)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(map[string]interface{}{"status": "added"}); err != nil {
			logger.Error("Can't marshall student json", sl.Err(err))
		}
	}
}

func GetStudentsFromGroup(logger *slog.Logger, storage *postgres.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		group, done := permissions.ValidatePermissionsInGroup(w, r, logger, storage)
		if done {
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(map[string]interface{}{"students": group.Students}); err != nil {
			logger.Error("Can't marshall students json", sl.Err(err))
		}
	}
}

func GetGroupsByStudentHandler(logger *slog.Logger, storage *postgres.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, claims, _ := jwtauth.FromContext(r.Context())
		studentId := uint(claims["id"].(float64))

		groups, err := storage.GetGroupsByStudent(studentId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(groups); err != nil {
			logger.Error("Can't marshall groups json", sl.Err(err))
		}
	}
}
