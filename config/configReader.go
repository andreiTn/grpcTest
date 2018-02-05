package config

import (
	"os"
	"fmt"
	"encoding/json"
	"log"
)

type Config struct {
	TemplateLayoutPath string
	TemplateIncludePath string
	Host string
	Port string
}

func (c *Config) LoadConfig(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("There was an error reading the file,", err)
		return err
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&c)
	if err != nil {
		log.Println("error:", err)
		return err
	}
	return nil
}

var AppConfig Config

func init() {
	//err := loadConfiguration("config/default.json")
	err := AppConfig.LoadConfig("config/default.json")

	if err != nil {
		fmt.Println("Could not load config: ", err)
	}
}