package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/nuttapon-first/task-clean-arch/configs"
	"github.com/nuttapon-first/task-clean-arch/modules/server"
	"github.com/nuttapon-first/task-clean-arch/pkg/databases"
)

func init() {
	err := godotenv.Load("local.env")
	if err != nil {
		log.Printf("[only on local machine] please consider environment variables: %s\n", err)
	}
}

func main() {
	config := new(configs.Configs)

	config.App.Port = os.Getenv("GIN_PORT")

	config.Database.DriverName = os.Getenv("DB_DRIVER")
	config.Database.DbName = os.Getenv("DB_TABLE_NAME")

	db, err := databases.NewSQLDBConnection(config)
	if err != nil {
		log.Fatalln(err.Error())
	}

	s := server.NewServer(config, db)
	s.Start()
}
