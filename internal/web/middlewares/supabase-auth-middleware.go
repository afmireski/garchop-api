package middlewares

import (
	"context"
	"encoding/json"
	"net/http"

	customErrors "github.com/afmireski/garchop-api/internal/errors"
	"github.com/nedpals/supabase-go"
)

func SupabaseAuthMiddleware(supabaseClient *supabase.Client) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			headers := r.Header
			authorizationToken := headers.Get("Authorization")
	
			if authorizationToken == "" {
				w.Header().Set("Content-Type", "application/json")
				err := customErrors.NewInternalError("Please, provide a valid authorization token", 403, []string{
					"This route is authenticated",
				})
				w.WriteHeader(err.HttpCode)
				json.NewEncoder(w).Encode(err)
				return
			}
			ctx := context.Background()
			tokenData, err := supabaseClient.Auth.User(ctx, authorizationToken)
	
			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				err := customErrors.NewInternalError("Invalid credentials", 401, []string{
					err.Error(),
				})
				w.WriteHeader(err.HttpCode)
				json.NewEncoder(w).Encode(err)
				return
			}

			r.Header.Set("UserId", tokenData.UserMetadata["id"].(string))
			r.Header.Set("UserRole", tokenData.UserMetadata["id"].(string))
			r.Header.Set("UserEmail", tokenData.UserMetadata["email"].(string))
			r.Header.Set("UserPhone", tokenData.UserMetadata["phone"].(string))
	
			next.ServeHTTP(w, r)
		})
	}
}