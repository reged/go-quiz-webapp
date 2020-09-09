package main

import (
	"log"
	"net/http"
	"time"

	"github.com/reged/go-quiz-webapp/internal/database"

	"github.com/reged/go-quiz-webapp/internal/presentation/controller"
	"github.com/reged/go-quiz-webapp/internal/presentation/router"
	"github.com/reged/go-quiz-webapp/internal/service"

	_ "github.com/go-sql-driver/mysql"
)

// Initiate web server
func main() {

	log.Print("Server started\n")
	db, err := database.CreateConn("quiz_admin", "123", "quiz")
	if err != nil {
		log.Fatal("Can't connect database")
	}
	quizService := &service.QuizService{}
	c := &controller.Controller{
		Db:          db,
		QuizService: quizService,
	}

	r := router.InitRouter(c)
	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:80",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	// wp = GetWikiPage(db, "test")
	// _ = UpdateWikiPage(db, "test", []byte("Hello"))

	log.Fatal(srv.ListenAndServe())
}
