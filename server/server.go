package server

import (
	"encoding/json"
	"github.com/lexbedwell/account-service/database"
	"log"
	"net/http"
	"os"
)

type postUserRequest struct {
	Email string `json:"email"`
}

type getPingResponse struct {
	Response   string `json:"response"`
}

type getIdResponse struct {
	Id   string `json:"id"`
	Email   string `json:"email"`
	Err string `json:"error"`
}

type postUserResponse struct {
	Email   string `json:"created"`
}

func Initialize() {

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = ":8080"
	}

	log.Println("Now listening on port: ", PORT)

	http.HandleFunc("/ping", getPongFromPing)
	http.HandleFunc("/user/", getInfoFromId)
	http.HandleFunc("/create", postUser)

	http.ListenAndServe(PORT, nil)

}

func getPongFromPing(w http.ResponseWriter, _ *http.Request) {
	pingResponse := getPingResponse{"pong"}

	pingResponseJson, err := json.Marshal(pingResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(pingResponseJson)
}

func getInfoFromId(w http.ResponseWriter, r *http.Request) {
	error := ""
	ids, ok := r.URL.Query()["id"]
	if !ok || len(ids[0]) < 1 {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	info, err := database.GetUserFromId(ids[0])
	if err != nil {
		error = err.Error()
	}

	idResponse := getIdResponse{ids[0], info, error}
	idResponseJson, err := json.Marshal(idResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(idResponseJson)

}

func postUser(w http.ResponseWriter, r *http.Request) {
	var request postUserRequest
	var err error

	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil || request.Email == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	err = database.CreateUser(request.Email)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userResponse := postUserResponse{request.Email}
	userResponseJson, err := json.Marshal(userResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(userResponseJson)

}
