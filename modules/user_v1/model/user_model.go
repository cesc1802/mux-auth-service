package model

import (
	"auth-service/entities"
	"auth-service/modules/auth_v1/dto"
)

const (
	EntityName = "User"
)

type UserModel struct {
	entities.User
}

func (UserModel) TableName() string {
	return "users"
}

func FromUserDTO(input dto.UserRegisterRequest) UserModel {
	user := entities.User{}
	if input.FirstName != nil {
		user.FirstName = input.FirstName
	}

	if input.LastName != nil {
		user.LastName = input.LastName
	}

	user.Username = input.Username

	return UserModel{
		user,
	}
}

func (u UserModel) ComparePassword(hashedPassword string) bool {
	return u.Password == hashedPassword
}

func (u UserModel) IsActive() bool {
	return u.Status == 1
}
