package config

type Config struct {
	Database *Database `mapstructure:"database"`
	Logger   *Logger   `mapstructure:"logger"`
	Auth     *Auth     `mapstructure:"auth"`
}
