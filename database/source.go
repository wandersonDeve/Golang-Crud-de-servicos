package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func DatabaseConnect() *sql.DB {
	connection := "postgres://whatsapp_30am_user:yJ05pBjzr5y34hjHCcmviUHQXyzJRD6r@dpg-cniirfol6cac7398nhj0-a.oregon-postgres.render.com/services-check"

	db, err := sql.Open("postgres", connection)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
