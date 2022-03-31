package mux_auth

import (
	"auth-service/common"
	"auth-service/component/appctx"
	"auth-service/component/tokenprovider/jwt"
	"auth-service/modules/auth_v1/biz"
	"auth-service/modules/auth_v1/dto"
	userstorage "auth-service/modules/user_v1/storage"
	"auth-service/pkg/hash"
	"encoding/json"
	"net/http"
)

func UserLogin(ac appctx.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input dto.UserLoginRequest

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			panic(common.NewCustomError(err, err.Error(), "ERR_BAD_REQUEST"))
		}

		mainDB := ac.GetMainDBConnection()
		tokenProvider := jwt.NewTokenJWTProvider(ac.SecretKey())
		md5Hash := hash.NewMd5Hash()

		userStore := userstorage.NewSQLStore(mainDB)
		authBiz := biz.NewUserLogin(userStore, md5Hash, tokenProvider)

		data, err := authBiz.UserLogin(r.Context(), &input)
		if err != nil {
			panic(common.NewCustomError(err, "", ""))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(common.SimpleSuccessResponse(data))
	}
}
