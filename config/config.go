package config

import (
	"log"
	"sync"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port                      string `envconfig:"PORT"`
	DefaultExpirationDuration int    `envconfig:"DEFAULT_EXPIRATION_DURATION"`
	PurgeExpiredItemsDuration int    `envconfig:"PURGE_EXPIRED_ITEMS_DURATION"`
}

var (
	conf Config
	once sync.Once
)

// Get are responsible to load env and get data an return the struct
func Get() *Config {
	once.Do(func() {
		log.Println("Service configuration initialized.")
		err := envconfig.Process("", &conf)
		if err != nil {
			log.Printf("err : %v \n", err)
		}
	})

	return &conf
}
