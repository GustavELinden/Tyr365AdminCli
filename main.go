package main

import (
	"fmt"

	"github.com/GustavELinden/Tyr365AdminCli/cmd"
	"github.com/GustavELinden/Tyr365AdminCli/internal/config"
)

func main() {
	if err := config.Initialize("config.json"); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	cmd.Execute()
}
