package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

//var myEnv map[string]string = utilities.LoadEnvVaribles()
var DB *sql.DB

// load env variables and store in the variables

var (
	DB_DRIVER   = os.Getenv("DB_DRIVER")
	DB_USER     = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_NAME     = os.Getenv("DB_NAME")
	DB_SERVER   = os.Getenv("DB_SERVER")
	DB_PORT     = os.Getenv("DB_PORT")
)

func GetDbConn() *sql.DB {

	connectDB()
	// return the connection
	return DB
}

func connectDB() {
	fmt.Println(DB_DRIVER, DB_USER+":"+DB_PASSWORD+"@tcp("+DB_SERVER+":"+DB_PORT+")/"+DB_NAME)
	db, err := sql.Open(DB_DRIVER, DB_USER+":"+DB_PASSWORD+"@tcp("+DB_SERVER+":"+DB_PORT+")/"+DB_NAME)

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// check the connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	DB = db

	log.Println("Successfully connected with database!")
}

func CloseDBConnection(db *sql.DB) {
	db.Close()
	log.Println("Successfully closed DB connection")
}
