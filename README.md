# Account Service

Service that stores and updates customer account information for [BuyAndCell](https://github.com/LexBedwell/BuyAndCell).
Utilizes Go and an accompanying PostgreSQL database.

## Commands
- `go run main.go` - start service

## Routes

| Method | Route | Call Definition | Description
| ------ | ----- | ----- | -----------
| GET | /metrics | handlers.NewPrometheusHandler | Metrics for Prometheus
| GET | /ping | handlers.NewGetPongFromPingHandler | Ping service
| GET | /user/?id=$userId | handlers.NewGetInfoFromIdHandler | Get account info from userId
| POST | /create | handlers.NewPostUserHandler | Create user
