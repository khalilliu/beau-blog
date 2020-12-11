package config

type JWT struct {
	SigningKey string `mapstructure:"singing-key" json:"signingKey" yaml:"signing-key"`
}