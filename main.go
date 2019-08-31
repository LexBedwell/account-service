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

    go database.SyncAndSeed()

    var svc service.AccountService
    dao := database.NewDao()
    svc = service.AccountService{Dao: dao}

    http.Handle("/", handlers.NewGetPongFromPingHandler(svc))
    http.Handle("/ping", handlers.NewGetPongFromPingHandler(svc))
    http.Handle("/user/", handlers.NewGetInfoFromIdHandler(svc))
    http.Handle("/create", handlers.NewPostUserHandler(svc))
    
    PORT := os.Getenv("PORT")
    if PORT == "" {
        PORT = "8080"
    }

    log.Println("Listening on port", PORT)
    http.ListenAndServe(":" + PORT, nil)
}
