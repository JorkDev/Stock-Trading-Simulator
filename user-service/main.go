
package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "your_project_path/models"
    "your_project_path/routes"
)

var DB *gorm.DB

func initDB() {
    var err error
    dsn := "host=localhost user=user password=password dbname=stock_simulator port=5432 sslmode=disable TimeZone=Asia/Shanghai"
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database: ", err)
    }

    DB.AutoMigrate(&models.User{}) // Auto-migrate user table
    log.Println("Database connected successfully.")
}

func main() {
    initDB() // Initialize the DB connection

    r := mux.NewRouter()
    routes.RegisterUserRoutes(r) // Register user routes

    log.Fatal(http.ListenAndServe(":8000", r)) // Listen on port 8000
}
