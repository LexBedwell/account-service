package main

import (
    "github.com/lexbedwell/account-service/internal/adapters/transport/handlers"
    "github.com/lexbedwell/account-service/internal/adapters/database"
    "github.com/lexbedwell/account-service/internal/usecase/service"
    "github.com/lexbedwell/account-service/internal/logging"
    gokitlog "github.com/go-kit/kit/log"
    "log"
    "net/http"
    "os"
)

func main() {
    log.Println("Starting account-service")

    go database.SyncAndSeed()

    logger := gokitlog.NewLogfmtLogger(os.Stderr)

    //set svc as service interface
    var svc service.AccountServiceInterface
    dao := database.NewDao()
    //set svc as type that implements that service interface (has the same methods)
    svc = &service.AccountService{Dao: dao}
    //set svc as another type (again!) that also implements that interface (has the same methods) 
    //whose implementation methods call next to the methods on the preceeding wrapped layers
    svc = &logging.LoggingMiddleware{Logger: logger, Next: svc}

    http.Handle("/ping", handlers.NewGetPongFromPingHandler(svc))
    http.Handle("/user/", handlers.NewGetInfoFromIdHandler(svc))
    http.Handle("/create", handlers.NewPostUserHandler(svc))

    http.Handle("/metrics", handlers.NewPrometheusHandler())
    
    PORT := os.Getenv("PORT")
    if PORT == "" {
        PORT = "8080"
    }

    log.Println("Listening on port", PORT)
    http.ListenAndServe(":" + PORT, nil)
}
