package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var db *sql.DB

func SyncAndSeed() {
	var err error

	DATABASE_URL := os.Getenv("DATABSE_URL")
	if DATABASE_URL == "" {
		DATABASE_URL = "user=postgres dbname=user_db sslmode=disable"
	}
	
	db, err = sql.Open("postgres", DATABASE_URL)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error: Could not establish a connection with the database")
	}

	const createUserTableQry = `CREATE TABLE IF NOT EXISTS users (
		id serial PRIMARY KEY,
		email VARCHAR (40) UNIQUE NOT NULL
		);`

	_, err = db.Exec(createUserTableQry)
	if err != nil {
		log.Fatal(err)
	}

	const createUserQry = `INSERT INTO users(email) 
		VALUES ('lexbedwell@gmail.com')
		ON CONFLICT (email) 
		DO NOTHING;`

	_, err = db.Exec(createUserQry)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("database sync and seed completed")

}

func GetUserFromId(userId string) (string, error) {
	var email string
	var err error
	err = db.QueryRow(`SELECT email FROM users WHERE id=$1`, userId).Scan(&email)
	if err != nil {
		return "", err
	}
	return email, err
}

func CreateUser (email string) error {
	var err error
	_ , err = db.Exec("INSERT INTO users(email) VALUES($1)", email)
	return err
}
