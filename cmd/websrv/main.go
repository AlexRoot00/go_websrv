package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/websrv.toml", "path to config file")
}
func main() {
	flag.Parse()

	config := websrv.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
	s := websrv.New()
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
