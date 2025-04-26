package cmd

import (
	"fmt"
	"lopher/log"
	"net/url"

	"github.com/spf13/cobra"
)

var parse string
var parsingCmd = &cobra.Command{
	Use:   "parsing",
	Short: "URL parsing and inspection toolkit",
	Run: func(cmd *cobra.Command, args []string) {
		if parse == "" {
			log.ErrorLogger.Println("Please provide a URL to parse")
			cmd.Help()
			return
		}
		parsedUrl, err := url.Parse(parse)
		if err != nil {
			log.ErrorLogger.Println(err)
			return
		}
		fmt.Println("🔎 Parsed URL Details:")
		fmt.Println("---------------------------")
		fmt.Printf("Scheme: %s\n", parsedUrl.Scheme)
		fmt.Printf("Host: %s\n", parsedUrl.Host)
		fmt.Printf("Path: %s\n", parsedUrl.Path)

		queryParams := parsedUrl.Query()
		if len(queryParams) > 0 {
			fmt.Println("\nQuery Parameters:")
			for key, values := range queryParams {
				for _, value := range values {
					fmt.Printf("  %s = %s\n", key, value)
				}
			}
		} else {
			fmt.Println("\nNo Query Parameters found.")
		}
	},
}

func init() {
	parsingCmd.PersistentFlags().StringVarP(&parse, "parse", "p", "", "URL to parse")
	parsingCmd.MarkPersistentFlagRequired("parse")
	rootCmd.AddCommand(parsingCmd)
}
