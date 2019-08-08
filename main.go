package main

import (
    "github.com/lexbedwell/account-service/database"
    "github.com/lexbedwell/account-service/server"
    "log"
)

func main() {
    log.Println("starting account-service")
    database.SyncAndSeed()
    server.Initialize()
}
