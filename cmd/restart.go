// Package cmd ...
package cmd

import (
	"fmt"

	action "psa/aws/ecs"

	"github.com/spf13/cobra"
)

var cluster string
var service string

// restartCmd represents the restart command
var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(stack) <= 0 || len(cluster) <= 0 || len(service) <= 0 {
			if len(stack) <= 0 {
				fmt.Println("you must define --stack (integration-vi)")
			}
			if len(cluster) <= 0 {
				fmt.Println("you must define --cluster (Data,Front,Tools)")
			}
			if len(service) <= 0 {
				fmt.Println("you must define --service micro-service with name capitalized")
			}
		} else {
			action.ActionRestart(stack, cluster, service)
		}

	},
}

func init() {
	ecsCmd.AddCommand(restartCmd)
	restartCmd.Flags().StringVarP(&stack, "stack", "", "", "stack name i.e. integration-v1")
	restartCmd.Flags().StringVarP(&cluster, "type", "", "", "cluster i.e. Front/Data/Tools")
	restartCmd.Flags().StringVarP(&service, "service", "", "", "cluster i.e. Front/Data/Tools")
}
