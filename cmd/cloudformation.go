//Package cmd is ...
package cmd

import (
	"github.com/spf13/cobra"
)

// cloudformationCmd represents the cloudformation command
var cloudformationCmd = &cobra.Command{
	Use:   "cloudformation",
	Short: "interact with psa aws cloudformation service",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(cloudformationCmd)
}
