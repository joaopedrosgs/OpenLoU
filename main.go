package main

//go:generate sqlboiler --wipe psql

import (
	"flag"
	"fmt"

	"github.com/spf13/viper"

	"github.com/joaopedrosgs/OpenLoU/app"
	_ "github.com/lib/pq"
)

func main() {
	configPtr := flag.String("config", "config", "Configuration File")
	flag.Parse()
	viper.AddConfigPath(".") // optionally look for config in the working directory
	viper.SetConfigName(*configPtr)
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Failed to read config file: %s \n", err))
	}
	app.Run()
}
