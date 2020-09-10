package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/reged/go-quiz-webapp/internal/presentation/controller"
)

// Route declaration
func InitRouter(c *controller.Controller) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", c.IndexPage).Methods("GET")
	r.HandleFunc("/register", c.RegisterPageHandler).Methods("GET")
	r.HandleFunc("/register", c.RegisterHandler).Methods("POST")
	r.HandleFunc("/login", c.LoginPageHandler).Methods("GET")
	r.HandleFunc("/login", c.LoginHandler).Methods("POST")
	r.HandleFunc("/rebuild", c.RebuildHandler).Methods("GET")
	r.HandleFunc("/test", c.TestHandler).Methods("GET")
	fileServer := http.FileServer(http.Dir("./static"))
	r.PathPrefix("/").Handler(http.StripPrefix("/static", fileServer))
	return r
}
