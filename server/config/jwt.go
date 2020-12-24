package config

type JWT struct {
	SigningKey string `mapstructure:"singing-key" json:"signingKey" yaml:"signing-key"`
	ExpiresTime int64  `mapstructure:"expires-time" json:"expiresTime" yaml:"expires-time"`
	BufferTime  int64  `mapstructure:"buffer-time" json:"bufferTime" yaml:"buffer-time"`
}