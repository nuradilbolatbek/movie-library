package handler

import (
	"VkTask/package/service"
	"context"
	"errors"
	"log"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtxKey          = "userId"
	roleCtxKey          = "role"
)

// UserIdentity is a middleware for parsing the user id from the JWT token
func UserIdentity(svc *service.AuthService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get(authorizationHeader)
			if header == "" {
				http.Error(w, "empty auth header", http.StatusUnauthorized)
				return
			}

			headerParts := strings.Split(header, " ")
			if len(headerParts) != 2 || headerParts[0] != "Bearer" {
				http.Error(w, "invalid auth header", http.StatusUnauthorized)
				return
			}

			if len(headerParts[1]) == 0 {
				http.Error(w, "token is empty", http.StatusUnauthorized)
				return
			}

			id, role, err := svc.ParseToken(headerParts[1])
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			log.Printf("Role set in context: %s", role)

			// Setting user id to the request context
			ctx := context.WithValue(r.Context(), userCtxKey, id)
			ctx = context.WithValue(ctx, roleCtxKey, role)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// GetUserID retrieves the user id from the request context
func GetUserID(r *http.Request) (int, error) {
	userId, ok := r.Context().Value(userCtxKey).(int)
	if !ok {
		return 0, errors.New("user id not found")
	}
	return userId, nil
}

func GetUserRole(r *http.Request) (string, error) {

	role, ok := r.Context().Value(roleCtxKey).(string)
	if !ok {
		return "", errors.New("user role not found")
	}
	return role, nil
}
