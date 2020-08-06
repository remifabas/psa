// Package cmd ...
package cmd

import (
	"github.com/spf13/cobra"
)

// ecsCmd represents the ecs command
var ecsCmd = &cobra.Command{
	Use:   "ecs",
	Short: "Interact with psa aws ecs service",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(ecsCmd)
}
