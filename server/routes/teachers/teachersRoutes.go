package teachers

import (
	"encoding/json"
	"github.com/go-chi/jwtauth"
	"log/slog"
	"net/http"
	communicationJson "server/internal/http_server/network/communication/json"
	"server/internal/lib/logger/sl"
	"server/internal/storage/postgres"
)

func LoginTeacher(logger *slog.Logger, storage *postgres.Storage, authToken *jwtauth.JWTAuth) http.HandlerFunc {
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

		claims := map[string]interface{}{"id": teacher.ID, "login": teacher.Login, "table": "teachers"}
		_, token, err := authToken.Encode(claims)
		if err != nil {
			logger.Error("Can't get jwt", sl.Err(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(map[string]interface{}{"token": token}); err != nil {
			logger.Error("Cannot encode token", sl.Err(err))
		}
	}
}
