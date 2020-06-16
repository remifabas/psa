package ecs

import (
	"errors"
	"fmt"
	"strings"

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

func getCluster(listClusters []*string, stack string, ecs string) (*string, error) {
	for _, c := range listClusters {
		if strings.Contains(*c, stack) && strings.Contains(*c, ecs) {
			return c, nil
		}
	}
	// If not found return error not found
	return nil, errors.New("ERROR :stack not found")
}

func getDataCluster(listClusters []*string, stack string) (*string, error) {
	return getCluster(listClusters, stack, "ECSData")
}

func getFrontCluster(listClusters []*string, stack string) (*string, error) {
	return getCluster(listClusters, stack, "ECSFront")
}

func getServices(ecsService *ecs.ECS, cluster *string) ([]*string, error) {

	input := &ecs.ListServicesInput{Cluster: cluster}
	output, err := ecsService.ListServices(input)
	if err != nil {
		return nil, err
	}
	return output.ServiceArns, nil
}

func getService(services []*string, service string) (*string, error) {
	for _, s := range services {
		if strings.Contains(*s, service) {
			return s, nil
		}
	}
	return nil, errors.New("ERROR : service not found")
}

func forceNewDeploy(ecsService *ecs.ECS, service *string, clusterArn *string) error {
	input := &ecs.UpdateServiceInput{
		ForceNewDeployment: aws.Bool(true),
		Service:            service,
		Cluster:            clusterArn,
	}
	_, err := ecsService.UpdateService(input)
	if err != nil {
		return err
	}
	return nil
}
