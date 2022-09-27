package config

type Webhook struct {
	Name    string   `mapstructure:"name"`
	URL     string   `mapstructure:"url"`
	Path    string   `mapstructure:"path"`
	Actions []string `mapstructure:"actions"`
	Secret  string   `mapstructure:"secret"`
}
