package middleware

import (
	"cloud-disk/core/helper"
	"net/http"
)

type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		aush := r.Header.Get("Authorization")
		if aush == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("token is required"))
			return
		}

		userClaim, err := helper.AnalyzeToken(aush)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}

		//将用户信息加到req
		r.Header.Set("UserId", string(rune(userClaim.Id))) //rune=int32 , 数字转为字符串 64=A
		next(w, r)
	}
}
