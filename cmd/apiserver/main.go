package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/RomanDovgii/go-restapi/internal/app/apiserver"
)

var (
	configPath string
)

func onot() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to congif file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()

	_, err := toml.DecodeFile(configPath, config)

	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.New(config)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
