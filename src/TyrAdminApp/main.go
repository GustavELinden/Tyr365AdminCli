/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"

	"github.com/GustavELinden/TyrAdminCli/365Admin/cmd"
	viperConfig "github.com/GustavELinden/TyrAdminCli/365Admin/config"
	getgov "github.com/GustavELinden/TyrAdminCli/365Admin/httpFuncs"
)




func main() {
	viper, err := viperConfig.InitViper("config.json")

	if err != nil {                                                          // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

fmt.Println(viper.ConfigFileUsed())
    // Now you can access your configuration settings using Viper
		fmt.Println("Using config file:", viper.GetString("M365managementAppClientId"))

    getgov.Get("GetProvisioningStatus", map[string]string{"requestId": "147999"})
	  cmd.Execute()
}
