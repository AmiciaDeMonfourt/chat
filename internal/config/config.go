package config

import (
	"log"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	Addr         string `yaml:"addr"`
	AuthAddr     string `yaml:"auth_addr"`
	ProfileDBURL string `yaml:"profiledb_url"`
	LogLevel     string `yaml:"log_level"`
}

type ConfigFiles struct {
	Profile map[string]*AppConfig `yaml:"app"`
}

func New(env string) *AppConfig {
	wr := "failed to read config file"
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("%s, %v", wr, err)
	}

	// load .env file
	if err := godotenv.Load(wd + "/.env"); err != nil {
		log.Fatalf("%s, %v", wr, err)
	}

	// load config.yaml file
	file, err := os.Open(wd + "/config.yaml")
	if err != nil {
		log.Fatalf("%s, %v", wr, err)
	}
	defer file.Close()

	// parse config.yaml file
	config := &ConfigFiles{}
	if err := yaml.NewDecoder(file).Decode(config); err != nil {
		log.Fatalf("%s, %v", wr, err)
	}

	// check if the required config is available
	cfg, ok := config.Profile[env]
	if !ok {
		log.Fatalf("%s, no such env %v", wr, err)
	}

	expandenv(cfg)

	// if cfg.LogLevel == "debug" {
	// }
	slog.SetLogLoggerLevel(slog.LevelDebug)

	return cfg
}

func expandenv(cfg *AppConfig) {
	wr := "failed to read config file"
	cfg.Addr = os.ExpandEnv(cfg.Addr)
	if cfg.Addr == "" {
		log.Fatalf("%s, missing cfg.Addr", wr)
	}

	cfg.ProfileDBURL = os.ExpandEnv(cfg.ProfileDBURL)
	if cfg.ProfileDBURL == "" {
		log.Fatalf("%s, missing ProfileDBURL", wr)
	}

	cfg.AuthAddr = os.ExpandEnv(cfg.AuthAddr)
	if cfg.ProfileDBURL == "" {
		log.Fatalf("%s, missing AuthAddr", wr)
	}
}
