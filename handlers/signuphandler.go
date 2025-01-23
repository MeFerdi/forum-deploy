package handlers

import "database/sql"

var GlobalDB *sql.DB

func InitDB(database *sql.DB) {
	GlobalDB = database
}

