// Package cmd is
package cmd

import (
	"fmt"

	action "psa/aws/cloudformation"

	"github.com/spf13/cobra"
)

var stackName string
var typeMongo string

// deleteMongoStackCmd represents the deleteStack command
var deleteMongoStackCmd = &cobra.Command{
	Use:   "deleteMongoStack",
	Short: "delete Mongo Stack not used you need to give a stack name (i.e integration-v1) and type of stack to delete (cms,data,front)",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println(stackName)

		if stackName == "" || typeMongo == "" {
			fmt.Println("you must define --stack (integration-vi) \nand mongo --type (cms,data,front)")
		} else {
			action.ActionDeleteStack(stackName, typeMongo)
		}

	},
}

func init() {
	cloudformationCmd.AddCommand(deleteMongoStackCmd)
	deleteMongoStackCmd.Flags().StringVarP(&stackName, "stack", "s", "", "stack name i.e. integration-v1")
	deleteMongoStackCmd.Flags().StringVarP(&typeMongo, "type", "t", "", "mongo type cms,data,front")
}
