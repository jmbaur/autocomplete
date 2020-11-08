package cmd

import (
	"github.com/spf13/cobra"
)

const (
	address           = "localhost:50051"
	port              = ":50051"
	defaultDictionary = "/usr/share/dict/words"
)

var (
	rootCmd = &cobra.Command{
		Use:   "autocomp",
		Short: "Short description",
		Long:  "Long Description",
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(cliCmd)
	rootCmd.AddCommand(shellCmd)
	rootCmd.AddCommand(serverCmd)
}
