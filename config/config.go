package config

import "flag"

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

// LoafConfig ...
func LoadConfig(filepath string, reader Reader) *Config {
	cfg := reader.ReadFile(filepath)
	return cfg
}

// GetConfiguration ...
func GetConfiguration(configFilePath string) (*Config, EnvConfigurationProvider) {
	cfgReader := NewFileConfigReader()
	cfg := LoadConfig("config.yaml", cfgReader)

	slogConfigurator := NewSlogConfigurator()
	slogConfigurator.ConfigureLogger(cfg.LogLevel)

	cfgProvider := NewConfigProvider(cfg, *env)
	return cfg, cfgProvider
}
