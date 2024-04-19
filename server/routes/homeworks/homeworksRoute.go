package homeworks

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
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
		w.Header().Set("Content-Type", "application/json")
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
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(homework); err != nil {
			logger.Error("Can't marshall homework json", sl.Err(err))
		}
	}
}

func AddHomeworkFiles(logger *slog.Logger, storage *postgres.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lesson, done := permissions.ValidateTeachersPermissionInLesson(w, r, logger, storage)
		if done {
			return
		}

		homeworkId, err := permissions.GetHomeworkIdFromRequest(w, r)
		if err != nil {
			return
		}

		if !storage.IsHomeworkInLesson(homeworkId, lesson.ID) {
			http.Error(w, "Homework not found", http.StatusNotFound)
			return
		}

		err = r.ParseMultipartForm(10 << 20)
		if err != nil {
			http.Error(w, "Can't parse form", http.StatusBadRequest)
			logger.Error("Can't parse form", sl.Err(err))
			return
		}

		form := r.MultipartForm
		files := form.File["files"]

		filePaths := make([]string, 0, len(files))
		for _, file := range files {
			filePath := fmt.Sprintf("files/%s", file.Filename)
			f, err := file.Open()
			if err != nil {
				http.Error(w, fmt.Sprintf("Can't open file %s", file.Filename), http.StatusBadRequest)
				return
			}
			defer f.Close()
			dst, err := os.Create(filePath)
			if err != nil {
				http.Error(w, "Can't create file", http.StatusInternalServerError)
				logger.Error("Can't create file", sl.Err(err))
				return
			}
			defer dst.Close()
			if _, err = io.Copy(dst, f); err != nil {
				http.Error(w, fmt.Sprintf("Can't copy file %s", file.Filename), http.StatusBadRequest)
				logger.Error("Can't copy file", sl.Err(err))
				return
			}

			filePaths = append(filePaths, filePath)
		}
		storage.AddHomeworkFiles(homeworkId, filePaths)
		if err = json.NewEncoder(w).Encode(map[string]interface{}{"added_files": filePaths}); err != nil {
			logger.Error("Can't marshall added files json", sl.Err(err))
		}
	}
}
