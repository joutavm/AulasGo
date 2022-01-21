package server

import (
	"crud/database"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type user struct {
	ID    uint32 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateUser(writter http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)

	if err != nil {
		writter.WriteHeader(http.StatusBadRequest)
		writter.Write([]byte("Error reading body"))
		return
	}

	var userRequest user
	if err = json.Unmarshal(body, &userRequest); err != nil {
		writter.WriteHeader(http.StatusBadRequest)
		writter.Write([]byte("Error parsing body"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		writter.WriteHeader(http.StatusInternalServerError)
		log.Println("Failed to connect in db", err)
		return
	}
	defer db.Close()

	var response user

	if err = db.QueryRow("INSERT INTO person (Name, Email) VALUES ($1, $2) RETURNING ID, name, email",
		userRequest.Name, userRequest.Email).Scan(&response.ID, &response.Name, &response.Email); err != nil {
		writter.WriteHeader(http.StatusInternalServerError)
		log.Println("Failed to insert user", err)

		return
	}

	responseBytes, err := json.Marshal(response)
	if err != nil {
		writter.WriteHeader(http.StatusInternalServerError)
		log.Println("Failed to parse response", err)

		return
	}
	writter.Header().Set("Content-Type", "application/json")
	writter.WriteHeader(http.StatusCreated)
	writter.Write(responseBytes)
}

func GetUsers(writter http.ResponseWriter, request *http.Request) {

}
