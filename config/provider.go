package config

// EnvConfigurationProvider is an abstraction for environment configuration
type EnvConfigurationProvider interface {
	AuthEnvCfg() EnvCfg
	ProfileEnvCfg() EnvCfg
}

// ConfigProvider is responsible for providing configurations
type ConfigProvider struct {
	cfg *Config
	env string
}

func NewConfigProvider(cfg *Config, env string) EnvConfigurationProvider {
	return &ConfigProvider{cfg: cfg, env: env}
}

func (p *ConfigProvider) AuthEnvCfg() EnvCfg {
	return p.cfg.Auth.ENV[p.env]
}

func (p *ConfigProvider) ProfileEnvCfg() EnvCfg {
	return p.cfg.Profile.ENV[p.env]
}
