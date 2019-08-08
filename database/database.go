package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func SyncAndSeed() {

	db, err := sql.Open("postgres", "user=postgres dbname=user_db sslmode=disable")
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

	const createUserQry = `insert into users (id, email) 
        values (1, 'lexbedwell@gmail.com')
        ON CONFLICT (id) 
        DO NOTHING;`

	_, err = db.Exec(createUserQry)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("database sync and seed completed")

}