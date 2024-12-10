package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	docs "github.com/clrajapaksha/to-do-list-app/docs"

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

// getAlbums responds with the list of all albums as JSON.
// @Tags Albums
// @Summary album summary
// @Description album description
// @Router / [get]
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func (server *APIServer) Run() error {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	// docs.SwaggerInfo.Host = "petstore.swagger.io"
	// docs.SwaggerInfo.BasePath = "/v2"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	//router := http.NewServeMux()
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	// router.HandleFunc("/l", hello)
	router.Get("/", hello)

	// router.HandleFunc("/swagger/*any", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/swagger/doc.json")))
	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), //The url pointing to API definition
	))

	return http.ListenAndServe(server.addresss, router)

}
