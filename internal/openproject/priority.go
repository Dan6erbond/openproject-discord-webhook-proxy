package openproject

type Priority struct {
	Type      string `json:"_type"`
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Position  int    `json:"position"`
	Color     string `json:"color"`
	IsDefault bool   `json:"isDefault"`
	IsActive  bool   `json:"isActive"`
	Links     struct {
		Self NamedLink `json:"self"`
	} `json:"_links"`
}
