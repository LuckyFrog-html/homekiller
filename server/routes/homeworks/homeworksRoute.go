package homeworks

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"server/internal/http_server/middlewares"
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
		var homeworkData communicationJson.HomeworkJson
		if err := json.NewDecoder(r.Body).Decode(&homeworkData); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			logger.Error("Can't unmarshal JSON", sl.Err(err))
			return
		}

		lesson, err := storage.GetLessonById(homeworkData.LessonId)
		if err != nil {
			http.Error(w, "Can't get lesson", http.StatusNotFound)
			logger.Error("Can't get lesson", sl.Err(err))
			return
		}

		teacherId, err := middlewares.GetTeacherIdFromContext(r.Context())
		if err != nil {
			http.Error(w, "Can't get teacherId", http.StatusNotFound)
			logger.Error("Can't get teacherId", sl.Err(err))
			return
		}

		if !storage.IsTeacherInGroup(lesson.GroupID, teacherId) {
			http.Error(w, "Teacher is not owner of this group", http.StatusForbidden)
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
		homeworkId, err := permissions.GetHomeworkIdFromRequest(r)
		if err != nil {
			http.Error(w, "You must send homeworkId as URL part like /homeworks/{homework_id}/files", http.StatusBadRequest)
			return
		}

		teacherId, err := middlewares.GetTeacherIdFromContext(r.Context())
		if err != nil {
			http.Error(w, "Can't get teacherId", http.StatusNotFound)
			logger.Error("Can't get teacherId", sl.Err(err))
			return
		}

		lesson, err := storage.GetLessonByHomeworkId(homeworkId)
		if err != nil {
			http.Error(w, "Can't get lesson", http.StatusNotFound)
			logger.Error("Can't get lesson", sl.Err(err))
			return
		}
		if lesson.Group.TeacherID != teacherId {
			http.Error(w, "Teacher is not owner of this group", http.StatusForbidden)
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

func AddHomeworkAnswer(logger *slog.Logger, storage *postgres.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var answerData communicationJson.HomeworkAnswerJson
		if err := json.NewDecoder(r.Body).Decode(&answerData); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			logger.Error("Can't unmarshal JSON", sl.Err(err))
			return
		}

		homework, err := storage.GetHomeworkById(answerData.HomeworkId)
		if err != nil {
			http.Error(w, "Can't get homework", http.StatusNotFound)
			logger.Error("Can't get homework", sl.Err(err))
			return
		}
		studentId := permissions.GetStudentIdFromContext(r)
		if !storage.IsStudentInGroup(homework.Lesson.GroupID, studentId) {
			http.Error(w, "Student is not in this group", http.StatusForbidden)
			return
		}

		err = storage.AddHomeworkAnswer(answerData.HomeworkId, studentId, answerData.Text)
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

func GetHomeworkSolvesByTeacher(logger *slog.Logger, storage *postgres.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		teacherId, err := middlewares.GetTeacherIdFromContext(r.Context())
		if err != nil {
			http.Error(w, "Can't get teacherId", http.StatusNotFound)
			logger.Error("Can't get teacherId", sl.Err(err))
			return
		}

		homeworkSolves, err := storage.GetHomeworkSolvesByTeacher(teacherId)
		if err != nil {
			http.Error(w, "Can't get homework answers", http.StatusForbidden)
			logger.Error("Can't get homework answers", sl.Err(err))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(map[string]interface{}{"solves": homeworkSolves}); err != nil {
			logger.Error("Can't marshall homeworks json", sl.Err(err))
		}
	}
}

func GetHomeworkSolveByTeacher(logger *slog.Logger, storage *postgres.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		teacherId, err := middlewares.GetTeacherIdFromContext(r.Context())
		if err != nil {
			http.Error(w, "Can't get teacherId", http.StatusNotFound)
			logger.Error("Can't get teacherId", sl.Err(err))
			return
		}

		solveId, err := permissions.GetSolveIdFromRequest(r)
		if err != nil {
			http.Error(w, "You must send solveId as URL part like /solves/{solve_id}", http.StatusBadRequest)
			return
		}

		solve, err := storage.GetHomeworkSolveById(solveId)
		if err != nil {
			http.Error(w, "Can't get solve", http.StatusForbidden)
			logger.Error("Can't get solve", sl.Err(err))
			return
		}

		if storage.IsTeacherInSolve(teacherId, solve.ID) {
			http.Error(w, "Teacher is not owner of homework this solve", http.StatusForbidden)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(solve); err != nil {
			logger.Error("Can't marshall solve json", sl.Err(err))
		}
	}
}

func AddFiles(logger *slog.Logger, storage *postgres.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		homeworkId, err := permissions.GetHomeworkIdFromRequest(r)
		if err != nil {
			http.Error(w, "You must send homeworkId as URL part like /homeworks/{homework_id}/files", http.StatusBadRequest)
			return
		}

		teacherId, err := middlewares.GetTeacherIdFromContext(r.Context())
		if err != nil {
			http.Error(w, "Can't get teacherId", http.StatusNotFound)
			logger.Error("Can't get teacherId", sl.Err(err))
			return
		}

		lesson, err := storage.GetLessonByHomeworkId(homeworkId)
		if err != nil {
			http.Error(w, "Can't get lesson", http.StatusNotFound)
			logger.Error("Can't get lesson", sl.Err(err))
			return
		}
		if lesson.Group.TeacherID != teacherId {
			http.Error(w, "Teacher is not owner of this group", http.StatusForbidden)
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
		_ = form.File["files"]

		// TODO: Дописать
		//filePaths := make([]string, 0, len(files))
		//for _, file := range files {
		//	fileId, err := storage.AddHomeworkFile(homeworkId, file.Filename)
		//	if err != nil {
		//		http.Error(w, "Can't add file", http.StatusInternalServerError)
		//		logger.Error("Can't add file", sl.Err(err))
		//		return
		//	}
		//
		//	splitted := strings.Split(file.Filename, ".")
		//	filePath := fmt.Sprintf("files/%d.%s", fileId, splitted[len(splitted)-1])
		//	f, err := file.Open()
		//	if err != nil {
		//		http.Error(w, fmt.Sprintf("Can't open file %s", file.Filename), http.StatusBadRequest)
		//		return
		//	}
		//	defer f.Close()
		//
		//	dst, err := os.Create(filePath)
		//	if err !=
	}
}
