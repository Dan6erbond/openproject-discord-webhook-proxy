package openproject

import "time"

type WorkPackageWebhookPayload struct {
	WebhookPayload
	WorkPackage WorkPackage `json:"work_package"`
}

type WorkPackage struct {
	Type                 string          `json:"_type"`
	ID                   int             `json:"id"`
	LockVersion          int             `json:"lockVersion"`
	Subject              string          `json:"subject"`
	Description          FormattableText `json:"description"`
	ScheduleManually     bool            `json:"scheduleManually"`
	StartDate            interface{}     `json:"startDate"`
	DueDate              interface{}     `json:"dueDate"`
	DerivedStartDate     interface{}     `json:"derivedStartDate"`
	DerivedDueDate       interface{}     `json:"derivedDueDate"`
	EstimatedTime        string          `json:"estimatedTime"`
	DerivedEstimatedTime interface{}     `json:"derivedEstimatedTime"`
	SpentTime            string          `json:"spentTime"`
	PercentageDone       int             `json:"percentageDone"`
	CreatedAt            time.Time       `json:"createdAt"`
	UpdatedAt            time.Time       `json:"updatedAt"`
	LaborCosts           string          `json:"laborCosts"`
	MaterialCosts        string          `json:"materialCosts"`
	OverallCosts         string          `json:"overallCosts"`
	RemainingTime        string          `json:"remainingTime"`
	Embedded             struct {
		Attachments   Collection[interface{}] `json:"attachments"`
		FileLinks     Collection[interface{}] `json:"fileLinks"`
		Relations     Collection[interface{}] `json:"relations"`
		Type          Type                    `json:"type"`
		Priority      Priority                `json:"priority"`
		Project       Project                 `json:"project"`
		Status        Status                  `json:"status"`
		Author        User                    `json:"author"`
		Responsible   User                    `json:"responsible"`
		Assignee      User                    `json:"assignee"`
		CustomActions []interface{}           `json:"customActions"`
		CostsByType   Collection[interface{}] `json:"costsByType"`
	} `json:"_embedded"`
	Links struct {
		Attachments                 Link         `json:"attachments"`
		AddAttachment               ApiLink      `json:"addAttachment"`
		FileLinks                   Link         `json:"fileLinks"`
		AddFileLink                 ApiLink      `json:"addFileLink"`
		Self                        NamedLink    `json:"self"`
		Update                      ApiLink      `json:"update"`
		Schema                      Link         `json:"schema"`
		UpdateImmediately           ApiLink      `json:"updateImmediately"`
		Delete                      ApiLink      `json:"delete"`
		LogTime                     NamedLink    `json:"logTime"`
		Move                        ResourceLink `json:"move"`
		Copy                        NamedLink    `json:"copy"`
		Pdf                         ResourceLink `json:"pdf"`
		Atom                        ResourceLink `json:"atom"`
		AvailableRelationCandidates NamedLink    `json:"availableRelationCandidates"`
		CustomFields                ResourceLink `json:"customFields"`
		ConfigureForm               ResourceLink `json:"configureForm"`
		Activities                  Link         `json:"activities"`
		AvailableWatchers           Link         `json:"availableWatchers"`
		Relations                   Link         `json:"relations"`
		Revisions                   Link         `json:"revisions"`
		Watchers                    Link         `json:"watchers"`
		AddWatcher                  PayloadLink[struct {
			User ApiLink `json:"user"`
		}] `json:"addWatcher"`
		RemoveWatcher      PayloadLink[interface{}] `json:"removeWatcher"`
		AddRelation        ApiLink                  `json:"addRelation"`
		AddChild           ApiLink                  `json:"addChild"`
		ChangeParent       ApiLink                  `json:"changeParent"`
		AddComment         ApiLink                  `json:"addComment"`
		PreviewMarkup      ApiLink                  `json:"previewMarkup"`
		TimeEntries        NamedLink                `json:"timeEntries"`
		Ancestors          []NamedLink              `json:"ancestors"`
		Category           Link                     `json:"category"`
		Type               NamedLink                `json:"type"`
		Priority           NamedLink                `json:"priority"`
		Project            NamedLink                `json:"project"`
		Status             NamedLink                `json:"status"`
		Author             NamedLink                `json:"author"`
		Responsible        NamedLink                `json:"responsible"`
		Assignee           NamedLink                `json:"assignee"`
		Version            Link                     `json:"version"`
		Parent             NamedLink                `json:"parent"`
		CustomActions      []interface{}            `json:"customActions"`
		LogCosts           ResourceLink             `json:"logCosts"`
		ShowCosts          ResourceLink             `json:"showCosts"`
		CostsByType        Link                     `json:"costsByType"`
		Github             NamedLink                `json:"github"`
		GithubPullRequests NamedLink                `json:"github_pull_requests"`
		ConvertBCF         PayloadLink[struct {
			ReferenceLinks []string `json:"reference_links"`
		}] `json:"convertBCF"`
	} `json:"_links"`
}
