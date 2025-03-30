package main

import (
	"log"

	"github.com/FT1006/blogator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	
	err = cfg.SetUser("FT1006")

	cfg, err = config.Read()
	if err != nil {
		log.Fatal(err)
	}
	
	println(cfg.DBUrl)
	println(cfg.CurrentUserName)
}
