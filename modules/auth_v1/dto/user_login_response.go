package dto

import "auth-service/component/tokenprovider"

type UserLoginResponse struct {
	AccessToken  tokenprovider.Token `json:"access_token"`
	RefreshToken tokenprovider.Token `json:"refresh_token"`
}
