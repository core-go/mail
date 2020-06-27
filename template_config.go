package mail

type TemplateConfig struct {
	Subject string `mapstructure:"subject"`
	Body    string `mapstructure:"body"`
}
