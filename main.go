/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/joho/godotenv"
	"log"
	"nootebook.com/cmd"
	"nootebook.com/config"
)

func main() {
	config.LoadConfigFile("./config/")
	cmd.Execute()
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
}
