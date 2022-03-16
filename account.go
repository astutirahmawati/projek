package account

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `gorm:"unique" json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Phone    string `gorm:"unique" json:"phone" form:"phone"`
	Address  string `json:"address" form:"address"`
}


func CreateAccount(*newUser Account) {
	tx := DB.Save(&newUser)
}

func GetAccount() {

}

func UpdateAccount() {

}

func DeleteAccount() {

}
