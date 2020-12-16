package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func Parse() {
	viper.SetConfigName("e")    // name of config file (without extension)
	viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name

	viper.AddConfigPath("../../") // path to look for the config file in
	err := viper.ReadInConfig()    // Find and read the config file
	if err != nil {                // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	fmt.Print(viper.GetString("db.project"))
}
