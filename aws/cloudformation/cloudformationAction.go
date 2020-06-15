package cloudformation

import (
	"fmt"
	"os"
)

// ActionDeleteStack delete all unused stack from stackName with type ExternalStack
func ActionDeleteStack(stackName string, typeExternalStack string) {
	stackToDelete, err := listStackNotUsed(stackName, typeExternalStack)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Number of not used stack : ", len(stackToDelete))
	err = deleteStack(stackToDelete)
	if err != nil {
		fmt.Println("Expected no error")
	}
}
