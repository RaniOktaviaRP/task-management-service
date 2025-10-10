package app

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	httpSwagger "github.com/swaggo/http-swagger"

	"task-management/controller"
	"task-management/middleware"
)

func WrapHandlerWithHttprouter(handler http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		handler.ServeHTTP(w, r)
	}
}

func WrapHandlerWithJWT(handler httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			handler(w, r, ps)
		})
		middleware.JWTAuth(h).ServeHTTP(w, r)
	}
}

func NewRouter(userController controller.UserController, profileController controller.ProfileController, projectController controller.ProjectController, taskController controller.TaskController) *httprouter.Router {
	router := httprouter.New()

	// Auth & user
	router.POST("/api/users", userController.Register)
	router.POST("/api/login", userController.Login)

	// Refresh token
	router.POST("/api/refresh", WrapHandlerWithJWT(userController.Refresh))

	// Logout
	router.POST("/api/logout", WrapHandlerWithJWT(userController.Logout))

	// CRUD users
	router.GET("/api/users", WrapHandlerWithJWT(userController.FindAll))          // get all users
	router.GET("/api/users/:userId", WrapHandlerWithJWT(userController.FindById)) // get user by ID
	router.PUT("/api/users/:userId", WrapHandlerWithJWT(userController.Update))   // update user
	router.DELETE("/api/users/:userId", WrapHandlerWithJWT(userController.Delete)) // delete user

	// Profile routes
	router.POST("/api/profiles", WrapHandlerWithJWT(profileController.Create))
	router.GET("/api/profiles", WrapHandlerWithJWT(profileController.FindAll))
	router.GET("/api/profiles/by-user/:userId", WrapHandlerWithJWT(profileController.FindByUserId))
	router.GET("/api/profiles/by-id/:profileId", WrapHandlerWithJWT(profileController.FindById))
	router.PUT("/api/profiles/by-id/:profileId", WrapHandlerWithJWT(profileController.Update))
	router.DELETE("/api/profiles/by-id/:profileId", WrapHandlerWithJWT(profileController.Delete))

	// Projects API
	router.POST("/api/projects", WrapHandlerWithJWT(projectController.Create))
	router.GET("/api/projects", WrapHandlerWithJWT(projectController.FindAll))
	router.GET("/api/projects/by-user/:userId", WrapHandlerWithJWT(projectController.FindByUserId))
	router.GET("/api/projects/by-id/:id", WrapHandlerWithJWT(projectController.FindById))
	router.PUT("/api/projects/by-id/:id", WrapHandlerWithJWT(projectController.Update))
	router.DELETE("/api/projects/by-id/:id", WrapHandlerWithJWT(projectController.Delete))

	// Tasks API
	router.POST("/api/tasks", WrapHandlerWithJWT(taskController.Create))
	router.GET("/api/tasks", WrapHandlerWithJWT(taskController.FindAll))
	// Gunakan path yang lebih spesifik dan tidak ambigu
	router.GET("/api/tasks/id/:id", WrapHandlerWithJWT(taskController.FindById))
	router.PUT("/api/tasks/:id", WrapHandlerWithJWT(taskController.Update))
	router.DELETE("/api/tasks/id/:id", WrapHandlerWithJWT(taskController.Delete))
	router.GET("/api/tasks/project/:projectId", WrapHandlerWithJWT(taskController.FindByProjectId))

	// swagger docs
	router.GET("/swagger/*any", WrapHandlerWithHttprouter(middleware.CORS(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:3001/swagger/doc.json"),
	))))

	return router
}
