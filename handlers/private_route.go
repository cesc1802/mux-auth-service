package handlers

import (
	"auth-service/common"
	"auth-service/component/appctx"
	"auth-service/modules/auth_v1/transports/mux_auth"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func PrivateRoute(ac appctx.AppContext, router *mux.Router) {

	heath := router.PathPrefix("/heath").Subrouter()
	{
		heath.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(common.SimpleSuccessResponse(true))
			return
		}).Methods("GET")
	}
	authV1 := router.PathPrefix("/auth").Subrouter()
	{
		authV1.HandleFunc("/login", mux_auth.UserLogin(ac)).Methods("POST")
	}
}

func PrivateRouteV2(ac appctx.AppContext, router *mux.Router) {
	authV1 := router.PathPrefix("/auth").Subrouter()
	{
		authV1.HandleFunc("/login", mux_auth.UserLogin(ac)).Methods("POST")
		authV1.HandleFunc("/register", mux_auth.UserRegister(ac)).Methods("POST")
	}
}
