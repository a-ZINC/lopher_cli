package main

import (
	"lopher/cmd"
	"lopher/log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.ErrorLogger.Println(err)
		return
	}
	cmd.Execute()
}
