package routes

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"server/internal/http_server/middlewares"
	communicationJson "server/internal/http_server/network/communication/json"
	"server/internal/lib/logger/sl"
	"server/internal/storage/postgres"
)

func AddStudentHandler(logger *slog.Logger, storage *postgres.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var studentData communicationJson.AddStudentJson

		if err := json.NewDecoder(r.Body).Decode(&studentData); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			logger.Error("Can't unmarshal JSON", sl.Err(err))
			return
		}

		student := storage.AddStudent(studentData.Name, studentData.Stage, studentData.Login, studentData.Password)

		if err := json.NewEncoder(w).Encode(student); err != nil {
			logger.Error("Can't marshall student json", sl.Err(err))
		}
	}
}

func LoginStudentHandler(logger *slog.Logger, storage *postgres.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var studentData communicationJson.GetStudentJson

		if err := json.NewDecoder(r.Body).Decode(&studentData); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			logger.Error("Can't unmarshal JSON", sl.Err(err))
		}

		student, err := storage.GetStudentByLogin(studentData.Login, studentData.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		jwtAuthToken, err := middlewares.GetAuthTokenFromContext(r.Context())
		if err != nil {
			logger.Error("Can't get jwt auth token", sl.Err(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		claims := map[string]interface{}{"id": student.ID, "login": student.Login, "table": "students"}
		_, token, err := jwtAuthToken.Encode(claims)
		if err != nil {
			logger.Error("Can't get jwt", sl.Err(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(map[string]interface{}{"token": token})
	}
}
