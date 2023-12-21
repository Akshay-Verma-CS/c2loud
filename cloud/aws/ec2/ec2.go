package ec2

import (
    "fmt"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ec2"
)

// EC2Service provides operations for EC2 resources.
type EC2Service struct {
    Client *ec2.EC2
}

// NewEC2Service creates a new EC2Service.
func NewEC2Service(sess *session.Session) *EC2Service {
    return &EC2Service{
        Client: ec2.New(sess),
    }
}

// LaunchInstance launches a new EC2 instance.
func (s *EC2Service) LaunchInstance(imageID, instanceType string) (*ec2.Reservation, error) {
    input := &ec2.RunInstancesInput{
        ImageId:      aws.String(imageID),
        InstanceType: aws.String(instanceType),
        MinCount:     aws.Int64(1),
        MaxCount:     aws.Int64(1),
    }

    result, err := s.Client.RunInstances(input)
    if err != nil {
        return nil, fmt.Errorf("failed to launch instance: %v", err)
    }
    return result, nil
}

// DescribeInstances describes EC2 instances.
func (s *EC2Service) DescribeInstances(instanceIDs []string) ([]*ec2.Instance, error) {
    input := &ec2.DescribeInstancesInput{
        InstanceIds: aws.StringSlice(instanceIDs),
    }

    result, err := s.Client.DescribeInstances(input)
    if err != nil {
        return nil, fmt.Errorf("failed to describe instances: %v", err)
    }

    var instances []*ec2.Instance
    for _, reservation := range result.Reservations {
        instances = append(instances, reservation.Instances...)
    }
    return instances, nil
}

// TerminateInstances terminates one or more EC2 instances.
func (s *EC2Service) TerminateInstances(instanceIDs []string) (*ec2.TerminateInstancesOutput, error) {
    input := &ec2.TerminateInstancesInput{
        InstanceIds: aws.StringSlice(instanceIDs),
    }

    result, err := s.Client.TerminateInstances(input)
    if err != nil {
        return nil, fmt.Errorf("failed to terminate instances: %v", err)
    }
    return result, nil
}

