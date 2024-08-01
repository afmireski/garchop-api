package helpers

import (
	"encoding/json"
	"net/url"

	customErrors "github.com/afmireski/garchop-api/internal/errors"
)

func ParseQueryParam[T any](query url.Values, paramName string, param *T)  *customErrors.InternalError {
	queryParam := query.Get(paramName)

	if len(queryParam) == 0 {
		return nil
	}

	parseErr := json.Unmarshal([]byte(queryParam), &param)
	if parseErr != nil {
		return customErrors.NewInternalError("fail on deserialize query parameter", 400, []string{parseErr.Error()})
	}

	return nil
}