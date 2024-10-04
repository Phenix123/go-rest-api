package db

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
	env "testApi/config"
)

func New() *sql.DB {
	db, _ := setDBConn()

	return db
}

func setDBConn() (*sql.DB, error) {
	config := env.GetConfig()
	cfg := mysql.Config{
		User:                 config.DB_USER,
		Passwd:               config.DB_PASSWORD,
		Addr:                 config.DB_HOST + ":" + config.DB_PORT,
		DBName:               config.DB_NAME,
		AllowNativePasswords: true,
		Net:                  "tcp",
	}
	// Get a database handle.
	var err error
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	return db, err
}
