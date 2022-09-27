package openproject

import "time"

type User struct {
	Type        string      `json:"_type"`
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	CreatedAt   time.Time   `json:"createdAt"`
	UpdatedAt   time.Time   `json:"updatedAt"`
	Login       string      `json:"login"`
	Admin       bool        `json:"admin"`
	FirstName   string      `json:"firstName"`
	LastName    string      `json:"lastName"`
	Email       string      `json:"email"`
	Avatar      string      `json:"avatar"`
	Status      string      `json:"status"`
	IdentityURL interface{} `json:"identityUrl"`
	Language    string      `json:"language"`
	Links       struct {
		Self              NamedLink    `json:"self"`
		Memberships       NamedLink    `json:"memberships"`
		ShowUser          ResourceLink `json:"showUser"`
		UpdateImmediately ApiLink      `json:"updateImmediately"`
		Lock              ApiLink      `json:"lock"`
		Delete            ApiLink      `json:"delete"`
		AuthSource        NamedLink    `json:"auth_source"`
	} `json:"_links"`
}
