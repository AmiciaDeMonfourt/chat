package config

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type AuthConfig struct {
	Addr        string `yaml:"addr"`
	DBURL       string `yaml:"db_url"`
	LogLevel    string `yaml:"log_level"`
	ProfileAddr string `yaml:"profile_addr"`
}

type ConfigFiles struct {
	Auth map[string]*AuthConfig `yaml:"auth"`
}

var (
	env = *flag.String("env", "dev", "Environment to use [dev/test]")
)

func New() *AuthConfig {
	flag.Parse()

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
	cfg, ok := config.Auth[env]
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

	cfg.ProfileAddr = os.ExpandEnv(cfg.ProfileAddr)
	if cfg.DBURL == "" {
		log.Fatalf("%s, missing ProfileAddr", wr)
	}

	return cfg
}
