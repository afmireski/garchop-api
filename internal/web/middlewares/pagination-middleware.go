package middlewares

import (
	"context"
	"encoding/json"
	"net/http"

	customErrors "github.com/afmireski/garchop-api/internal/errors"
	myTypes "github.com/afmireski/garchop-api/internal/types"
)

func PaginationMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			query := r.URL.Query()
			var queryPagination myTypes.QueryPagination
			err := json.Unmarshal([]byte(query.Get("pagination")), &pagination);
			if err != nil {
				pagination = myTypes.Pagination{
					Limit: 10,
					Page: 1,
					Offset: 0,
				}
			}


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
			_, err := supabaseClient.Auth.User(ctx, authorizationToken)
	
			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				err := customErrors.NewInternalError("Invalid credentials", 401, []string{
					err.Error(),
				})
				w.WriteHeader(err.HttpCode)
				json.NewEncoder(w).Encode(err)
				return
			}
	
			next.ServeHTTP(w, r)
		})
	}
}