
package models

type User struct {
    ID        uint   `json:"id" gorm:"primary_key"`
    Username  string `json:"username"`
    Email     string `json:"email"`
    Password  string `json:"password"`
}
