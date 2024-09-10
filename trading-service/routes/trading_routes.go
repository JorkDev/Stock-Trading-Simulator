
package routes

import (
    "github.com/gorilla/mux"
    "your_project_path/handlers"
)

func RegisterTradingRoutes(router *mux.Router) {
    router.HandleFunc("/trade/buy", handlers.BuyStock).Methods("POST")
    router.HandleFunc("/trade/sell", handlers.SellStock).Methods("POST")
}
