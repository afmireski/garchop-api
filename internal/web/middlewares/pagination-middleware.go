package middlewares

import (
	"encoding/json"
	"net/http"
	"strconv"

	myTypes "github.com/afmireski/garchop-api/internal/types"
	customErrors "github.com/afmireski/garchop-api/internal/errors"
)

func PaginationMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			query := r.URL.Query()
			var queryPagination myTypes.QueryPagination

			err := json.Unmarshal([]byte(query.Get("pagination")), &queryPagination);
			if err != nil {
				queryPagination = myTypes.QueryPagination{
					Limit: 10,
					Page: 1,
				}
			}

			if queryPagination.Limit <= 0 {
				err := customErrors.NewInternalError("The pagination limit should be greater than 0", 400, []string{})
				w.WriteHeader(err.HttpCode)
				json.NewEncoder(w).Encode(err)
				return
			} else if queryPagination.Page <= 0 {
				err := customErrors.NewInternalError("The pagination page should be greater than 0", 400, []string{})
				w.WriteHeader(err.HttpCode)
				json.NewEncoder(w).Encode(err)
				return
			} 


			pagination := myTypes.Pagination{
				Limit: queryPagination.Limit,
				Page: queryPagination.Page,
				Offset: (queryPagination.Page - 1) * queryPagination.Limit,
			}

			r.Header.Set("pagination-limit", strconv.Itoa(queryPagination.Limit))
			r.Header.Set("pagination-page", strconv.Itoa(queryPagination.Page))
			r.Header.Set("pagination-offset", strconv.Itoa(pagination.Offset))
	
			next.ServeHTTP(w, r)
		})
	}
}