package main

import (
	"akimbaev/database"
	"akimbaev/routes"
	"log"
	"net/http"
)

var secretKey = []byte("secret-key")

func main() {
	database.Init()
	api := http.NewServeMux()
	api.Handle("/api/v1/user/", http.StripPrefix("/api/v1/user", routes.UserMux()))
	log.Println("Listening on :8080")

	http.ListenAndServe(":8080", api)
}

// func CreateUser(w http.ResponseWriter, r *http.Request) {
// 	type params struct {
// 		Name     string `json:"name"`
// 		Email    string `json:"email"`
// 		Password string `json:"password"`
// 	}

// 	decoder := json.NewDecoder(r.Body)

// 	p := params{}
// 	err := decoder.Decode(&p)

// 	if err != nil {
// 		response.Json(w, http.StatusBadRequest, "Invalid JSON")
// 		fmt.Println("Decoding error:", err)
// 		return
// 	}

// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)

// 	if err != nil {
// 		response.Json(w, http.StatusInternalServerError, "Error while hashing password")
// 	}

// 	NewUser := models.User{
// 		Email:    p.Email,
// 		Name:     p.Name,
// 		Password: string(hashedPassword),
// 	}

// 	database.DB.Create(&NewUser)

// 	response.Json(w, http.StatusOK, NewUser)
// }
