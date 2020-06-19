package ec2

import (
	"fmt"
	"testing"
)

func TestGetSession(t *testing.T) {
	_, errGetEC2 := getEC2Session()
	if errGetEC2 != nil {
		t.Errorf("Error : get ec2 session ... %v", errGetEC2)
	}
}

func TestGetInstances(t *testing.T) {
	ec2Sess, errGetEC2 := getEC2Session()
	if errGetEC2 != nil {
		t.Errorf("Error : get ec2 session ... %v", errGetEC2)
	}
	_, errGetInstances := getInstances(ec2Sess)
	if errGetInstances != nil {
		t.Errorf("Error : get connection ... %v", errGetEC2)
	}
}

func TestGetMongoDBInstances(t *testing.T) {
	ec2Sess, errGetEC2 := getEC2Session()
	if errGetEC2 != nil {
		t.Errorf("Error : get ec2 session ... %v", errGetEC2)
	}
	instances, errGetInstances := getInstances(ec2Sess)
	if errGetInstances != nil {
		t.Errorf("Error : get connection ... %v", errGetEC2)
	}
	fmt.Println(len(instances))
	mongoInstances, errMongoDB := getMongoDBInstances(instances)
	if errMongoDB != nil {
		t.Errorf("Error : get Mongo instances ... %v", errGetEC2)
	}
	fmt.Println(len(mongoInstances))
}
