
package handlers

import (
    "encoding/json"
    "net/http"
    "your_project_path/models"
    "your_project_path/main"
    "golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
    var user models.User
    _ = json.NewDecoder(r.Body).Decode(&user)

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
    if err != nil {
        http.Error(w, "Error hashing password", http.StatusInternalServerError)
        return
    }

    user.Password = string(hashedPassword)

    result := main.DB.Create(&user) // Save user to the database
    if result.Error != nil {
        http.Error(w, "Could not create user", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}

func Login(w http.ResponseWriter, r *http.Request) {
    var user models.User
    _ = json.NewDecoder(r.Body).Decode(&user)

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode("Logged in")
}
