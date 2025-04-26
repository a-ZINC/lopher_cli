package cmd

import (
	"lopher/algo"
	"lopher/log"
	"lopher/utils"
	"os"

	"github.com/spf13/cobra"
)

var alg string
var data string
var file string

var allowedAlgorithms = []string{"sha256", "sha512", "md5"}

var hashCmd = &cobra.Command{
	Use:   "hash",
	Short: "Hash files or strings (algorithms: sha256, sha512, md5)",
	Long: `Hash input using cryptographic algorithms.

	Available algorithms:
	- sha256
	- sha512
	- md5

	Examples:
	gopher hash --input "hello" --alg sha256
	gopher hash -i "world" -a sha512`,
	Run: func(cmd *cobra.Command, args []string) {
		if data == "" && file == "" {
			log.ErrorLogger.Println("Please provide input data or a file")
			cmd.Help()
			return
		}
		if data != "" && file != "" {
			log.ErrorLogger.Println("Please provide either input data or a file")
			cmd.Help()
			return
		}
		if !utils.Contains(allowedAlgorithms, alg) {
			log.ErrorLogger.Printf("Invalid algorithm: %s\n", alg)
			closestAlgo := algo.ClosestDistance(allowedAlgorithms, alg, 3)
			log.WarnLogger.Printf("Did you mean %s?\n", closestAlgo)
			return
		}
		var inputByte []byte
		var err error
		if data != "" {
			inputByte = []byte(data)
		} else {
			inputByte, err = os.ReadFile(file)
			if err != nil {
				log.ErrorLogger.Println(err)
				return
			}
		}

		hash := utils.Hash(alg, inputByte)
		file, err := os.Create("hash.txt")
		if err != nil {
			log.ErrorLogger.Println(err)
			return
		}
		defer file.Close()
		_, err = file.WriteString(hash)
		if err != nil {
			log.ErrorLogger.Println(err)
			return
		}
		log.InfoLogger.Printf("\n Algorithm: %s \n Hash: %s\n Data: %s\n", alg, hash, string(inputByte))
	},
}

func init() {
	hashCmd.PersistentFlags().StringVarP(&alg, "alg", "a", "sha256", "Hash algorithm")
	hashCmd.PersistentFlags().StringVarP(&data, "input", "i", "", "Data to hash")
	hashCmd.RegisterFlagCompletionFunc("alg", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"sha256", "sha512", "md5"}, cobra.ShellCompDirectiveNoFileComp
	})
	hashCmd.PersistentFlags().StringVarP(&file, "file", "f", "", "File to hash")
	hashCmd.MarkPersistentFlagRequired("data")
	rootCmd.AddCommand(hashCmd)
}
