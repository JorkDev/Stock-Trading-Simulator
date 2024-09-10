
package models

type Portfolio struct {
    ID          uint    `json:"id" gorm:"primary_key"`
    UserID      uint    `json:"user_id"`
    CashBalance float64 `json:"cash_balance"`
}
