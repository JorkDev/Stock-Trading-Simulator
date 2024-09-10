
package models

import "time"

type Order struct {
    ID          uint      `json:"id" gorm:"primary_key"`
    UserID      uint      `json:"user_id"`
    StockSymbol string    `json:"stock_symbol"`
    OrderType   string    `json:"order_type"` // buy or sell
    Quantity    int       `json:"quantity"`
    Price       float64   `json:"price"`
    CreatedAt   time.Time `json:"created_at"`
}
