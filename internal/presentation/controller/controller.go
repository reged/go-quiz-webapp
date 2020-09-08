package controller

import (
	"database/sql"
	"github.com/hackersandslackers/golang-helloworld/internal/service"
)

type Controller struct {
	Db          *sql.DB
	QuizService *service.QuizService
}
