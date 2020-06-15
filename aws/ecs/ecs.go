package ecs

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecs"
)

func getECSClientSession() *ecs.ECS {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file. (~/.aws/credentials).
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-1")},
	)
	if err != nil {
		fmt.Println("Error:", err)
	}
	ecsService := ecs.New(sess)

	return ecsService
}

func listClusters(ecsService *ecs.ECS) ([]*string, error) {

	input := &ecs.ListClustersInput{}
	output, err := ecsService.ListClusters(input)

	if err != nil {
		return nil, err
	}

	clusterARN := output.ClusterArns

	return clusterARN, nil
}
