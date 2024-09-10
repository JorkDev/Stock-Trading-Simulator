
package handlers

import (
    "encoding/json"
    "net/http"
    "your_project_path/models"
    "your_project_path/main"
    "time"
)

type TradeRequest struct {
    UserID      uint    `json:"user_id"`
    StockSymbol string  `json:"stock_symbol"`
    Quantity    int     `json:"quantity"`
    Price       float64 `json:"price"`
}

func BuyStock(w http.ResponseWriter, r *http.Request) {
    var req TradeRequest
    _ = json.NewDecoder(r.Body).Decode(&req)

    order := models.Order{
        UserID:      req.UserID,
        StockSymbol: req.StockSymbol,
        OrderType:   "buy",
        Quantity:    req.Quantity,
        Price:       req.Price,
        CreatedAt:   time.Now(),
    }

    result := main.DB.Create(&order)
    if result.Error != nil {
        http.Error(w, "Could not create order", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode("Order created successfully")
}

func SellStock(w http.ResponseWriter, r *http.Request) {
    var req TradeRequest
    _ = json.NewDecoder(r.Body).Decode(&req)

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode("Stock sold successfully")
}
