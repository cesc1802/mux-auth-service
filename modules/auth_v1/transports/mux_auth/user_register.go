package mux_auth

import (
	"auth-service/common"
	"auth-service/component/appctx"
	"auth-service/modules/auth_v1/biz"
	"auth-service/modules/auth_v1/dto"
	userstorage "auth-service/modules/user_v1/storage"
	"auth-service/pkg/hash"
	"encoding/json"
	"net/http"
)

func UserRegister(ac appctx.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input dto.UserRegisterRequest

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			panic(common.NewCustomError(err, err.Error(), "ERR_BAD_REQUEST"))
		}

		mainDB := ac.GetMainDBConnection()
		userStore := userstorage.NewSQLStore(mainDB)
		md5Hash := hash.NewMd5Hash()
		biz := biz.NewUserRegister(userStore, md5Hash)
		if err := biz.UserRegister(r.Context(), &input); err != nil {
			panic(common.NewCustomError(err, "", "ERR_BAD_REQUEST"))
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(common.SimpleSuccessResponse(true))
	}
}
