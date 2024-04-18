package homeworks

import (
	"encoding/json"
	"log/slog"
	"net/http"
	communicationJson "server/internal/http_server/network/communication/json"
	"server/internal/http_server/permissions"
	"server/internal/lib/logger/sl"
	"server/internal/storage/postgres"
)

func GetHomeworks(logger *slog.Logger, storage *postgres.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lesson, done := permissions.ValidatePermissionsInLesson(w, r, logger, storage)
		if done {
			return
		}

		homeworks := lesson.Homeworks
		if err := json.NewEncoder(w).Encode(homeworks); err != nil {
			logger.Error("Can't marshall homeworks json", sl.Err(err))
		}
	}
}

func AddHomework(logger *slog.Logger, storage *postgres.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lesson, done := permissions.ValidateTeachersPermissionInLesson(w, r, logger, storage)
		if done {
			return
		}

		var homeworkData communicationJson.HomeworkJson
		if err := json.NewDecoder(r.Body).Decode(&homeworkData); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			logger.Error("Can't unmarshal JSON", sl.Err(err))
			return
		}

		homework := storage.AddHomework(homeworkData.Description, lesson.ID, homeworkData.Deadline, homeworkData.MaxScore)
		if err := json.NewEncoder(w).Encode(homework); err != nil {
			logger.Error("Can't marshall homework json", sl.Err(err))
		}
	}
}
