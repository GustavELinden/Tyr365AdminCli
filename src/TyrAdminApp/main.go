/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"

	"github.com/GustavELinden/TyrAdminCli/365Admin/authentication"
	"github.com/GustavELinden/TyrAdminCli/365Admin/cmd"
	"github.com/spf13/viper"
)




func main() {
	viper := viper.New()
	viper.SetConfigName("config")

	viper.SetConfigType("json")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(`C:\Users\fintv\Desktop\Testmapp`) // path to look for the config file in                                             // optionally look for config in the working directory
	err := viper.ReadInConfig()                                              // Find and read the config file
	if err != nil {                                                          // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}


    // Now you can access your configuration settings using Viper
		fmt.Println("Using config file:", viper.GetString("M365managementAppClientId"))

    accessToken := authentication.GetTokenForGovernanceApi()
    fmt.Println(accessToken)
	  cmd.Execute()
}
