package cmd

import (
	"fmt"
	"os"

	"vivid-bridge/cmd/vivid-guardiand/cmd/guardian"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vivid-guardiand",
	Short: "vivid-bridge guardian node",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("exiting rootCmd")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(guardian.GuardianCmd)

}
