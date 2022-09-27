package openproject

type Status struct {
	Type             string      `json:"_type"`
	ID               int         `json:"id"`
	Name             string      `json:"name"`
	IsClosed         bool        `json:"isClosed"`
	Color            string      `json:"color"`
	IsDefault        bool        `json:"isDefault"`
	IsReadonly       bool        `json:"isReadonly"`
	DefaultDoneRatio interface{} `json:"defaultDoneRatio"`
	Position         int         `json:"position"`
	Links            struct {
		Self NamedLink `json:"self"`
	} `json:"_links"`
}
