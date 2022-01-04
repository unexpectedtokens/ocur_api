package main

import (
	"os"

	_ "github.com/jackc/pgx"
	"github.com/joho/godotenv"
	"github.com/unexpectedtokens/ocur_api/db"
	"github.com/unexpectedtokens/ocur_api/migrations"
	"github.com/unexpectedtokens/ocur_api/router"
)

func main() {
	if os.Getenv("mode") != "production" {
		err := godotenv.Load()
		if err != nil {
			panic(err)
		}
	}
	if len(os.Args) > 1 {
		for _, x := range os.Args {
			if x == "migrate" {
				if err := migrations.RunMigrations(); err != nil {
					panic(err)
				}
			}
		}
	} else {
		db.InitDB()
		router.SetUpRoutes()
	}

}
