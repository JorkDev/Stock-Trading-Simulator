package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "your_project_path/routes"
)

func main() {
    r := mux.NewRouter()
    routes.RegisterUserRoutes(r)

    log.Fatal(http.ListenAndServe(":8000", r))
}
