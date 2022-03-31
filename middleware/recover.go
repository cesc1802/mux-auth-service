package middleware

import (
	"auth-service/common"
	"auth-service/component/appctx"
	"encoding/json"
	"net/http"
)

func Recover(ac appctx.AppContext) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					w.Header().Set("Content-Type", "application/json")
					if appErr, ok := err.(*common.AppError); ok {
						w.WriteHeader(appErr.StatusCode)
						json.NewEncoder(w).Encode(appErr)
						return
					}

					appErr := common.ErrInternal(err.(error))
					w.WriteHeader(appErr.StatusCode)
					json.NewEncoder(w).Encode(appErr)
					return
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}
