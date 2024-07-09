package middlewares

import (
	"encoding/json"
	"net/http"

	customErrors "github.com/afmireski/garchop-api/internal/errors"
	"github.com/afmireski/garchop-api/internal/types/enums"
)

func UserRoleMiddleware(expectedRole enums.UserRoleEnum) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			headers := r.Header
			role := headers.Get("User-Role")
	
			if role != string(expectedRole) {
				w.Header().Set("Content-Type", "application/json")
				err := customErrors.NewInternalError("You cannot have necessary roles for access this resources", 403, []string{})
				w.WriteHeader(err.HttpCode)
				json.NewEncoder(w).Encode(err)
				return
			}
	
			next.ServeHTTP(w, r)
		})
	}
}