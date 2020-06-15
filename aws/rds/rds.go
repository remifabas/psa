package rds

import (
	"errors"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rds"
)

func getRdsCLientSession() *rds.RDS {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file. (~/.aws/credentials).
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-1")},
	)
	if err != nil {
		fmt.Println("Error:", err)
	}
	rdsService := rds.New(sess)

	return rdsService
}

func listDBInstances(rdsService *rds.RDS) []*rds.DBInstance {

	input := &rds.DescribeDBInstancesInput{}
	output, _ := rdsService.DescribeDBInstances(input)

	return output.DBInstances
}

// getDBInstanceID return the db identifier from the specific env and type of db
// typeDB can only be rdstecodc or rdsranges
func getDBInstanceID(listDBInstances []*rds.DBInstance, env string, typeDB string) (string, error) {
	var dbInstanceIdentifier string = ""
	for _, dbInstance := range listDBInstances {
		if strings.Contains(*dbInstance.DBSubnetGroup.DBSubnetGroupName, env) {
			if strings.Contains(*dbInstance.DBSubnetGroup.DBSubnetGroupName, typeDB) {
				dbInstanceIdentifier = *dbInstance.DBInstanceIdentifier
			}
		}
	}
	if len(dbInstanceIdentifier) <= 0 {
		return "", errors.New("dbInstanceIdentifier not found")
	}
	return dbInstanceIdentifier, nil
}

func createSnapshot(rdsService *rds.RDS, dbInstanceID string, nameSnap string) (string, error) {
	var status string = "OK"
	input := &rds.CreateDBSnapshotInput{
		DBInstanceIdentifier: aws.String(dbInstanceID),
		DBSnapshotIdentifier: aws.String(nameSnap),
	}
	_, err := rdsService.CreateDBSnapshot(input)
	if err != nil {
		status = "KO"
	}
	return status, err
}
