package cloudformation

import (
	"fmt"
	"log"
	"strings"
	"testing"
)

func TestGetStackByName(t *testing.T) {
	theStack, err := getStackByName("integration-v2")
	if err != nil {
		fmt.Println("Something is wrong with getStackByName")
		t.Errorf("getStackByName theStack is %s", theStack)
	}
	typeof := fmt.Sprintf("%T", theStack)
	expected := "*cloudformation.Stack"

	if expected != typeof && strings.Contains(*theStack.StackName, "integration-v2") {
		t.Errorf("Expected %s got %s", expected, typeof)
	}

	theStack, err = getStackByName("error-v2")
	// this stack must be nil -> dos not exist
	if theStack != nil {
		t.Errorf("Expected nil got %s", theStack)
	}
	if !strings.Contains(err.Error(), "ValidationError: Stack with id error-v2 does not exist") {
		t.Errorf("Expected ValidationError: Stack with id error-v2 does not exist got %s", err.Error())
	}
}

func TestGetMongoDataStackName(t *testing.T) {
	theStack, err := getStackByName("integration-v2")
	if err != nil {
		log.Fatal(err)
	}
	name := getMongoDataStackName(theStack)
	if !(strings.Contains(name, "integration-v2") && strings.Contains(name, "mongo-data")) {
		t.Errorf("Mongo stack name doesn't contains mongo-data or integration-v2 : %s", name)
	}
}

func TestGetMongoCmsStackName(t *testing.T) {
	theStack, err := getStackByName("integration-v2")
	if err != nil {
		log.Fatal(err)
	}
	name := getMongoCmsStackName(theStack)
	if !(strings.Contains(name, "integration-v2") && strings.Contains(name, "mongo-cms")) {
		t.Errorf("Mongo stack name doesn't contains mongo-data or integration-v2 : %s", name)
	}
}
func TestGetMongoFrontStackName(t *testing.T) {
	theStack, err := getStackByName("integration-v2")
	if err != nil {
		log.Fatal(err)
	}
	name := getMongoFrontStackName(theStack)
	if !(strings.Contains(name, "integration-v2") && strings.Contains(name, "mongo-front")) {
		t.Errorf("Mongo stack name doesn't contains mongo-data or integration-v2 : %s", name)
	}
}

func TestFetSnapshotDatabaseRanges(t *testing.T) {
	theStack, err := getStackByName("integration-v2")
	if err != nil {
		log.Fatal(err)
	}
	snap := getSnapshotDatabaseRanges(theStack)
	if !strings.Contains(snap, "rds-ranges") && !strings.Contains(snap, "rds") {
		t.Errorf("Mongo stack name doesn't contains mongo-data or integration-v2 : %s", snap)
	}
}

func TestFilterStacks(t *testing.T) {
	filteredStack := filterStacks("validation-v2", "cms")
	fmt.Println(filteredStack)
}

func TestGetExternalStackInUse(t *testing.T) {
	stackInUse, err := getExternalStackInUse("validation-v2", "data")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(stackInUse)
}

func TestListStackNoUsed(t *testing.T) {
	stacks, err := listStackNotUsed("validation-v2", "cms")
	if err != nil {
		t.Error(err)
	}
	fmt.Println("TestListStackNoUsed : validation-v2")
	fmt.Println(len(stacks))
	for _, s := range stacks {
		fmt.Println(s)
	}
}

func TestDeleteStack(t *testing.T) {
	stackToDelete, err := listStackNotUsed("integration-v2", "data")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(len(stackToDelete))
	err = deleteStack(stackToDelete)
	if err != nil {
		t.Errorf("Expected no error")
	}
}
