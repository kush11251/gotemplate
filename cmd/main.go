package main

import (
	"fmt"
	"gotemplate/pkg/config"
	"gotemplate/pkg/logger"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Fatal(err)
	}
	fmt.Println(cfg)
}