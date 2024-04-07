package main

import "github.com/gutkedu/go-expert-api/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBHost)
}
