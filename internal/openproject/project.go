package openproject

import "time"

type Project struct {
	Type              string          `json:"_type"`
	ID                int             `json:"id"`
	Identifier        string          `json:"identifier"`
	Name              string          `json:"name"`
	Active            bool            `json:"active"`
	Public            bool            `json:"public"`
	Description       FormattableText `json:"description"`
	CreatedAt         time.Time       `json:"createdAt"`
	UpdatedAt         time.Time       `json:"updatedAt"`
	StatusExplanation FormattableText `json:"statusExplanation"`
	Links             struct {
		Self                         NamedLink     `json:"self"`
		CreateWorkPackage            ApiLink       `json:"createWorkPackage"`
		CreateWorkPackageImmediately ApiLink       `json:"createWorkPackageImmediately"`
		WorkPackages                 ApiLink       `json:"workPackages"`
		Storages                     []interface{} `json:"storages"`
		Categories                   ApiLink       `json:"categories"`
		Versions                     ApiLink       `json:"versions"`
		Memberships                  ApiLink       `json:"memberships"`
		Types                        ApiLink       `json:"types"`
		Update                       ApiLink       `json:"update"`
		UpdateImmediately            ApiLink       `json:"updateImmediately"`
		Delete                       ApiLink       `json:"delete"`
		Schema                       ApiLink       `json:"schema"`
		Ancestors                    []interface{} `json:"ancestors"`
		Parent                       struct {
			Href interface{} `json:"href"`
		} `json:"parent"`
		Status NamedLink `json:"status"`
	} `json:"_links"`
}
