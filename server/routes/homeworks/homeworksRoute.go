package homeworks

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"mime"
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

		_, params, err := mime.ParseMediaType(r.Header.Get("Content-Disposition"))
		if err != nil {
			http.Error(w, "Can't parse file", http.StatusBadRequest)
			logger.Error("Can't parse file", sl.Err(err))
			return
		}
		splited := strings.Split(params["filename"], ".")
		extension := splited[len(splited)-1]

		fileId, err := storage.AddHomeworkFile(homeworkId, extension)
		if err != nil {
			http.Error(w, "Can't add file", http.StatusInternalServerError)
			logger.Error("Can't add file", sl.Err(err))
			return
		}

		filePath := fmt.Sprintf("files/teachers/%d.%s", fileId, extension)
		f, err := os.Create(filePath)
		if err != nil {
			http.Error(w, "Can't create file", http.StatusInternalServerError)
			logger.Error("Can't create file", sl.Err(err))
			return
		}
		defer f.Close()

		if _, err = io.Copy(f, r.Body); err != nil {
			http.Error(w, "Can't copy file", http.StatusInternalServerError)
			logger.Error("Can't copy file", sl.Err(err))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err = json.NewEncoder(w).Encode(map[string]interface{}{"file_id": fileId}); err != nil {
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

		solution, err := storage.AddHomeworkAnswer(answerData.HomeworkId, studentId, answerData.Text)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(solution); err != nil {
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
		studentId := permissions.GetStudentIdFromContext(r)
		solutionId, err := permissions.GetSolutionIdFromRequest(r)
		if err != nil {
			http.Error(w, "You must send homeworkId as URL part like /solutions/{solution_id}/files", http.StatusBadRequest)
			return
		}

		solution, err := storage.GetSolutionById(solutionId)
		if err != nil {
			http.Error(w, "Solution not found", http.StatusNotFound)
			logger.Error("Solution not found", sl.Err(err))
			return
		}

		if solution.Student.ID != studentId {
			http.Error(w, "Student is not owner of this solution", http.StatusForbidden)
			return
		}

		_, params, err := mime.ParseMediaType(r.Header.Get("Content-Disposition"))
		if err != nil {
			http.Error(w, "Can't parse file", http.StatusBadRequest)
			logger.Error("Can't parse file", sl.Err(err))
			return
		}
		splited := strings.Split(params["filename"], ".")
		extension := splited[len(splited)-1]

		fileId, err := storage.AddSolutionFile(solutionId, extension)
		if err != nil {
			http.Error(w, "Can't add file", http.StatusInternalServerError)
			logger.Error("Can't add file", sl.Err(err))
			return
		}

		filePath := fmt.Sprintf("files/students/%d.%s", fileId, extension)
		f, err := os.Create(filePath)
		if err != nil {
			http.Error(w, "Can't create file", http.StatusInternalServerError)
			logger.Error("Can't create file", sl.Err(err))
			return
		}
		defer f.Close()

		if _, err = io.Copy(f, r.Body); err != nil {
			http.Error(w, "Can't copy file", http.StatusInternalServerError)
			logger.Error("Can't copy file", sl.Err(err))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(map[string]interface{}{"file_id": fileId}); err != nil {
			logger.Error("Can't marshall file_id json", sl.Err(err))
		}
	}
}

func GetHomeworkByIdByTeacher(logger *slog.Logger, storage *postgres.Storage) http.HandlerFunc {
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
		teacherId, err := middlewares.GetTeacherIdFromContext(r.Context())
		if err != nil {
			http.Error(w, "Can't get teacherId", http.StatusNotFound)
			logger.Error("Can't get teacherId", sl.Err(err))
			return
		}
		if !storage.IsTeacherInGroup(homework.Lesson.GroupID, teacherId) {
			http.Error(w, "Teacher is not owner of this group", http.StatusForbidden)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(homework); err != nil {
			logger.Error("Can't marshall homework json", sl.Err(err))
		}
	}
}

func GetHomeworkSolutions(logger *slog.Logger, storage *postgres.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		homeworkId, err := permissions.GetHomeworkIdFromRequest(r)
		if err != nil {
			http.Error(w, "You must send homeworkId as URL part like /homeworks/{homework_id}/solves", http.StatusBadRequest)
			return
		}
		teacherId, err := middlewares.GetTeacherIdFromContext(r.Context())
		if err != nil {
			http.Error(w, "Can't get teacherId", http.StatusNotFound)
			logger.Error("Can't get teacherId", sl.Err(err))
			return
		}
		homework, err := storage.GetHomeworkById(homeworkId)
		if err != nil {
			http.Error(w, "Can't get homework", http.StatusNotFound)
			logger.Error("Can't get homework", sl.Err(err))
			return
		}

		if isTeacherInGroup := storage.IsTeacherInGroup(homework.Lesson.GroupID, teacherId); !isTeacherInGroup {
			http.Error(w, "Teacher is not owner of this group", http.StatusForbidden)
			return
		}

		solutions, err := storage.GetHomeworkSolutions(homeworkId)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(map[string]any{"solutions": solutions}); err != nil {
			logger.Error("Can't marshall homework solutions json", sl.Err(err))
		}
	}
}

func GetHomeworkSolveReviewsByIdReviews(logger *slog.Logger, storage *postgres.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		solveId, err := permissions.GetSolveIdFromRequest(r)
		if err != nil {
			http.Error(w, "You must send solveId as URL part like /solves/{solve_id}/reviews", http.StatusBadRequest)
			return
		}
		studentId := permissions.GetStudentIdFromContext(r)

		solve, err := storage.GetHomeworkSolveById(solveId)
		if err != nil {
			http.Error(w, "Can't get solve", http.StatusNotFound)
			logger.Error("Can't get solve", sl.Err(err))
			return
		}

		if solve.StudentID != studentId {
			http.Error(w, "Student is not owner of this solve", http.StatusForbidden)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(map[string]any{"reviews": solve}); err != nil {
			logger.Error("Can't marshall solve json", sl.Err(err))
		}
	}
}

func GetHomeworkSolves(logger *slog.Logger, storage *postgres.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		homeworkId, err := permissions.GetHomeworkIdFromRequest(r)
		if err != nil {
			http.Error(w, "You must send homeworkId as URL part like /homeworks/{homework_id}/solutions", http.StatusBadRequest)
			return
		}
		studentId := permissions.GetStudentIdFromContext(r)

		homework, err := storage.GetHomeworkById(homeworkId)
		if err != nil {
			http.Error(w, "Can't get homework", http.StatusNotFound)
			logger.Error("Can't get homework", sl.Err(err))
			return
		}

		if !storage.IsStudentInGroup(homework.Lesson.GroupID, studentId) {
			http.Error(w, "Student is not in this group", http.StatusForbidden)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(map[string]any{"solutions": homework.HomeworkAnswers}); err != nil {
			logger.Error("Can't marshall homework solves json", sl.Err(err))
		}
	}
}

func AddHomeworkSolveReview(logger *slog.Logger, storage *postgres.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		solveId, err := permissions.GetSolveIdFromRequest(r)
		if err != nil {
			http.Error(w, "You must send solveId as URL part like /solves/{solve_id}/reviews", http.StatusBadRequest)
			return
		}

		teacherId, err := middlewares.GetTeacherIdFromContext(r.Context())

		if !storage.IsTeacherInSolve(teacherId, solveId) {
			http.Error(w, "Teacher is not owner of this solve", http.StatusForbidden)
			return
		}

		var reviewData communicationJson.HomeworkSolveReviewJson
		if err := json.NewDecoder(r.Body).Decode(&reviewData); err != nil {
			http.Error(w, "Can't unmarshal JSON", http.StatusBadRequest)
		}

		review, err := storage.AddHomeworkSolveReview(solveId, reviewData.Comment, reviewData.Score)
		if err != nil {
			http.Error(w, "Can't add review", http.StatusInternalServerError)
			logger.Error("Can't add review", sl.Err(err))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(review); err != nil {
			logger.Error("Can't marshall review json", sl.Err(err))
		}
	}
}
