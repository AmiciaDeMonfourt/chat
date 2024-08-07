package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type ProfileConfig struct {
	Addr     string `yaml:"addr"`
	DBURL    string `yaml:"db_url"`
	LogLevel string `yaml:"log_level"`
}

type ConfigFiles struct {
	Profile map[string]*ProfileConfig `yaml:"profile"`
}

func New(env string) *ProfileConfig {
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

	// load env vars
	cfg.Addr = os.ExpandEnv(cfg.Addr)
	if cfg.Addr == "" {
		log.Fatalf("%s, missing cfg.Addr", wr)
	}

	cfg.DBURL = os.ExpandEnv(cfg.DBURL)
	if cfg.DBURL == "" {
		log.Fatalf("%s, missing DBURL", wr)
	}

	return cfg
}
