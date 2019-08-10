# Account Service

Service that stores and updates customer account information for [BuyAndCell](https://github.com/LexBedwell/BuyAndCell).
Utilizes Go and an accompanying PostgreSQL database.

## Commands
- `go run main.go` - start service

## Routes

| Method | Route | Call Definition | Description
| ------ | ----- | ----- | -----------
| GET | /ping | server.getPongFromPing | Ping service
| GET | /user/?id=$userId | server.getInfoFromId | Get user info from userId
| POST | /create | server.postUser | Create user
