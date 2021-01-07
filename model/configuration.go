package model

import (
	"encoding/json"
	"os"
	"time"
)

type Configuration struct {
	DBUser          string
	DBPass          string
	DBName          string
	DBHost          string
	DBPort          int
	DBMaxOpenConns  int
	DBMaxIdleConns  int
	ConnMaxLifeTime time.Duration
	DBLogMode       bool
	DBMigrate       bool
	HTTPPort        int
	RedisHost       string
	RedisDB         int
	RedisPassword   string
}

//ReadFile
func (c *Configuration) ReadFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&c)
	if err != nil {
		panic(err)
	}
}

//NewConfiguration
func NewConfiguration(path string) *Configuration {
	var configuration Configuration
	configuration.ReadFile(path)
	return &configuration
}
