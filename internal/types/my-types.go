package types

type Any = interface{}

type AnyMap = map[string]Any

type Where = map[string]map[string]string

type Order = map[string]OrderOptions

type OrderOptions struct {
	Ascending    *bool   `json:"ascending"`
	ForeignTable *string `json:"foreignTable"`
	NullsFirst   *bool   `json:"nullsFirst"`
}
