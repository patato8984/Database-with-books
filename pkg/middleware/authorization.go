package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type Authorization struct {
	Kye string
}

func NewAuthorization(kye string) *Authorization {
	return &Authorization{Kye: kye}
}
func (a *Authorization) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			http.Error(w, `{"status":"error"}`, http.StatusUnauthorized)
			return
		}
		parts := strings.Split(auth, " ")
		if parts[0] != "Bearer" || len(parts) != 2 {
			http.Error(w, `{"status":"error"}`, http.StatusUnauthorized)
			return
		}
		tokenString := parts[1]
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("error jwt: %v", t.Header["alg"])
			}
			return []byte(a.Kye), nil
		})
		if err != nil || !token.Valid {
			http.Error(w, `{"status":"error"}`, http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
