package main

import (
	"fmt"
	"os"

	_ "github.com/jackc/pgx"
	"github.com/joho/godotenv"
	"github.com/unexpectedtokens/ocur_api/migrations"
)




func main(){
	err := godotenv.Load()
	if err != nil{
		panic(err)
	}
	fmt.Println(len(os.Args))
	//db.InitDB()
	if len(os.Args) > 1{
		for _, x := range os.Args{
			if x == "migrate"{
				if err = migrations.RunMigrations(); err != nil{
					panic(err)
				}
			}
		}
	}

}