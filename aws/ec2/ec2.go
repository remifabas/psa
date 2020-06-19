package ec2

import (
	"errors"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func getEC2Session() (*ec2.EC2, error) {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file. (~/.aws/credentials).
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-1")},
	)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	ec2Service := ec2.New(sess)

	return ec2Service, nil
}

func getInstances(ec2Sess *ec2.EC2) ([]*ec2.Reservation, error) {

	var allReservation []*ec2.Reservation

	describeInstanceInput := &ec2.DescribeInstancesInput{}
	output, err := ec2Sess.DescribeInstances(describeInstanceInput)
	allReservation = append(allReservation, output.Reservations...)
	if err != nil {
		return nil, err
	}
	if output.NextToken != nil {
		describeInstanceInputLoop := &ec2.DescribeInstancesInput{NextToken: output.NextToken}
		output, err = ec2Sess.DescribeInstances(describeInstanceInputLoop)
		if err != nil {
			return nil, err
		}
		allReservation = append(allReservation, output.Reservations...)
	}
	return allReservation, nil
}

func getMongoDBInstances(list []*ec2.Reservation) ([]*ec2.Reservation, error) {
	var returnRes []*ec2.Reservation
	for _, r := range list {
		tags := r.Instances[0].Tags
		name, errGetValue := getValueTagName(tags)
		if errGetValue != nil {
			return nil, errGetValue
		}
		if strings.Contains(name, "MongoDB") {
			fmt.Println(" [OK] name :", name)
			returnRes = append(returnRes, r)
		}
	}
	if len(returnRes) <= 0 {
		return nil, errors.New("Error, no value found")
	}
	return returnRes, nil
}

func getValueTagName(tags []*ec2.Tag) (string, error) {
	for _, t := range tags {
		if strings.EqualFold(*t.Key, "Name") {
			return *t.Value, nil
		}
	}
	return "", errors.New("No KEY found with 'Name'")
}
