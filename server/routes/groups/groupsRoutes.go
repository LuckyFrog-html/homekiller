package groups

import (
	"encoding/json"
	"github.com/go-chi/jwtauth"
	"log/slog"
	"net/http"
	"server/internal/http_server/middlewares"
	communicationJson "server/internal/http_server/network/communication/json"
	"server/internal/http_server/permissions"
	"server/internal/lib/logger/sl"
	"server/internal/storage/postgres"
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
		teacherId, err := middlewares.GetTeacherIdFromContext(r.Context())
		if err != nil {
			http.Error(w, "Can't get teacher id", http.StatusInternalServerError)
			logger.Error("Can't get teacher id", sl.Err(err))
			return
		}

		groupId, err := permissions.GetGroupIdFromRequest(r)

		var groupData communicationJson.AddStudentToGroupJson

		if err := json.NewDecoder(r.Body).Decode(&groupData); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			logger.Error("Can't unmarshal JSON", sl.Err(err))
			return
		}

		if !storage.IsTeacherInGroup(groupId, teacherId) {
			http.Error(w, "Teacher is not owner of this group", http.StatusForbidden)
			return
		}

		err = storage.AddStudentsToGroup(groupId, groupData.StudentsIds)
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
		studentId := permissions.GetStudentIdFromContext(r)

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

func GetGroupsByTeacher(logger *slog.Logger, storage *postgres.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		teacherId, err := middlewares.GetTeacherIdFromContext(r.Context())
		if err != nil {
			http.Error(w, "Can't get teacher id", http.StatusInternalServerError)
			logger.Error("Can't get teacher id", sl.Err(err))
			return
		}

		groups, err := storage.GetGroupsByTeacher(teacherId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(map[string]interface{}{"groups": groups}); err != nil {
			logger.Error("Can't marshall groups json", sl.Err(err))
		}
	}
}
