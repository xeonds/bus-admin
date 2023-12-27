package main

import (
	"os"

	"github.com/spf13/viper"
)

//TODO: add config validate method & add default config generate method

func init_config() error {
	if _, err := os.Stat("config.yaml"); err != nil {
		if err := func() error {
			data := []byte(`database.address: localhost:3306`)
			err := os.WriteFile("config.yaml", data, 0644)
			if err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return err
		}
	}
	if err := func() error {
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
		err := viper.ReadInConfig()
		if err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return err
	}

	return nil
}
