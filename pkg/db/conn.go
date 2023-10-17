package db

import "database/sql"

var db *sql.DB

func InitDb(dataSourceName string) error {
	var err error
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return err
	}
	return db.Ping()
}

func getDB() *sql.DB {
	return db
}
