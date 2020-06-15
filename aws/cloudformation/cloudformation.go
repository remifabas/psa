package cloudformation

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
)

func getStackByName(stackName string) (*cloudformation.Stack, error) {

	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file. (~/.aws/credentials).
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-1")},
	)
	if err != nil {
		fmt.Println("Error:", err)
	}
	cfsvc := cloudformation.New(sess)

	params := &cloudformation.DescribeStacksInput{
		NextToken: aws.String("NextToken"),
		StackName: aws.String(stackName),
	}
	resp, err := cfsvc.DescribeStacks(params)

	if err != nil {
		return nil, err
	}

	return resp.Stacks[0], nil
	//typeof := fmt.Sprintf("%T", stack.Parameters)
}

func getMongoDataStackName(stack *cloudformation.Stack) string {
	ExternalStackNameMongoData := ""
	for _, s := range stack.Parameters {
		if *s.ParameterKey == "ExternalStackNameMongoData" {
			ExternalStackNameMongoData = *s.ParameterValue
		}
	}
	return ExternalStackNameMongoData
}

func getMongoFrontStackName(stack *cloudformation.Stack) string {
	ExternalStackNameMongoFront := ""
	for _, s := range stack.Parameters {
		if *s.ParameterKey == "ExternalStackNameMongoFront" {
			ExternalStackNameMongoFront = *s.ParameterValue
		}
	}
	return ExternalStackNameMongoFront
}

func getMongoCmsStackName(stack *cloudformation.Stack) string {
	ExternalStackNameMongoCms := ""
	for _, s := range stack.Parameters {
		if *s.ParameterKey == "ExternalStackNameMongoCms" {
			ExternalStackNameMongoCms = *s.ParameterValue
		}
	}
	return ExternalStackNameMongoCms
}

func getSnapshotDatabaseRanges(stack *cloudformation.Stack) string {
	SnapshotDatabaseRanges := ""
	for _, s := range stack.Parameters {
		if *s.ParameterKey == "SnapshotDatabaseRanges" {
			SnapshotDatabaseRanges = *s.ParameterValue
		}

	}
	return SnapshotDatabaseRanges
}

func listStacksOuput(cfsvc *cloudformation.CloudFormation, inputParam *cloudformation.ListStacksInput) (*cloudformation.ListStacksOutput, error) {
	return cfsvc.ListStacks(inputParam)
}

func getAllStacks() []*cloudformation.StackSummary {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-1")},
	)
	if err != nil {
		fmt.Println("Error:", err)
	}

	cfsvc := cloudformation.New(sess)

	/* Possible StackStatusFilter
	enum value set: [REVIEW_IN_PROGRESS, CREATE_FAILED, UPDATE_ROLLBACK_FAILED, UPDATE_ROLLBACK_IN_PROGRESS, CREATE_IN_PROGRESS,
	IMPORT_ROLLBACK_IN_PROGRESS, UPDATE_ROLLBACK_COMPLETE_CLEANUP_IN_PROGRESS, ROLLBACK_IN_PROGRESS, IMPORT_IN_PROGRESS,
	DELETE_COMPLETE, UPDATE_COMPLETE, UPDATE_IN_PROGRESS, DELETE_FAILED, IMPORT_COMPLETE, DELETE_IN_PROGRESS,
	ROLLBACK_COMPLETE, ROLLBACK_FAILED, IMPORT_ROLLBACK_COMPLETE, UPDATE_COMPLETE_CLEANUP_IN_PROGRESS, CREATE_COMPLETE, I
	MPORT_ROLLBACK_FAILED, UPDATE_ROLLBACK_COMPLETE]]
	*/
	//First call without NextToken
	params := &cloudformation.ListStacksInput{
		StackStatusFilter: []*string{aws.String("UPDATE_COMPLETE"), aws.String("ROLLBACK_COMPLETE"),
			aws.String("UPDATE_ROLLBACK_COMPLETE"), aws.String("CREATE_COMPLETE")},
	}
	listStacksO, err := listStacksOuput(cfsvc, params)
	if err != nil {
		fmt.Println("Error:", err)
	}
	listStackSum := listStacksO.StackSummaries
	theNextToken := listStacksO.NextToken
	for theNextToken != nil {
		loopParams := &cloudformation.ListStacksInput{
			NextToken: theNextToken,
			StackStatusFilter: []*string{aws.String("UPDATE_COMPLETE"), aws.String("ROLLBACK_COMPLETE"),
				aws.String("UPDATE_ROLLBACK_COMPLETE"), aws.String("CREATE_COMPLETE")},
		}

		loopListStackOuput, err := listStacksOuput(cfsvc, loopParams)
		if err != nil {
			fmt.Println("Error:", err)
		}
		theNextToken = loopListStackOuput.NextToken
		loopListStackSum := loopListStackOuput.StackSummaries
		listStackSum = append(listStackSum, loopListStackSum...)
	}
	return listStackSum
}

func filterStacks(stackName string, typeStack string) []string {
	allStack := getAllStacks()
	var searchedStacks []string
	for _, s := range allStack {
		if strings.Contains(*s.StackName, stackName) && strings.Contains(*s.StackName, typeStack) {
			searchedStacks = append(searchedStacks, *s.StackName)
		}
	}
	return searchedStacks
}

func getExternalStackInUse(stackName string, typeExternalStack string) (string, error) {
	currentStack, err := getStackByName(stackName)
	if err != nil {
		fmt.Println("Something is wrong with getStackByName")
		return "", err
	}
	externalStack := ""
	if "cms" == typeExternalStack {
		externalStack = getMongoCmsStackName(currentStack)
	}
	switch os := typeExternalStack; os {
	case "cms":
		externalStack = getMongoCmsStackName(currentStack)
	case "data":
		externalStack = getMongoDataStackName(currentStack)
	case "front":
		externalStack = getMongoFrontStackName(currentStack)
	default:
		externalStack = ""
	}
	return externalStack, nil
}

func listStackNotUsed(stackName string, typeExternalStack string) ([]string, error) {
	var listStackNotUsed []string

	filteredStack := filterStacks(stackName, typeExternalStack)
	stackInUse, err := getExternalStackInUse(stackName, typeExternalStack)
	if err != nil {
		return nil, err
	}

	for _, currentStack := range filteredStack {
		if currentStack == stackInUse {
			fmt.Println("Current Stack in used : ", currentStack)
		} else {
			listStackNotUsed = append(listStackNotUsed, currentStack)
		}
	}
	return listStackNotUsed, nil
}

func deleteStack(toDelete []string) error {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file. (~/.aws/credentials).
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-1")},
	)
	if err != nil {
		fmt.Println("Error:", err)
	}
	cfsvc := cloudformation.New(sess)

	for _, stackToDelete := range toDelete {
		params := &cloudformation.DeleteStackInput{
			StackName: aws.String(stackToDelete),
		}

		_, errLoop := cfsvc.DeleteStack(params)
		fmt.Println("Delete stacl : ", stackToDelete)

		if errLoop != nil {
			fmt.Println(errLoop)
			break
		}
	}

	return err

}
