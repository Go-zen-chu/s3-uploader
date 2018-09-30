package server

type ConfigRepository interface {
	LoadFromEnv() *Config
}