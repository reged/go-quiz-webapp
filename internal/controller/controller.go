package controller

import "database/sql"

type Controller struct {
	Db *sql.DB
}
