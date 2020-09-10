package controller

import (
	"database/sql"

	"github.com/reged/go-quiz-webapp/internal/service"
)

type Controller struct {
	Db          *sql.DB
	QuizService *service.QuizService
}
