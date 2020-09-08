package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hackersandslackers/golang-helloworld/internal/models"
	"log"
)

// CreateConn Create connection to local MySQL DB with login and pass
func CreateConn(uname string, upass string, dbname string) (*sql.DB, error) {
	// db, err := sql.Open("mysql", "quiz_admin:123@tcp(localhost:3306)/quiz")
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s", uname, upass, dbname))
	if err != nil {
		return nil, err
	}
	log.Printf("Connected to local db with name %s", dbname)
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

// CreateUser Create and insert new user into local db
func CreateUser(db *sql.DB, userName string, userPass string) (int64, error) {

	stmt, err := db.Prepare("INSERT INTO users (username, password) VALUES (?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(userName, userPass)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

// CheckUserPassword Return true if password correct
func CheckUserPassword(db *sql.DB, userName string, userPass string) (bool, error) {

	result, err := db.Query("SELECT password FROM users WHERE username = ?", userName)
	if err != nil {
		return false, err
	}
	defer result.Close()
	var savedPassword string
	for result.Next() {
		err := result.Scan(&savedPassword)
		if err != nil {
			return false, err
		}
	}
	if result.Err() != nil {
		return false, result.Err()
	}
	return savedPassword == userPass, nil
}

// GetWikiPage loads wiki page by title from db
func GetWikiPage(db *sql.DB, title string) (*models.WikiPage, error) {
	result, err := db.Query("SELECT id, title, body FROM wiki WHERE title = ?", title)
	if err != nil {
		return nil, err
	}
	var wp models.WikiPage
	for result.Next() {
		err := result.Scan(&wp.ID, &wp.Title, &wp.Body)
		if err != nil {
			return nil, err
		}
	}
	return &wp, nil
}

// UpdateWikiPage updates body of page with "title"
func UpdateWikiPage(db *sql.DB, title string, body []byte) error {
	result, err := db.Prepare("UPDATE wiki SET body = ? WHERE title = ?")
	if err != nil {
		return err
	}
	_, err = result.Exec(body, title)
	if err != nil {
		return err
	}
	return nil
}
