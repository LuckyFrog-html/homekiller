package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"log/slog"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"server/internal/config"
	"server/internal/http_server/middlewares"
	"server/internal/lib/logger/sl"
	"server/internal/storage/file_storage"
	"server/internal/storage/postgres"
	"server/routes/files"
	"server/routes/groups"
	"server/routes/homeworks"
	"server/routes/lessons"
	"server/routes/students"
	"server/routes/teachers"
)

func main() {
	Start()
}

func Start() {
	dir, _ := os.Getwd()
	var configName string
	if tempPath := os.Getenv("CONFIG_PATH"); tempPath == "" {
		configName = "local.yaml"
	} else {
		configName = tempPath
	}
	configPath := path.Join(dir, "config", configName)
	cfg := config.MustLoad(configPath)
	log := sl.SetupLogger(cfg.Env)

	storage, err := postgres.New(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.DBConf.Host, cfg.DBConf.User, cfg.DBConf.Password, cfg.DBConf.DBName, cfg.DBConf.Port))

	if err != nil {
		log.Error("Database is not connected", sl.Err(err))
		fmt.Println(err)
		panic("Database is not connected!")
	}

	file_storage.InitFileStorage(cfg.StoragePath, log)

	router := CreateRouter(log, storage)

	// Сборка документации
	//doc := docgen.MarkdownDoc{Router: router}
	//err = doc.Generate()
	//os.WriteFile("res.md", []byte(doc.String()), 0666)

	panic(http.ListenAndServe(":8080", router))
}

func CreateRouter(log *slog.Logger, storage *postgres.Storage) chi.Router {
	tokenAuth := jwtauth.New("HS256", []byte("secret"), nil)

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middlewares.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat) // Хз, надо ли оно нам
	//router.Use(middlewares.JWTAuthHolder(tokenAuth))

	router.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Use(middlewares.TeacherAuth)

		// Авторизованные запросы для учителей

		r.Post("/students", students.AddStudentHandler(log, storage))
		r.Post("/groups", groups.AddGroup(log, storage))
		r.Get("/teacher/groups", groups.GetGroupsByTeacher(log, storage))
		r.Post("/students", groups.AddStudentsToGroup(log, storage))
		r.Get("/groups/{group_id}/students", groups.GetStudentsFromGroup(log, storage))

		r.Post("/lessons", lessons.AddLesson(log, storage))
		r.Post("/lessons/{lesson_id}", lessons.MarkStudentAttendance(log, storage)) // TODO: Дописать геттер для списка отмеченных учеников

		r.Post("/homeworks", homeworks.AddHomework(log, storage))
		r.Get("/solves", homeworks.GetHomeworkSolvesByTeacher(log, storage))
		// TODO: Геттер для конкретного решения
		r.Post("/homeworks/{homework_id}/files", homeworks.AddHomeworkFiles(log, storage))
		r.Get("/students/{student_id}/homeworks", homeworks.GetHomeworksByStudentIdInRequest(log, storage))
		// TODO: Геттер для файлов
		// TODO: DELETE для студента
		// TODO: DELETE для группы
		// TODO: DELETE для прикрепления студента к группе
		// TODO: DELETE для урока
		// TODO: DELETE для домашки
		// TODO: DELETE для файла из домашки
	})
	router.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)

		// Авторизованные запросы для студентов
		//r.Get("/groups/{group_id}/lessons", lessons.GetLessons(log, storage))
		//r.Get("/groups/{group_id}/lessons/{lesson_id}", lessons.GetLessonByGroup(log, storage))
		//r.Get("/groups/{group_id}/lessons/{lesson_id}/homeworks", homeworks.GetHomeworks(log, storage))
		//r.Get("/groups", groups.GetGroupsByStudentHandler(log, storage))
		r.Get("/homeworks", homeworks.GetHomeworksByStudent(log, storage))
		r.Get("/homeworks/{homework_id}", homeworks.GetHomeworkById(log, storage))
		r.Post("/homeworks", homeworks.AddHomeworkAnswer(log, storage))
		// TODO: Добавление ответов студентов на домашку
		// TODO: Геттер для конкретного решения
		// TODO: DELETE для ответов на домашку
		// TODO: DELETE для файлов ответа
		// TODO: самовыпил
	})
	router.Group(func(r chi.Router) {
		// Неавторизованные запросы
		r.Post("/login", students.LoginStudentHandler(log, storage, tokenAuth))
		r.Post("/teachers/login", teachers.LoginTeacher(log, storage, tokenAuth))
		dir, _ := os.Getwd()
		filesDir := http.Dir(filepath.Join(dir, "files"))
		r.Handle("/files/*", files.FileHandler(log, storage, filesDir))
	})

	return router
}
