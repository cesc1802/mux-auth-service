package entities

import "auth-service/common"

type User struct {
	common.SQLModel
	Username  string  `gorm:"column:username"`
	Password  string  `gorm:"column:password"`
	FirstName *string `gorm:"column:first_name"`
	LastName  *string `gorm:"column:last_name"`
	Salt      string  `gorm:"column:salt"`
}
