package rds

import (
	"fmt"
	"strings"
	"testing"
)

func TestInfoRds(t *testing.T) {
	rdsService := getRdsCLientSession()
	expected := "*rds.RDS"
	rdsType := fmt.Sprintf("%T", rdsService)
	if expected != rdsType {
		t.Errorf("Expected %s got rdsType %s", expected, rdsType)
	}
}

func TestListDBInstances(t *testing.T) {
	rdsService := getRdsCLientSession()
	list := listDBInstances(rdsService)
	if len(list) < 0 {
		t.Errorf("Expected more than %d DB Instances", len(list))
	}
}

func TestGetRangesDBFromEnv(t *testing.T) {
	rdsService := getRdsCLientSession()
	list := listDBInstances(rdsService)

	dbInstanceID := getDBInstanceID(list, "integration-v2", "rdsranges")
	fmt.Println(dbInstanceID)
}

func TestCreateSnapshot(t *testing.T) {
	rdsService := getRdsCLientSession()
	list := listDBInstances(rdsService)

	dbInstanceID := getDBInstanceID(list, "integration-v2", "rdsranges")

	_, err := createSnapshot(rdsService, dbInstanceID, "rdsrangestestingGO")

	if !strings.Contains(err.Error(), "DBSnapshotAlreadyExists: Cannot create the snapshot because a snapshot with the identifier rdsrangestestinggo already exists.") {
		t.Errorf("This error must be launched : %s", err.Error())
	}
}
