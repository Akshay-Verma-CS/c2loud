package sns

import (
    "fmt"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/sns"
)

// SNSService provides operations for SNS resources.
type SNSService struct {
    Client *sns.SNS
}

// NewSNSService creates a new SNSService.
func NewSNSService(sess *session.Session) *SNSService {
    return &SNSService{
        Client: sns.New(sess),
    }
}

// CreateTopic creates a new SNS topic.
func (s *SNSService) CreateTopic(topicName string) (string, error) {
    result, err := s.Client.CreateTopic(&sns.CreateTopicInput{
        Name: aws.String(topicName),
    })
    if err != nil {
        return "", fmt.Errorf("failed to create topic: %v", err)
    }
    return aws.StringValue(result.TopicArn), nil
}

// Subscribe subscribes an endpoint to an SNS topic.
func (s *SNSService) Subscribe(topicArn, protocol, endpoint string) (string, error) {
    result, err := s.Client.Subscribe(&sns.SubscribeInput{
        TopicArn: aws.String(topicArn),
        Protocol: aws.String(protocol),
        Endpoint: aws.String(endpoint),
    })
    if err != nil {
        return "", fmt.Errorf("failed to subscribe to topic: %v", err)
    }
    return aws.StringValue(result.SubscriptionArn), nil
}

// Publish sends a message to an SNS topic.
func (s *SNSService) Publish(topicArn, message string) (string, error) {
    result, err := s.Client.Publish(&sns.PublishInput{
        TopicArn: aws.String(topicArn),
        Message:  aws.String(message),
    })
    if err != nil {
        return "", fmt.Errorf("failed to publish message: %v", err)
    }
    return aws.StringValue(result.MessageId), nil
}

// DeleteTopic deletes an SNS topic.
func (s *SNSService) DeleteTopic(topicArn string) error {
    _, err := s.Client.DeleteTopic(&sns.DeleteTopicInput{
        TopicArn: aws.String(topicArn),
    })
    if err != nil {
        return fmt.Errorf("failed to delete topic: %v", err)
    }
    return nil
}

