package main

import (
	"fmt"

	"github.com/GustavELinden/TyrAdminCli/cmd"
	viperConfig "github.com/GustavELinden/TyrAdminCli/config"
)

func main() {
	_, err := viperConfig.InitViper("config.json")
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	cmd.Execute()
}
