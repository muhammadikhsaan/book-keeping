package config

import (
	"database/sql"
	"log"
)

var databases *sql.DB

//Getdatabase is init database connection
func Getdatabase() *sql.DB {
	if databases == nil {
		var connection = Getenv("DB_USER") + ":" + Getenv("DB_PASS") + "@tcp(" + Getenv("DB_HOST") + ")/" + Getenv("DB_NAME")
		db, err := sql.Open(Getenv("DB"), connection)

		if err != nil {
			log.Fatal(err)
		}

		if err := db.Ping(); err != nil {
			log.Fatal(err)
		}
		databases = db
	}
	return databases
}
