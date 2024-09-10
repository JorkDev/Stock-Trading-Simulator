
package main

import (
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
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

    DB.AutoMigrate(&models.User{}, &models.Order{}, &models.Portfolio{}) // Auto-migrate models
    log.Println("Database connected and models migrated successfully.")
}

func main() {
    initDB()

    r := mux.NewRouter()
    routes.RegisterTradingRoutes(r) // Register trading routes

    // Start HTTP server in a goroutine
    srv := &http.Server{
        Handler: r,
        Addr:    ":8001",
    }

    go func() {
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("ListenAndServe(): %s", err)
        }
    }()
    log.Println("Trading Service HTTP server started on :8001")

    // Set up Kafka consumer in another goroutine
    go consumeStockPrices()

    // Graceful shutdown
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    log.Println("Shutting down Trading Service...")
}
