package main

import (
	"api-go/cmd/api"
	"api-go/db"
	"database/sql"
	"log"

	config "api-go/config"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	env := config.InitConfig()

	db, err := db.NewMYSQLStorage(mysql.Config{
		User:                 env.DBUser,
		Passwd:               env.DBPassword,
		DBName:               env.DBName,
		Addr:                 env.DBAddress,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(":8080", db)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

}

func initStorage(db *sql.DB) {
	err := db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Db is successfully connected")
}
