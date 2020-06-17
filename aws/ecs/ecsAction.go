package ecs

import (
	"log"
)

// ActionRestart for package cmd
func ActionRestart(stack string, cluster string, service string) {
	ecsS := getECSClientSession()

	clusters, err := listClusters(ecsS)
	if err != nil {
		log.Fatalf("List of cluster ... %v should not appear", err)
	}

	c, nerr := getCluster(clusters, stack, cluster)
	if nerr != nil {
		log.Fatalf("This %v should not appear", nerr)
	}

	servicesArn, serviceerr := getServices(ecsS, c)
	if serviceerr != nil {
		log.Fatalf("This %v shouln't appear", serviceerr)
	}
	theServ, getServ := getService(servicesArn, service)
	if getServ != nil {
		log.Fatalf("The service must be returned, %v shouln't appear", getServ)
	}

	forceNewDeploy(ecsS, theServ, c)
}
