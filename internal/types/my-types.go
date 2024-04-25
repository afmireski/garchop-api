package types

type Any = interface{}

type AnyMap = map[string]Any

type Where = map[string]map[string]string

type Order = map[string]OrderOptions

type OrderOptions struct {
	Ascending bool    `json:"ascending"`
	Child     *string `json:"child"`
	Field     string  `json:"field"`
}
