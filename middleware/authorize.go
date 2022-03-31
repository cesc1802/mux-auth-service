package middleware

import (
	"auth-service/common"
	"auth-service/component/appctx"
	"auth-service/component/tokenprovider/jwt"
	usermodel "auth-service/entities"
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type AuthStore interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

func ErrWrongAuthHeader(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("wrong authen header"),
		fmt.Sprintf("ErrWrongAuthHeader"),
	)
}

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")
	//"Authorization" : "Bearer {token}"

	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(nil)
	}

	return parts[1], nil
}

func muxExtractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")
	//"Authorization" : "Bearer {token}"

	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(nil)
	}

	return parts[1], nil
}

func MuxRequiredAuth(appCtx appctx.AppContext, authStore AuthStore) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())
			token, err := muxExtractTokenFromHeaderString(r.Header.Get("Authorization"))
			if err != nil {
				panic(err)
			}

			payload, err := tokenProvider.Validate(token)
			if err != nil {
				panic(err)
			}

			user, err := authStore.FindUser(r.Context(), map[string]interface{}{"id": payload.UserId})

			if err != nil {
				panic(err)
			}

			if user.Status != 1 {
				panic(common.ErrNoPermission(errors.New("user has been deleted or banned")))
			}

			user.Mask(common.DbTypeUser)
			ctxWithUser := context.WithValue(r.Context(), common.CurrentUser, user)

			next.ServeHTTP(w, r.WithContext(ctxWithUser))
		})
	}
}
