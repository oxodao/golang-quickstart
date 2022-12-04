package main

import (
	"fmt"
	"os"

	"github.com/oxodao/go-quickstart/cmd"
	"github.com/oxodao/go-quickstart/config"
	"github.com/oxodao/go-quickstart/services"
)

func main() {
	if err := config.Load(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := services.Load(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cmd.Execute()
}
