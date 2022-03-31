package biz

import (
	"auth-service/common"
	"auth-service/component/tokenprovider"
	"auth-service/modules/auth_v1/dto"
	"auth-service/modules/user_v1/model"
	"context"
	"fmt"
)

type UserStore interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*model.UserModel, error)
}

type userLogin struct {
	store       UserStore
	hasher      common.Hasher
	tokProvider tokenprovider.Provider
}

func NewUserLogin(store UserStore, hasher common.Hasher, tokProvider tokenprovider.Provider) *userLogin {
	return &userLogin{store: store, hasher: hasher, tokProvider: tokProvider}
}

func (biz *userLogin) UserLogin(ctx context.Context, input *dto.UserLoginRequest) (*dto.UserLoginResponse, error) {
	user, err := biz.store.FindUser(ctx, map[string]interface{}{
		"username": input.Username,
	})

	if err != nil {
		return nil, common.ErrCannotGetEntity(model.EntityName, err)
	}

	if !user.IsActive() {
		return nil, common.NewCustomError(nil, "user has been block", "ERR_USER_HAS_BEEN_BLOCK")
	}

	hashedPassword := biz.hasher.Hash(fmt.Sprintf("%s%s", input.Password, user.Salt))
	if !user.ComparePassword(hashedPassword) {
		return nil, common.NewCustomError(nil, "username or password invalid", "ERR_USERNAME_PASSWORD")
	}

	payload := tokenprovider.TokenPayload{UserId: user.Id}
	accessToken, err := biz.tokProvider.Generate(payload, 60*60*24*7)
	if err != nil {

	}
	refreshToken, err := biz.tokProvider.Generate(payload, 60*60*24*30)
	if err != nil {

	}
	return &dto.UserLoginResponse{AccessToken: *accessToken, RefreshToken: *refreshToken}, nil
}
