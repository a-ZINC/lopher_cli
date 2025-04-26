package cmd

import (
	"lopher/algo"
	"lopher/log"
	"lopher/utils"

	"github.com/spf13/cobra"
)

var alg string
var data string

var allowedAlgorithms = []string{"sha256", "sha512", "md5"}

var hashCmd = &cobra.Command{
	Use:   "hash",
	Short: "Hash files or strings (algorithms: sha256, sha512, md5)",
	Long:  `Hash input using cryptographic algorithms.

	Available algorithms:
	- sha256
	- sha512
	- md5

	Examples:
	gopher hash --input "hello" --alg sha256
	gopher hash -i "world" -a sha512`,
	Run: func(cmd *cobra.Command, args []string) {
		if data == "" {
			cmd.Help()
			return
		}
		if !utils.Contains(allowedAlgorithms, alg) {
			log.ErrorLogger.Printf("Invalid algorithm: %s\n", alg)
			closestAlgo := algo.ClosestDistance(allowedAlgorithms, alg, 3)
			log.WarnLogger.Printf("Did you mean %s?\n", closestAlgo)
			return
		}
		hash := utils.Hash(alg, data)
		log.InfoLogger.Printf("\n Algorithm: %s \n Hash: %s\n Data: %s\n", alg, hash, data)
	},
}

func init() {
	hashCmd.PersistentFlags().StringVarP(&alg, "alg", "a", "sha256", "Hash algorithm")
	hashCmd.PersistentFlags().StringVarP(&data, "input", "i", "", "Data to hash")
	hashCmd.RegisterFlagCompletionFunc("alg", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"sha256", "sha512", "md5"}, cobra.ShellCompDirectiveNoFileComp
	})
	hashCmd.MarkPersistentFlagRequired("data")
	rootCmd.AddCommand(hashCmd)
}
