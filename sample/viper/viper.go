package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	// config := viper.New()
	// config.SetConfigName("esms")
	// config.AddConfigPath("./config")
	// err := config.ReadInConfig()
	//
	// if err != nil {
	// 	fmt.Println("No configuration file loaded - using defaults")
	// }
	// theMessage := config.GetStringMapString("influxdb")
	// fmt.Printf("\n%s\n\n", theMessage)
	//
	// config1 := viper.New()
	// config1.SetConfigName("loinguyen")
	// config1.AddConfigPath("./config1")
	// err = config1.ReadInConfig()
	//
	// if err != nil {
	// 	fmt.Println("No configuration file loaded - using defaults")
	// }
	//
	// message := config1.GetString("loi.name")
	// fmt.Printf("\n%s\n\n", message)
	config := viper.New()
	config.SetConfigName("v1")
	// config.AddConfigPath("/host")
	config.AddRemoteProvider("consul", "10.60.3.231:8500", "/v1/kv/order")
	fmt.Println(config.GetStringMapString("host"))

}
