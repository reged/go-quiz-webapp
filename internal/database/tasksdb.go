package database

import (
	"database/sql"

	"github.com/reged/go-quiz-webapp/internal/models"

	// Load sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

// CreateNewTask Insert new task into db
func CreateNewTask(db *sql.DB, t models.Task) error {
	stmt, err := db.Prepare("INSERT INTO tasks (title, description) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(t.Title, t.Description)
	if err != nil {
		return err
	}

	return nil
}

func GetTasks(db *sql.DB) ([]models.Task, error) {
	result, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var tasks []models.Task
	var temp models.Task

	for result.Next() {
		err := result.Scan(&temp.ID, &temp.Title, &temp.Description)
		if err != nil {
			return tasks, err
		}
		tasks = append(tasks, temp)
	}

	return tasks, nil
}
