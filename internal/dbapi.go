package internal

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// CreateConn Create connection to local MySQL DB with login and pass
func CreateConn(uname string, upass string, dbname string) *sql.DB {
	// db, err := sql.Open("mysql", "quiz_admin:123@tcp(localhost:3306)/quiz")
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s", uname, upass, dbname))
	if err != nil {
		panic(err.Error())
	}
	log.Printf("Connected to local db with name %s", dbname)
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	return db
}

// CreateUser Create and insert new user into local db
func CreateUser(db *sql.DB, uname string, upass string) (int64, error) {

	stmt, err := db.Prepare("INSERT INTO users (username, password) VALUES (?, ?)")
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(uname, upass)
	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

// CheckUserPassword Return true if password correct
func CheckUserPassword(db *sql.DB, uname string, upass string) (bool, error) {

	result, err := db.Query("SELECT password FROM users WHERE username = ?", uname)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var userpassword string
	for result.Next() {
		err := result.Scan(&userpassword)
		if err != nil {
			panic(err.Error())
		}
	}
	return userpassword == upass, nil
}

// GetWikiPage loads wiki page by title from db
func GetWikiPage(db *sql.DB, title string) *WikiPage {
	result, err := db.Query("SELECT id, title, body FROM wiki WHERE title = ?", title)
	if err != nil {
		panic(err.Error())
	}
	var wp WikiPage
	for result.Next() {
		err := result.Scan(&wp.ID, &wp.Title, &wp.Body)
		if err != nil {
			panic(err.Error())
		}
	}
	return &wp
}

// UpdateWikiPage updates body of page with "title"
func UpdateWikiPage(db *sql.DB, title string, body []byte) error {
	result, err := db.Prepare("UPDATE wiki SET body = ? WHERE title = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = result.Exec(body, title)
	if err != nil {
		panic(err.Error())
	}
	return nil
}
