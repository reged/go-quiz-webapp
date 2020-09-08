package main

import (
	"github.com/hackersandslackers/golang-helloworld/internal/database"
	"github.com/hackersandslackers/golang-helloworld/internal/presentation/controller"
	"github.com/hackersandslackers/golang-helloworld/internal/presentation/router"
	"github.com/hackersandslackers/golang-helloworld/internal/service"
	"log"
	"net/http"
	"time"

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
