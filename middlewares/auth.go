package middlewares

import (
	"context"
	"net/http"
	"task5-vix/helpers"
)

func Auth( next http.Handler ) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		accesToken := r.Header.Get("Authorization")

		if accesToken == "" {
			helpers.Response(w, 401, "Unauthorized", nil)
			return
		}

		user, err := helpers.ValidateToken(accesToken)
		if err != nil {
			helpers.Response(w, 401, err.Error(), nil)
			return
		}
		

		ctx := context.WithValue(r.Context(),"userinfo", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}