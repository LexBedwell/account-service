package main

import (
    "github.com/lexbedwell/account-service/internal/adapters/transport/handlers"
    "github.com/lexbedwell/account-service/internal/adapters/database"
    "github.com/lexbedwell/account-service/internal/usecase/service"
    "github.com/lexbedwell/account-service/internal/logging"
    "github.com/lexbedwell/account-service/internal/instrumentation"
    promhttp "github.com/prometheus/client_golang/prometheus/promhttp"
    "log"
    "net/http"
    "os"
)

func main() {
    log.Println("Starting account-service")

    go database.SyncAndSeed()

    var svc service.AccountServiceInterface
    dao := database.NewDao()
    svc = &service.AccountService{Dao: dao}
    svc = logging.NewLoggingMiddleware(svc)
    svc = instrumentation.NewInstrumentingMiddleware(svc)

    http.Handle("/ping", handlers.NewGetPongFromPingHandler(svc))
    http.Handle("/user/", handlers.NewGetInfoFromIdHandler(svc))
    http.Handle("/create", handlers.NewPostUserHandler(svc))

    http.Handle("/metrics", handlers.NewPrometheusHandler(promhttp.Handler()))
    
    PORT := os.Getenv("PORT")
    if PORT == "" {
        PORT = "8080"
    }

    log.Println("Listening on port", PORT)
    http.ListenAndServe(":" + PORT, nil)
}
