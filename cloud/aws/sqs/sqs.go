package sqs

import (
    "fmt"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/sqs"
)

// SQSService provides operations for SQS resources.
type SQSService struct {
    Client *sqs.SQS
}

// NewSQSService creates a new SQSService.
func NewSQSService(sess *session.Session) *SQSService {
    return &SQSService{
        Client: sqs.New(sess),
    }
}

// CreateQueue creates a new SQS queue.
func (s *SQSService) CreateQueue(queueName string) (string, error) {
    result, err := s.Client.CreateQueue(&sqs.CreateQueueInput{
        QueueName: aws.String(queueName),
    })
    if err != nil {
        return "", fmt.Errorf("failed to create queue: %v", err)
    }
    return aws.StringValue(result.QueueUrl), nil
}

// SendMessage sends a message to an SQS queue.
func (s *SQSService) SendMessage(queueURL, messageBody string) error {
    _, err := s.Client.SendMessage(&sqs.SendMessageInput{
        QueueUrl:    aws.String(queueURL),
        MessageBody: aws.String(messageBody),
    })
    if err != nil {
        return fmt.Errorf("failed to send message: %v", err)
    }
    return nil
}

// ReceiveMessage receives messages from an SQS queue.
func (s *SQSService) ReceiveMessage(queueURL string) ([]*sqs.Message, error) {
    result, err := s.Client.ReceiveMessage(&sqs.ReceiveMessageInput{
        QueueUrl: aws.String(queueURL),
    })
    if err != nil {
        return nil, fmt.Errorf("failed to receive messages: %v", err)
    }
    return result.Messages, nil
}

// DeleteQueue deletes an SQS queue.
func (s *SQSService) DeleteQueue(queueURL string) error {
    _, err := s.Client.DeleteQueue(&sqs.DeleteQueueInput{
        QueueUrl: aws.String(queueURL),
    })
    if err != nil {
        return fmt.Errorf("failed to delete queue: %v", err)
    }
    return nil
}

// DeleteMessage deletes a message from the SQS queue.
func (s *SQSService) DeleteMessage(queueURL, receiptHandle string) error {
    _, err := s.Client.DeleteMessage(&sqs.DeleteMessageInput{
        QueueUrl:      aws.String(queueURL),
        ReceiptHandle: aws.String(receiptHandle),
    })
    if err != nil {
        return fmt.Errorf("failed to delete message: %v", err)
    }
    return nil
}

// ChangeMessageVisibility sets the visibility timeout for a specific message.
func (s *SQSService) ChangeMessageVisibility(queueURL, receiptHandle string, visibilityTimeout int64) error {
    _, err := s.Client.ChangeMessageVisibility(&sqs.ChangeMessageVisibilityInput{
        QueueUrl:          aws.String(queueURL),
        ReceiptHandle:     aws.String(receiptHandle),
        VisibilityTimeout: aws.Int64(visibilityTimeout),
    })
    if err != nil {
        return fmt.Errorf("failed to change message visibility: %v", err)
    }
    return nil
}

// GetQueueAttributes retrieves attributes for an SQS queue.
func (s *SQSService) GetQueueAttributes(queueURL string) (map[string]string, error) {
    result, err := s.Client.GetQueueAttributes(&sqs.GetQueueAttributesInput{
        QueueUrl:       aws.String(queueURL),
        AttributeNames: aws.StringSlice([]string{"All"}),
    })
    if err != nil {
        return nil, fmt.Errorf("failed to get queue attributes: %v", err)
    }
    return aws.StringValueMap(result.Attributes), nil
}

// SetQueueAttributes sets attributes for an SQS queue.
func (s *SQSService) SetQueueAttributes(queueURL string, attributes map[string]string) error {
    _, err := s.Client.SetQueueAttributes(&sqs.SetQueueAttributesInput{
        QueueUrl:   aws.String(queueURL),
        Attributes: aws.StringMap(attributes),
    })
    if err != nil {
        return fmt.Errorf("failed to set queue attributes: %v", err)
    }
    return nil
}
