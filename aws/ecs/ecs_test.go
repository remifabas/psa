package ecs

import (
	"fmt"
	"testing"
)

func TestECSClientSession(t *testing.T) {
	ecsService := getECSClientSession()
	expected := "*ecs.ECS"
	ecsType := fmt.Sprintf("%T", ecsService)
	if expected != ecsType {
		t.Errorf("Expected %s got rdsType %s", expected, ecsType)
	}
}

func TestListClusters(t *testing.T) {
	clusters, err := listClusters(getECSClientSession())
	if err != nil {
		t.Errorf("This %s should not appear", err)
	}
	//for _, c := range clusters {
	//	fmt.Println(*c)
	//}
	fmt.Println(&clusters)
}

func TestGetCluster(t *testing.T) {
	ecsSession := getECSClientSession()
	clusters, err := listClusters(ecsSession)
	if err != nil {
		t.Errorf("This %s should not appear", err)
	}
	c, err2 := getCluster(clusters, "integration-v2", "ECSData")
	if err2 != nil {
		t.Errorf("This %v should not appear", err2)
	}
	fmt.Println(*c)
}

func TestGetDataCluster(t *testing.T) {
	clusters, err := listClusters(getECSClientSession())
	if err != nil {
		t.Errorf("This %v should not appear", err)
	}

	c, nerr := getDataCluster(clusters, "integration-v2")

	if nerr != nil {
		t.Errorf("This %v should not appear", nerr)
	}

	fmt.Println(*c)
}

func TestGetFrontCluster(t *testing.T) {
	clusters, err := listClusters(getECSClientSession())
	if err != nil {
		t.Errorf("This %v should not appear", err)
	}

	c, nerr := getFrontCluster(clusters, "integration-v2")

	if nerr != nil {
		t.Errorf("This %v should not appear", nerr)
	}

	fmt.Println(*c)
}

func TestGetServices(t *testing.T) {
	ecsS := getECSClientSession()

	clusters, err := listClusters(ecsS)
	if err != nil {
		t.Errorf("This %v should not appear", err)
	}

	c, nerr := getFrontCluster(clusters, "integration-v2")
	if nerr != nil {
		t.Errorf("This %v should not appear", nerr)
	}

	servicesArn, serviceerr := getServices(ecsS, c)
	if serviceerr != nil {
		t.Errorf("This %v shouln't appear", serviceerr)
	}
	fmt.Println(servicesArn)
}

func TestGetService(t *testing.T) {
	ecsS := getECSClientSession()

	clusters, err := listClusters(ecsS)
	if err != nil {
		t.Errorf("List of cluster ... %v should not appear", err)
	}

	c, nerr := getFrontCluster(clusters, "integration-v2")
	if nerr != nil {
		t.Errorf("This %v should not appear", nerr)
	}

	servicesArn, serviceerr := getServices(ecsS, c)
	if serviceerr != nil {
		t.Errorf("This %v shouln't appear", serviceerr)
	}
	theServ, getServ := getService(servicesArn, "Content")
	if getServ != nil {
		t.Errorf("The service must be returned, %v shouln't appear", getServ)
	}
	fmt.Println(*theServ)
}

func TestForceNewDeploy(t *testing.T) {
	ecsS := getECSClientSession()

	clusters, err := listClusters(ecsS)
	if err != nil {
		t.Errorf("List of cluster ... %v should not appear", err)
	}

	c, nerr := getFrontCluster(clusters, "integration-v2")
	if nerr != nil {
		t.Errorf("This %v should not appear", nerr)
	}

	servicesArn, serviceerr := getServices(ecsS, c)
	if serviceerr != nil {
		t.Errorf("This %v shouln't appear", serviceerr)
	}
	theServ, getServ := getService(servicesArn, "Content")
	if getServ != nil {
		t.Errorf("The service must be returned, %v shouln't appear", getServ)
	}
	fmt.Println(*theServ)

	serviceARN, errDeploy := forceNewDeploy(ecsS, theServ, c)
	if errDeploy != nil {
		t.Errorf("Service Arn must be returned, Error : %v", errDeploy)
	}

	fmt.Println(*serviceARN)
}
