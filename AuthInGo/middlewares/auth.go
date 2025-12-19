package middlewares

import (
	dbConfig "AuthInGo/config/db"
	env "AuthInGo/config/env"
	repo "AuthInGo/db/repositories"
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Authorization header must start with Bearer", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		if token == "" {
			http.Error(w, "Token is required", http.StatusUnauthorized)
			return
		}

		claims := jwt.MapClaims{}

		_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(env.GetString("JWT_SECRET", "TOKEN")), nil
		})

		if err != nil {
			http.Error(w, "Invalid token:"+err.Error(), http.StatusUnauthorized)
			return
		}

		userId, okId := claims["id"].(float64)

		email, okEmail := claims["email"].(string)

		if !okId || !okEmail {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		fmt.Println("Authenticated user Id:", int64(userId), "Email:", email)

		ctx := context.WithValue(r.Context(), UserIDKey, strconv.FormatFloat(userId, 'f', 0, 64))

		ctx = context.WithValue(ctx, EmailKey, email)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func RequireAllRoles(roles ...string) func(http.Handler) http.Handler {
	// function that can create middleware to check for roles
	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			userIdStr := r.Context().Value(UserIDKey).(string)

			userId, err := strconv.ParseInt(userIdStr, 10, 64)

			if err != nil {
				http.Error(w, "Invalid user ID", http.StatusUnauthorized)
				return
			}

			dbConn, dbErr := dbConfig.SetupDB()

			if dbErr != nil {
				http.Error(w, "Database connection error:"+dbErr.Error(), http.StatusInternalServerError)
				return
			}

			urr := repo.NewUserRoleRepository(dbConn)

			hasAllRoles, err := urr.HasAllRoles(userId, roles)
			if err != nil {
				http.Error(w, "Error checking roles", http.StatusInternalServerError)
				return
			}
			if !hasAllRoles {
				http.Error(w, "Insufficient permissions", http.StatusForbidden)
				return
			}
			fmt.Println("User", userId, "has all required roles:", roles)

			next.ServeHTTP(w, r)

		})

	}
}

func RequireAnyRole(roles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userIdStr := r.Context().Value(UserIDKey).(string)
			userId, err := strconv.ParseInt(userIdStr, 10, 64)
			if err != nil {
				http.Error(w, "Invalid user ID", http.StatusUnauthorized)
				return
			}
			dbConn, dbErr := dbConfig.SetupDB()
			if dbErr != nil {
				http.Error(w, "Database connection error:"+dbErr.Error(), http.StatusInternalServerError)
				return
			}
			urr := repo.NewUserRoleRepository(dbConn)
			hasAnyRole, err := urr.HasAnyRole(userId, roles)
			fmt.Println("hasAnyRole:", hasAnyRole)
			if err != nil {
				http.Error(w, "Error checking roles", http.StatusInternalServerError)
				return
			}
			if !hasAnyRole {
				http.Error(w, "Insufficient permissions", http.StatusForbidden)
				return
			}
			fmt.Println("User", userId, "has any required role:", roles)
			next.ServeHTTP(w, r)
		})
	}
}
