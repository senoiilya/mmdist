package middlewares

import (
	"context"
	"net/http"

	"github.com/senoiilya/mmdist/utils"
)

func IsAuthorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		claims, err := utils.ParseToken(cookie.Value)
		if err != nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		// Сохранение роли в контексте запроса
		ctx := context.WithValue(r.Context(), "role", claims.Role)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
