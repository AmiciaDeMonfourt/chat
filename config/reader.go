package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Reader interface {
	ReadFile(filepath string) *Config
}

type FileConfigReader struct{}

func NewFileConfigReader() Reader {
	return &FileConfigReader{}
}

func (r *FileConfigReader) ReadFile(filepath string) *Config {
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
