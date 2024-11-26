package main

import (
	"log"
	"module/cmd/api"
	"module/config"
	"module/db"

	"github.com/go-sql-driver/mysql"
)

func main() {
	mysql, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Env.DBUser,
		Passwd:               config.Env.DBPassword,
		Addr:                 config.Env.DBAddress,
		DBName:               config.Env.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}

	db.InitMySQLStorage(mysql)

	server := api.NewAPISever(":8080", mysql)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
