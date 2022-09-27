package openproject

type FormattableText struct {
	Format string `json:"format"`
	Raw    string `json:"raw"`
	HTML   string `json:"html"`
}
