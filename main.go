package main

import (
    "log"
    "github.com/lexbedwell/account-service/database"
)

func main() {
    log.Println("starting account-service")
    database.SyncAndSeed()
}
