package middleware

import (
	"github.com/rs/cors"
	"net/http"
)

func AllowCors(next http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{http.MethodPost, http.MethodGet, http.MethodPut, http.MethodDelete, http.MethodOptions},
	})
	return c.Handler(next)
}
