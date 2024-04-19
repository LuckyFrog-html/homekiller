package middlewares

import (
	"context"
	"fmt"
	"github.com/go-chi/jwtauth"
	"net/http"
)

func JWTAuthHolder(tokenAuth *jwtauth.JWTAuth) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		handler := func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = context.WithValue(ctx, "jwt_auth", tokenAuth)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(handler)
	}
}

func TeacherAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		_, claims, err := jwtauth.FromContext(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}
		if claims["table"] == nil || claims["table"] != "teachers" {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}

		ctx = context.WithValue(ctx, "teacher_id", claims["id"])

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetTeacherIdFromContext(ctx context.Context) (uint, error) {
	val := ctx.Value("teacher_id")
	if val == nil {
		return 0, fmt.Errorf("can't find teacher_id value")
	}
	res, ok := val.(float64)
	if !ok {
		return 0, fmt.Errorf("can't use ctx value as uint")
	}
	return uint(res), nil
}

func GetAuthTokenFromContext(ctx context.Context) (*jwtauth.JWTAuth, error) {
	val := ctx.Value("jwt_auth")
	if val == nil {
		return nil, fmt.Errorf("can't find jwt_auth value")
	}
	res, ok := val.(*jwtauth.JWTAuth)
	if !ok {
		return nil, fmt.Errorf("can't use ctx value as jwtauth.JWTAuth")
	}
	return res, nil
}
