package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/jMurad/statistics-collection/internal/app/statserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/statserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	cfg := statserver.NewConfig()
	_, err := toml.DecodeFile(configPath, cfg)
	if err != nil {
		log.Fatal(err)
	}

	s := statserver.New(cfg)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
