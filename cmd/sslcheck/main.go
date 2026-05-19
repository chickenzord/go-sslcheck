package main

import (
	"os"

	"github.com/chickenzord/go-sslcheck"
	"gopkg.in/yaml.v2"
)

func main() {
	address := os.Args[1]

	result, err := sslcheck.CheckTCP(address)
	if err != nil {
		panic(err)
	}

	yaml.NewEncoder(os.Stdout).Encode(result)
}
