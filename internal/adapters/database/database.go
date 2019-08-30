package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type dao struct{}

var db *sql.DB

func NewDao() *dao {
	return &dao{}
}

func SyncAndSeed() {
	var err error

	DATABASE_URL := os.Getenv("DATABASE_URL")
	if DATABASE_URL == "" {
		DATABASE_URL = "user=postgres dbname=user_db sslmode=disable"
	}

	log.Println("Connecting to DATABASE_URL at", DATABASE_URL)
	
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

func (_ *dao) GetUserFromId(userId string) (string, error) {
	var email string
	var err error
	err = db.QueryRow(`SELECT email FROM users WHERE id=$1`, userId).Scan(&email)
	if err != nil {
		return "", err
	}
	return email, err
}

func (_ *dao) CreateUser (email string) (string, error) {
	var id string
	var err error
	_ , err = db.Exec("INSERT INTO users(email) VALUES($1)", email)
	if err != nil && err.Error() != "pq: duplicate key value violates unique constraint \"users_email_key\"" {
		return "", err
	}
	err = db.QueryRow(`SELECT id FROM users WHERE email=$1`, email).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, err
}
