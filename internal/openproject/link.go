package openproject

type Link struct {
	Href string `json:"href"`
}

type NamedLink struct {
	Link
	Title string `json:"title"`
}

type ApiLink struct {
	NamedLink
	Method string `json:"method"`
}

type PayloadLink[T interface{}] struct {
	ApiLink
	Payload   T    `json:"payload"`
	Templated bool `json:"templated"`
}

type ResourceLink struct {
	NamedLink
	Type string `json:"type"`
}
