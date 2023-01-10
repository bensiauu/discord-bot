package config

type Config struct {
	Authorization Auth `yaml:"authorization"`
}

type Auth struct {
	Token string `yaml:"token"`
}
