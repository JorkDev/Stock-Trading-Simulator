package handlers

import (
    "encoding/json"
    "net/http"
    "your_project_path/models"
)

func Register(w http.ResponseWriter, r *http.Request) {
    var user models.User
    _ = json.NewDecoder(r.Body).Decode(&user)

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}

func Login(w http.ResponseWriter, r *http.Request) {
    var user models.User
    _ = json.NewDecoder(r.Body).Decode(&user)

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode("Logged in")
}
