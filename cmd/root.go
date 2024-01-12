package cmd

import (
	"os"

	"github.com/arfurlaneto/goexpert-challenge-stress-test/internal"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goexpert-stresser",
	Short: "A Golang stress tester",
	Long:  "A Golang stress tester made as one of the final challenges of https://goexpert.fullcycle.com.br/pos-goexpert/",
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.PersistentFlags().GetString("url")
		requests, _ := cmd.PersistentFlags().GetInt64("requests")
		concurrency, _ := cmd.PersistentFlags().GetInt64("concurrency")
		internal.RunStressTester(url, requests, concurrency)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("url", "u", "", "URL to be tested")
	rootCmd.PersistentFlags().Int64P("requests", "r", 0, "total requests to be made")
	rootCmd.PersistentFlags().Int64P("concurrency", "c", 0, "max number of concurrent requests")
	rootCmd.MarkPersistentFlagRequired("url")
	rootCmd.MarkPersistentFlagRequired("requests")
	rootCmd.MarkPersistentFlagRequired("concurrency")
}
