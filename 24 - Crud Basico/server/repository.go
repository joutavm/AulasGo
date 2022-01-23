package server

import (
	"crud/database"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type user struct {
	ID    uint32 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateUser(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Error reading body"))
		return
	}

	var userRequest user
	if err = json.Unmarshal(body, &userRequest); err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Error parsing body"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println("Failed to connect in db", err)
		return
	}
	defer db.Close()

	var response user

	if err = db.QueryRow("INSERT INTO person (Name, Email) VALUES ($1, $2) RETURNING ID, name, email",
		userRequest.Name, userRequest.Email).Scan(&response.ID, &response.Name, &response.Email); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println("Failed to insert user", err)

		return
	}

	responseBytes, err := json.Marshal(response)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println("Failed to parse response", err)

		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	writer.Write(responseBytes)
}

func GetUsers(writer http.ResponseWriter, request *http.Request) {
	db, err := database.Connect()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println("Failed to connect in db", err)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM person")
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println("Failed to query users", err)
		return
	}
	defer rows.Close()

	var users []user

	for rows.Next() {
		var user user
		if err = rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			log.Println("Failed to scan user", err)
			return
		}
		users = append(users, user)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(users); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println("Failed to parse response", err)
		return
	}

}

func GetUser(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	ID, err := strconv.ParseUint(params["id"], 10, 32)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Error parsing id"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println("Failed to connect in db", err)
		return
	}
	defer db.Close()

	row, err := db.Query("SELECT * FROM person WHERE id = $1", ID)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println("Failed to query user", err)
		return
	}

	var user user
	if row.Next() {
		if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			log.Println("Failed to scan user", err)
			return
		}
	}

	if user.ID == 0 {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("User not found"))
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(user); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println("Failed to parse response", err)
		return
	}

}

func UpdateUser(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	ID, err := strconv.ParseUint(params["id"], 10, 32)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Error parsing id"))
		return
	}

	body, err := ioutil.ReadAll(request.Body)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Error reading body"))
		return
	}

	var userRequest user
	if err = json.Unmarshal(body, &userRequest); err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Error parsing body"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println("Failed to connect in db", err)
		return
	}
	defer db.Close()

	var response user

	if err = db.QueryRow("UPDATE person SET name = $1, email = $2 WHERE id = $3 RETURNING id, name, email",
		userRequest.Name, userRequest.Email, ID).Scan(&response.ID, &response.Name, &response.Email); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println("Failed to update user", err)

		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(response); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println("Failed to parse response", err)
		return
	}
}

func DeleteUser(writter http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	ID, err := strconv.ParseUint(params["id"], 10, 32)

	if err != nil {
		writter.WriteHeader(http.StatusBadRequest)
		writter.Write([]byte("Error parsing id"))
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

	if err = db.QueryRow("DELETE FROM person WHERE id = $1 RETURNING id, name, email", ID).Scan(&response.ID,
		&response.Name, &response.Email); err != nil {
		writter.WriteHeader(http.StatusInternalServerError)
		log.Println("Failed to delete user", err)
		return
	}

	if response.ID == 0 {
		writter.WriteHeader(http.StatusNotFound)
		writter.Write([]byte("User not found"))
		return
	}

	writter.Header().Set("Content-Type", "application/json")
	writter.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writter).Encode(response); err != nil {
		writter.WriteHeader(http.StatusInternalServerError)
		log.Println("Failed to parse response", err)
		return
	}
}
