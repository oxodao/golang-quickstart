package cmd

import (
	"fmt"
	"strings"

	"github.com/oxodao/go-quickstart/utils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   strings.ToLower(utils.APP_NAME),
	Short: "The main app",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Oxodao's Golang quickstart")
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
