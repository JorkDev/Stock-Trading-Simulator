
package routes

import (
    "github.com/gorilla/mux"
    "your_project_path/handlers"
)

func RegisterUserRoutes(router *mux.Router) {
    router.HandleFunc("/register", handlers.Register).Methods("POST")
    router.HandleFunc("/login", handlers.Login).Methods("POST")
}
