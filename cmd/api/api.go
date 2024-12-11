package api

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/clrajapaksha/to-do-list-app/cache"
	"github.com/clrajapaksha/to-do-list-app/config"
	docs "github.com/clrajapaksha/to-do-list-app/docs"
	"github.com/clrajapaksha/to-do-list-app/entities"
	"github.com/clrajapaksha/to-do-list-app/repository"
	"github.com/clrajapaksha/to-do-list-app/services/task"
	"github.com/clrajapaksha/to-do-list-app/utils"

	httpSwagger "github.com/swaggo/http-swagger"
)

type APIServer struct {
	addresss string
	db       *sql.DB
}

func NewAPIServer(address string, db *sql.DB) *APIServer {
	return &APIServer{
		addresss: address,
		db:       db,
	}
}

// health responds with success beat signal.
// @Tags Health
// @Summary Health endpoint
// @Description health check
// @Router /health [get]
func health(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, http.StatusOK, map[string]string{"status": "success"})
}

// @contact.name   Chathuranga Rajapaksha
// @contact.url    https://www.linkedin.com/in/clrajapaksha/
// @contact.email  clrajapaksha@gmail.com

func (server *APIServer) Run() error {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "To Do List API"
	docs.SwaggerInfo.Description = "This is a sample REST API in Golang."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	//docs.SwaggerInfo.Host = "http://localhost:8080"
	//docs.SwaggerInfo.BasePath = "/api/v1"

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	cache := cache.New[string, entities.Task]()

	taskRepository := repository.NewDynamoDBRepository()
	taskHandler := task.NewHandler(taskRepository, cache)
	taskRouter := chi.NewRouter()
	taskHandler.RegisterRoutes(taskRouter)

	router.Mount("/", taskRouter)
	router.Get("/health", health)
	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(config.Envs.AppUrl+"/swagger/doc.json"),
	))

	return http.ListenAndServe(server.addresss, router)

}
