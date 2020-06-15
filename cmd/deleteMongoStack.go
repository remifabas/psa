// Package cmd is
package cmd

import (
	"fmt"

	action "psa/aws/cloudformation"

	"github.com/spf13/cobra"
)

var stack string
var typeMongo string

// deleteMongoStackCmd represents the deleteStack command
var deleteMongoStackCmd = &cobra.Command{
	DisableFlagParsing: true,
	Use:                "deleteMongoStack",
	Short:              "delete Mongo Stack not used you need to give a stack name (i.e integration-v1) and type of stack to delete (cms,data,front)",
	Long:               ``,
	Run: func(cmd *cobra.Command, args []string) {

		if stack == "" || typeMongo == "" {
			fmt.Println("you must define stackName (integration-vi) \nand mongo type (cms,data,front)")
		} else {
			action.ActionDeleteStack(stack, typeMongo)
		}

	},
}

func init() {
	cloudformationCmd.AddCommand(deleteMongoStackCmd)
	deleteMongoStackCmd.Flags().StringVarP(&stack, "stackName", "s", "", "stack name i.e. integration-v1")
	deleteMongoStackCmd.Flags().StringVarP(&typeMongo, "typeMongo", "t", "", "mongo type cms,data,front")
}
