package iam

import (
    "fmt"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/iam"
)

// IAMService provides operations for IAM resources.
type IAMService struct {
    Client *iam.IAM
}

// NewIAMService creates a new IAMService.
func NewIAMService(sess *session.Session) *IAMService {
    return &IAMService{
        Client: iam.New(sess),
    }
}

// CreateUser creates a new IAM user.
func (s *IAMService) CreateUser(userName string) (*iam.CreateUserOutput, error) {
    result, err := s.Client.CreateUser(&iam.CreateUserInput{
        UserName: aws.String(userName),
    })
    if err != nil {
        return nil, fmt.Errorf("failed to create user: %v", err)
    }
    return result, nil
}

// AttachUserPolicy attaches a policy to an IAM user.
func (s *IAMService) AttachUserPolicy(userName, policyArn string) error {
    _, err := s.Client.AttachUserPolicy(&iam.AttachUserPolicyInput{
        UserName:  aws.String(userName),
        PolicyArn: aws.String(policyArn),
    })
    if err != nil {
        return fmt.Errorf("failed to attach policy to user: %v", err)
    }
    return nil
}

// CreateAccessKey creates a new access key for an IAM user.
func (s *IAMService) CreateAccessKey(userName string) (*iam.AccessKey, error) {
    result, err := s.Client.CreateAccessKey(&iam.CreateAccessKeyInput{
        UserName: aws.String(userName),
    })
    if err != nil {
        return nil, fmt.Errorf("failed to create access key: %v", err)
    }
    return result.AccessKey, nil
}

// DeleteUser deletes an IAM user.
func (s *IAMService) DeleteUser(userName string) error {
    _, err := s.Client.DeleteUser(&iam.DeleteUserInput{
        UserName: aws.String(userName),
    })
    if err != nil {
        return fmt.Errorf("failed to delete user: %v", err)
    }
    return nil
}

