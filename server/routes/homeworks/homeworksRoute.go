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
	"strings"
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

		homework, err := storage.AddHomework(homeworkData.Description, lesson.ID, homeworkData.Deadline, homeworkData.MaxScore)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

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

		homeworkId, err := permissions.GetHomeworkIdFromRequest(r)
		if err != nil {
			http.Error(w, "You must send homeworkId as URL part like /groups/{group_id}/lessons/{lesson_id}/homeworks/{homework_id}", http.StatusBadRequest)
			return
		}

		if !storage.IsHomeworkInLesson(homeworkId, lesson.ID) {
			http.Error(w, "Homework not found", http.StatusNotFound)
			return
		}

		err = r.ParseMultipartForm(10 << 20)
		if err != nil {
			http.Error(w, "Can't parse form: "+err.Error(), http.StatusBadRequest)
			logger.Error("Can't parse form", sl.Err(err))
			return
		}

		form := r.MultipartForm
		files := form.File["files"]

		filePaths := make([]string, 0, len(files))
		for _, file := range files {
			fileId, err := storage.AddHomeworkFile(homeworkId, file.Filename)
			if err != nil {
				http.Error(w, "Can't add file", http.StatusInternalServerError)
				logger.Error("Can't add file", sl.Err(err))
				return
			}
			splitted := strings.Split(file.Filename, ".")
			filePath := fmt.Sprintf("files/%d.%s", fileId, splitted[len(splitted)-1])
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
		w.Header().Set("Content-Type", "application/json")
		if err = json.NewEncoder(w).Encode(map[string]interface{}{"added_files": filePaths}); err != nil {
			logger.Error("Can't marshall added files json", sl.Err(err))
		}
	}
}

func GetHomeworksByStudent(logger *slog.Logger, storage *postgres.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		studentId := permissions.GetStudentIdFromContext(r)
		homeworks, err := storage.GetHomeworksByStudent(studentId)
		if err != nil {
			http.Error(w, "Can't get homeworks", http.StatusForbidden)
			logger.Error("Can't get homeworks", sl.Err(err))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(map[string]interface{}{"homeworks": homeworks}); err != nil {
			logger.Error("Can't marshall homeworks json", sl.Err(err))
		}
	}
}

func GetHomeworksByStudentIdInRequest(logger *slog.Logger, storage *postgres.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		studentId, err := permissions.GetStudentIdFromRequest(r)
		if err != nil {
			http.Error(w, "Can't get studentId. You should use /students/{student_id}/homeworks route", http.StatusBadRequest)
			return
		}
		homeworks, err := storage.GetHomeworksByStudent(studentId)
		if err != nil {
			http.Error(w, "Can't get homeworks", http.StatusForbidden)
			logger.Error("Can't get homeworks", sl.Err(err))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(map[string]interface{}{"homeworks": homeworks}); err != nil {
			logger.Error("Can't marshall homeworks json", sl.Err(err))
		}
	}
}

func GetHomeworkById(logger *slog.Logger, storage *postgres.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		homeworkId, err := permissions.GetHomeworkIdFromRequest(r)
		if err != nil {
			http.Error(w, "You must send homeworkId as URL part like /homeworks/{homework_id}", http.StatusBadRequest)
			return
		}
		homework, err := storage.GetHomeworkById(homeworkId)
		if err != nil {
			http.Error(w, "Can't get homework", http.StatusForbidden)
			logger.Error("Can't get homework", sl.Err(err))
			return
		}
		groupId := homework.Lesson.GroupID
		if !storage.IsStudentInGroup(groupId, permissions.GetStudentIdFromContext(r)) {
			http.Error(w, "Student is not in this group", http.StatusForbidden)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(homework); err != nil {
			logger.Error("Can't marshall homework json", sl.Err(err))
		}
	}
}
