package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"lopher/log"
	"net/http"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

var parse string
var short bool
var parsingCmd = &cobra.Command{
	Use:   "parsing",
	Short: "URL parsing and inspection toolkit",
	Run: func(cmd *cobra.Command, args []string) {
		if parse == "" {
			log.ErrorLogger.Println("Please provide a URL to parse")
			cmd.Help()
			return
		}

		if short {
			shortUrl, err := shortenUrl(parse)
			if err != nil {
				log.ErrorLogger.Println(err)
				return
			}
			fmt.Println("🔗 Shortened URL:")
			fmt.Println("---------------------------")
			fmt.Printf("%s\n", shortUrl)
			fmt.Println()
			fmt.Println("🔎 Original URL:")
			fmt.Println("---------------------------")
			fmt.Printf("%s\n", parse)
		} else {
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
		}
	},
}

func init() {
	parsingCmd.PersistentFlags().StringVarP(&parse, "parse", "p", "", "URL to parse")
	parsingCmd.PersistentFlags().BoolVarP(&short, "short", "s", false, "Short output")
	parsingCmd.MarkPersistentFlagRequired("parse")
	rootCmd.AddCommand(parsingCmd)
}

func shortenUrl(url string) (string, error) {
	jsonBody := fmt.Sprintf(`{"url": "%s"}`, url)
	tinyUrl := os.Getenv("TINY_API_URL")

	resp, err := http.Post(tinyUrl, "application/json", bytes.NewBuffer([]byte(jsonBody)))
	if err != nil {
		log.ErrorLogger.Println(err)
		return "", err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.ErrorLogger.Println(err)
		return "", err
	}
	if val, ok := result["data"].(map[string]interface{})["tiny_url"]; ok {
		return val.(string), nil
	}
	return "", fmt.Errorf("no url found in response")
}
