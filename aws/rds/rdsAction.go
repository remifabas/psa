package rds

import (
	"fmt"
	"os"
)

// ActionCreateSnapshot create some snap from
func ActionCreateSnapshot(stackName string, typeRDS string, snapshotName string) {
	rdsService := getRdsCLientSession()
	list := listDBInstances(rdsService)

	dbInstanceID := getDBInstanceID(list, stackName, typeRDS)

	status, err := createSnapshot(rdsService, dbInstanceID, snapshotName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Printf("Status %s, snapshot name : %s", status, snapshotName)
	}
}
