package main

import (
	"fmt"

	"github.com/GustavELinden/Tyr365AdminCli/cmd"
	viperConfig "github.com/GustavELinden/Tyr365AdminCli/config"
)

func main() {
	_, err := viperConfig.InitViper("config.json")
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	cmd.Execute()
}
