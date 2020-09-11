package database

import (
	"database/sql"

	// Load sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

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

// GetUserCookie Return user cookie
func GetUserCookie(db *sql.DB, userName string) (string, error) {

	result, err := db.Query("SELECT cookie FROM users WHERE username = ?", userName)
	if err != nil {
		return "", err
	}
	defer result.Close()
	var cookie string
	for result.Next() {
		err := result.Scan(&cookie)
		if err != nil {
			return "", err
		}
	}
	if result.Err() != nil {
		return "", result.Err()
	}
	return cookie, nil
}

// SetUserCookie Set new cookie for user
func SetUserCookie(db *sql.DB, userName string, cookie string) error {

	result, err := db.Prepare("UPDATE users SET cookie = ? WHERE username = ?")
	if err != nil {
		return err
	}
	defer result.Close()
	_, err = result.Exec(cookie, userName)
	if err != nil {
		return err
	}
	return nil
}
