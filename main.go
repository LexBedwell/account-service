package main

import (
    "database/sql"
    _ "github.com/lib/pq"
    "fmt"
    "log"    
)

func main() {
    fmt.Println("app started")

    db, err := sql.Open("postgres", "user=postgres dbname=user_db sslmode=disable")
    if err != nil {
      log.Fatal(err)
    }

    err = db.Ping()
    if err != nil {
        log.Fatal("Error Could not establish a connection with the database")
    }

    var lastName string
    err = db.QueryRow("SELECT last_name FROM users WHERE id=$1", 15).Scan(&lastName)
    if err == sql.ErrNoRows {
        log.Fatal("No Results Found")
    }
    if err != nil {
        log.Fatal(err)
    }

}