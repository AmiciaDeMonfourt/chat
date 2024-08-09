package config

import (
	"flag"
	"log"
	"log/slog"
	"os"

	"gopkg.in/yaml.v3"
)

var env *string

func init() {
	env = flag.String("env", "dev", "launch environment [dev/test]")
	flag.Parse()
}

type Config struct {
	LogLevel string        `yaml:"loglevel"`
	App      ServiceConfig `yaml:"app"`
	Auth     ServiceConfig `yaml:"auth"`
	Profile  ServiceConfig `yaml:"profile"`
}

type ServiceConfig struct {
	Addr string            `yaml:"addr"`
	ENV  map[string]EnvCfg `yaml:"env"`
}

type DBConfig struct {
	URL string `yaml:"url"`
}

type Addr struct {
	Addr string `yaml:"addr"`
}

type EnvCfg struct {
	ExtAddr string `yaml:"extaddr"`
	DBURL   string `yaml:"dburl"`
}

// GetConfiguration ...
func LoadConfiguration(filepath string) (*Config, EnvConfigurationProvider) {
	cfg := loadConfig(filepath)
	configureLogger(cfg.LogLevel)
	cfgProvider := NewConfigProvider(cfg, *env)

	return cfg, cfgProvider
}

func loadConfig(filepath string) *Config {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var cfg Config
	if err := yaml.NewDecoder(file).Decode(&cfg); err != nil {
		log.Fatal(err)
	}

	return &cfg
}
func configureLogger(logLevel string) {
	switch logLevel {
	case "debug":
		slog.SetLogLoggerLevel(slog.LevelDebug)
	case "info":
		slog.SetLogLoggerLevel(slog.LevelInfo)
	default:
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}
}
