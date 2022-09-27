package openproject

type Collection[T interface{}] struct {
	Type     string `json:"_type"`
	Total    int    `json:"total"`
	Count    int    `json:"count"`
	Embedded struct {
		Elements []T `json:"elements"`
	} `json:"_embedded"`
	Links struct {
		Self NamedLink `json:"self"`
	} `json:"_links"`
}
