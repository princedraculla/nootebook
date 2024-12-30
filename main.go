/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"log"
	"nootebook.com/cmd"
	"nootebook.com/config"
	"nootebook.com/logger"
)

var loger *zap.Logger

func main() {
	loger = logger.InitializeLogger()
	defer loger.Sync()
	config.LoadConfigFile("./config/")
	loger.Info("config file load success")
	loger.Info("command execute success")
	cmd.Execute()
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

}
