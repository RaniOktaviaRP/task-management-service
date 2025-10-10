package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"

	"task-management/app"
	"task-management/controller"
	"task-management/helper"
	"task-management/middleware"
	"task-management/repository"
	"task-management/service"

	_ "task-management/docs" // Swagger docs
)

// @title User API
// @version 1.0
// @description API untuk manajemen users dengan PostgreSQL.
// @host localhost:3001
// @BasePath /api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	// Inisialisasi database
	db := app.NewDB()

	// Inisialisasi validator
	validate := validator.New()

	// Buat repository
	userRepository := repository.NewUserRepository(db)
	profileRepository := repository.NewProfileRepository(db)
	refreshTokenRepository := repository.NewRefreshTokenRepository()

	// JWT secret langsung dari konfigurasi service
	jwtSecret := []byte("rahasia") // atau ambil dari file konfigurasi/service

	// Buat service (kirim repository + db + jwtSecret)
	userService := service.NewUserService(userRepository, profileRepository, refreshTokenRepository, db, jwtSecret)

	// Buat profile service
	profileService := service.NewProfileService(profileRepository, db)

	// Buat project repository
	projectRepository := repository.NewProjectRepository(db)

	// Buat project service
	projectService := service.NewProjectService(projectRepository, db)

	// Buat task repository dengan sql.DB
	taskRepository := repository.NewTaskRepository(db)

	// Buat task service dengan validator
	taskService := service.NewTaskService(taskRepository, validate)

	// Buat controller
	userController := controller.NewUserController(userService)
	profileController := controller.NewProfileController(profileService)
	projectController := controller.NewProjectController(projectService)
	taskController := controller.NewTaskController(taskService)

	// Update router initialization
	router := app.NewRouter(userController, profileController, projectController, taskController)

	// Jalankan server dengan middleware CORS
	server := &http.Server{
		Addr:    "localhost:3001",
		Handler: middleware.CORS(router),
	}

	helper.PanicIfError(server.ListenAndServe())
}
