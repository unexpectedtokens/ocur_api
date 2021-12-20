package db

import (
	"fmt"
	"os"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

var DBCon *sqlx.DB

const (
	pghost = "PGHOST"
	pgport = "PGPORT"
	pguser = "PGUSER"
	pgpw   = "PGPW"
	pgdb   = "PGDB"
)

//InitDB is a function to initialize the database connection. The query is to be replaced with a migration system
func InitDB() {
	var err error
	config := dbConfig()
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config[pghost], config[pgport], config[pguser], config[pgpw], config[pgdb])
	DBCon, err = sqlx.Connect("pgx", psqlInfo)

	if err != nil {
		panic(err)
	}
	err = DBCon.Ping()
	if err != nil {
		panic(err)
	}
	// DBCon.SetConnMaxLifetime(5*time.Minute)
	// DBCon.SetMaxOpenConns(5)
	// DBCon.SetMaxIdleConns(5000)

	fmt.Println("Succesfully connected to DB")
}

func dbConfig() map[string]string {
	conf := make(map[string]string)
	host, ok := os.LookupEnv(pghost)
	if !ok {
		panic("PGHOST environment variable required but not set")
	}
	port, ok := os.LookupEnv(pgport)
	if !ok {
		panic("PGPORT environment variable required but not set")
	}
	user, ok := os.LookupEnv(pguser)
	if !ok {
		panic("PGUSER environment variable required but not set")
	}
	password, ok := os.LookupEnv(pgpw)
	if !ok {
		panic("PGPASS environment variable required but not set")
	}
	name, ok := os.LookupEnv(pgdb)
	if !ok {
		panic("PGNAME environment variable required but not set")
	}
	// portConvertedToInt := int(port)
	conf[pghost] = host
	conf[pgport] = port
	conf[pguser] = user
	conf[pgpw] = password
	conf[pgdb] = name
	return conf
}
