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
	listClusters(getECSClientSession())
}
