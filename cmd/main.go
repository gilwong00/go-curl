package main

import (
	"fmt"
	"os"

	"github.com/gilwong00/go-curl/internal/command"
)

func main() {
	if err := command.CreateRootCommand().Execute(); err != nil {
		fmt.Printf("error: %v\n", err.Error())
		os.Exit(1)
	}
}
