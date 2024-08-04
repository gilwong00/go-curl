package main

import "github.com/gilwong00/go-curl/internal/command"

func main() {
	if err := command.CreateRootCommand().Execute(); err != nil {
		panic(err)
	}
}
