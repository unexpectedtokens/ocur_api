package migrations

import (
	"fmt"
	"io/ioutil"

	"github.com/unexpectedtokens/ocur_api/db"
)

//RunMigrations runs the migrations in the migrations.sql file
func RunMigrations() error{
	defer func(){
		if r := recover(); r !=nil{
			fmt.Println(r)
		}
	}()
	
	db.InitDB()
	defer db.DBCon.Close()
	file, err := ioutil.ReadFile("db/migrations/migrations.sql")
	if err != nil{
		panic(err)

	}
	fmt.Println("Running migrations")
	_, err = db.DBCon.Exec(string(file))
	if err != nil {
		panic(fmt.Errorf("error migrating: %s", err.Error()))
	}
	fmt.Println("[SUCCES]: Migrations ran succesfully")
	return nil
}

//Flush deletes all tables from postgres
// func Flush(){
// 	query := "DROP TABLE users CASCADE; DROP TABLE profiles; DROP TABLE jwt_auth; DROP TABLE recipes CASCADE; DROP TABLE ingredients_from_recipe; DROP TABLE methods_from_recipe; DROP TABLE food_ingredient CASCADE; DROP TABLE ingredients_from_foodingredient_from_recipe;"
// 	db.InitDB()
// 	defer db.DBCon.Close()
// 	_, err := db.DBCon.Exec(query)
// 	if err != nil{
// 		panic(err)
// 	}
// 	fmt.Println("[SUCCES]: Flush ran succesfully")
// }