package main

import (
	"log"

	"github.com/spf13/viper"
)

func main() {
	if err := init_config(); err != nil {
		log.Fatalf("Error initializing config: %v", err)
	}
	if err := init_db(); err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	init_router()

	r.Run(":" + viper.GetString("server.port"))
}
