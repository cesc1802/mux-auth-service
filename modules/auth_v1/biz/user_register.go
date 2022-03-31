package biz

import (
	"auth-service/common"
	"auth-service/modules/auth_v1/dto"
	"auth-service/modules/user_v1/model"
	"auth-service/pkg/stringutil"
	"context"
	"fmt"
)

type UserRegisterStore interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*model.UserModel, error)
	Create(ctx context.Context, data *model.UserModel) error
}

type userRegister struct {
	store  UserRegisterStore
	hasher common.Hasher
}

func NewUserRegister(store UserRegisterStore, hasher common.Hasher) *userRegister {
	return &userRegister{store: store, hasher: hasher}
}

func (biz *userRegister) UserRegister(ctx context.Context, input *dto.UserRegisterRequest) error {
	user, err := biz.store.FindUser(ctx, map[string]interface{}{
		"username": input.Username,
	})

	if user != nil {
		return common.ErrEntityExisted(model.EntityName, err)
	}

	if err != nil && err != common.ErrRecordNotFound {
		return common.ErrCannotGetEntity(model.EntityName, err)
	}

	salt := stringutil.GenSalt(50)
	hashedPassword := biz.hasher.Hash(fmt.Sprintf("%s%s", input.Password, salt))

	userModel := model.FromUserDTO(*input)
	userModel.Salt = salt
	userModel.Password = hashedPassword

	fmt.Println(userModel)
	if err := biz.store.Create(ctx, &userModel); err != nil {
		return common.ErrCannotCreateEntity(model.EntityName, err)
	}

	return nil
}
