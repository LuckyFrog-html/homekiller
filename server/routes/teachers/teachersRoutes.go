package teachers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"server/internal/http_server/middlewares"
	communicationJson "server/internal/http_server/network/communication/json"
	"server/internal/lib/logger/sl"
	"server/internal/storage/postgres"
)

func LoginTeacher(logger *slog.Logger, storage *postgres.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var teacherData communicationJson.GetTeacherJson

		if err := json.NewDecoder(r.Body).Decode(&teacherData); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			logger.Error("Can't unmarshal JSON", sl.Err(err))
		}
		teacher, err := storage.GetTeacher(teacherData.Login, teacherData.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		jwtAuthToken, err := middlewares.GetAuthTokenFromContext(r.Context())
		if err != nil {
			logger.Error("Can't get jwt auth token", sl.Err(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		claims := map[string]interface{}{"id": teacher.ID, "login": teacher.Login, "table": "teachers"}
		_, token, err := jwtAuthToken.Encode(claims)
		if err != nil {
			logger.Error("Can't get jwt", sl.Err(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		if err := json.NewEncoder(w).Encode(map[string]interface{}{"token": token}); err != nil {
			logger.Error("Cannot encode token", sl.Err(err))
		}
	}
}
