package api

import (
	"go-hello/utils"
	"net/http"
	"strings"
)

type AuthMiddleware struct{}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (s *AuthMiddleware) AuthLoginMiddle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.EncodeJSON(w, http.StatusOK, utils.H{
				"message": "请登录",
			})
			http.Redirect(w, r, "/user/login", http.StatusFound)
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			utils.EncodeJSON(w, http.StatusOK, utils.H{
				"message": "无效的token",
			})
			http.Redirect(w, r, "/user/login", http.StatusFound)
			return
		}
		if !PasreJWT(parts[1]) {
			utils.EncodeJSON(w, http.StatusOK, utils.H{
				"message": "无效的token",
			})
			http.Redirect(w, r, "user/login", http.StatusFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}
