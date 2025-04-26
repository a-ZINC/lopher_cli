package cmd

import (
	"lopher/log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Lopher-cli",
	Short: "Lopher CLI - Bro Army Knife for Developers",
	Long: `A versatile CLI tool that handles files, URLs, hashing, time formatting, and more.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.ErrorLogger.Println(err)
		os.Exit(1)
	}
}

