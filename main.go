package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/suhaibkhatr/go-practice/cmd/api"
	"github.com/suhaibkhatr/go-practice/config"
	"github.com/suhaibkhatr/go-practice/db"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatalln(err)
	}

	initStorage(db)

	server := api.NewAPIServer(":8080", db)

	if err := server.Run(); err != nil {
		log.Fatalln(err)
	}
}
func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("DB: Successfully connected!")
}
