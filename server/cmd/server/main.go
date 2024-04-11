package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/docgen"
	"github.com/go-chi/jwtauth"
	"log/slog"
	"net/http"
	"os"
	"path"
	"server/internal/config"
	"server/internal/http_server/middlewares"
	"server/internal/lib/logger/sl"
	"server/internal/storage/postgres"
	"server/routes/groups"
	"server/routes/students"
	"server/routes/teachers"
)

func main() {
	Start()
}

func Start() {
	dir, _ := os.Getwd()
	configPath := path.Join(dir, "config", "local.yaml")
	cfg := config.MustLoad(configPath)
	log := sl.SetupLogger(cfg.Env)

	storage, err := postgres.New(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.DBConf.Host, cfg.DBConf.User, cfg.DBConf.Password, cfg.DBConf.DBName, cfg.DBConf.Port))

	if err != nil {
		log.Error("Database is not connected", sl.Err(err))
		fmt.Println(err)
		panic("Database is not connected!")
	}

	router := CreateRouter(log, storage)

	// Сборка документации
	doc := docgen.MarkdownDoc{Router: router}
	err = doc.Generate()
	os.WriteFile("res.md", []byte(doc.String()), 0666)

	http.ListenAndServe(":8080", router)
}

func CreateRouter(log *slog.Logger, storage *postgres.Storage) chi.Router {
	tokenAuth := jwtauth.New("HS256", []byte("secret"), nil)

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middlewares.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat) // Хз, надо ли оно нам
	router.Use(middlewares.JWTAuthHolder(tokenAuth))

	router.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Use(middlewares.TeacherAuth)

		// Авторизованные запросы

		r.Post("/students", students.AddStudentHandler(log, storage))
		r.Post("/groups", groups.AddGroup(log, storage))
		r.Post("/groups/{group_id}/students", groups.AddStudentsToGroup(log, storage))
		r.Get("/groups/{group_id}/students", groups.GetStudentsFromGroup(log, storage))
	})
	router.Group(func(r chi.Router) {
		// Неавторизованные запросы
		r.Post("/login", students.LoginStudentHandler(log, storage))
		r.Post("/teachers/login", teachers.LoginTeacher(log, storage))
	})

	return router
}
