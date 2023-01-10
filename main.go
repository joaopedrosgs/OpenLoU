package main

import (
	"flag"
	"fmt"

	"github.com/spf13/viper"

	_ "github.com/lib/pq"
	"openlou/app"
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
