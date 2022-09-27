package openproject

import "time"

type Type struct {
	Type        string    `json:"_type"`
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Color       string    `json:"color"`
	Position    int       `json:"position"`
	IsDefault   bool      `json:"isDefault"`
	IsMilestone bool      `json:"isMilestone"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Links       struct {
		Self NamedLink `json:"self"`
	} `json:"_links"`
}
