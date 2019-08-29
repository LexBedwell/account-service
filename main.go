package main

import (
    "github.com/lexbedwell/account-service/internal/adapters/transport/handlers"
    "github.com/lexbedwell/account-service/internal/adapters/database"
    "github.com/lexbedwell/account-service/internal/usecase/service"
    "log"
    "net/http"
    "os"
)

func main() {
    log.Println("starting account-service")

    var svc service.AccountService
    DAO := database.NewDao()
    svc = service.AccountService{DAO: DAO}

    go database.SyncAndSeed()

    http.Handle("/", handlers.NewGetPongFromPingHandler(svc))
    http.Handle("/ping", handlers.NewGetPongFromPingHandler(svc))
    http.Handle("/user/", handlers.NewGetInfoFromIdHandler(svc))
    http.Handle("/create", handlers.NewPostUserHandler(svc))
    
    PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
    }
    
    http.ListenAndServe(":" + PORT, nil)
    log.Println("Now listening on port", PORT)
}
