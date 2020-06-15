// Package cmd is
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	action "psa/aws/rds"
)

var typeRDS string
var snapName string
var stackSnap string

// createSnapshotCmd represents the createSnapshot command
var createSnapshotCmd = &cobra.Command{
	Use:   "createSnapshot",
	Short: "create Snapshot rds",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("createSnapshot called")
		if stackSnap == "" || typeRDS == "" || snapName == "" {
			fmt.Println("you must define stackName (integration-vi) \nand rds-type (rdstecdoc rdsranges) \nand snapshot name")
		} else {
			action.ActionCreateSnapshot(stack, typeRDS, snapName)
		}

	},
}

func init() {
	rdsCmd.AddCommand(createSnapshotCmd)
	createSnapshotCmd.Flags().StringVarP(&stackSnap, "stack", "s", "", "stack name i.e. integration-v1")
	createSnapshotCmd.Flags().StringVarP(&typeRDS, "typerds", "t", "", "mongo type cms,data,front")
	createSnapshotCmd.Flags().StringVarP(&snapName, "snapName", "n", "", "mongo type cms,data,front")
}
