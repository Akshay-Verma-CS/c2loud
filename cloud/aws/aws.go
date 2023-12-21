package aws

import (
    "github.com/Akshay-Verma-CS/c2loud/cloud/aws/appconfig"
    "github.com/Akshay-Verma-CS/c2loud/cloud/aws/ec2"
    "github.com/Akshay-Verma-CS/c2loud/cloud/aws/iam"
    "github.com/Akshay-Verma-CS/c2loud/cloud/aws/s3"
    "github.com/Akshay-Verma-CS/c2loud/cloud/aws/sns"
    "github.com/Akshay-Verma-CS/c2loud/cloud/aws/sqs"
    "github.com/aws/aws-sdk-go/aws/session"
	"sync"
)

var (
    // instances holds the map of AWSProvider instances
    instances = make(map[string]*AWSProvider)
    // mutex is used to ensure thread-safe access to the instances map
    mutex sync.Mutex
)

// AWSProvider holds the clients for various AWS services.
type AWSProvider struct {
    S3Service        *s3.S3Service
    SQSService       *sqs.SQSService
    SNSService       *sns.SNSService
    IAMService       *iam.IAMService
    EC2Service       *ec2.EC2Service
    AppConfigService *appconfig.AppConfigService
}

// AWSConfig defines the configuration for AWSProvider.
type AWSConfig struct {
    UseIAMRole     bool
    AccessKey      string
    SecretKey      string
    IAMRoleARN     string
    SessionName    string
    Region         string
}

func NewAWSProvider(config *AWSConfig) (*AWSProvider, error) {
    var sess *session.Session
    var err error

    if config.UseIAMRole {
        // Assume an IAM role if specified
        sess, err = session.NewSession(&aws.Config{Region: aws.String(config.Region)})
        if err == nil {
            creds := stscreds.NewCredentials(sess, config.IAMRoleARN, func(p *stscreds.AssumeRoleProvider) {
                p.RoleSessionName = config.SessionName
            })
            sess.Config.Credentials = creds
        }
    } else {
        // Use access keys if provided
        sess, err = session.NewSession(&aws.Config{
            Region:      aws.String(config.Region),
            Credentials: credentials.NewStaticCredentials(config.AccessKey, config.SecretKey, ""),
        })
    }

    if err != nil {
        return nil, err
    }

    return &AWSProvider{
        S3Service:        s3.NewS3Service(sess),
        SQSService:       sqs.NewSQSService(sess),
        SNSService:       sns.NewSNSService(sess),
        IAMService:       iam.NewIAMService(sess),
        EC2Service:       ec2.NewEC2Service(sess),
        AppConfigService: appconfig.NewAppConfigService(sess),
    }, nil
}

func GetInstance(key string) (*AWSProvider, error) {
    mutex.Lock()
    defer mutex.Unlock()

    if instance, exists := instances[key]; exists {
        return instance, nil
    }

    // Create a new instance if it does not exist
    instance, err := newAWSProvider()
    if err != nil {
        return nil, err
    }

    instances[key] = instance
    return instance, nil
}

func (p *AWSProvider) LaunchEC2Instance(imageID, instanceType string) (*ec2.Reservation, error) {
    return p.EC2Service.LaunchInstance(imageID, instanceType)
}

// CreateS3Bucket creates a new S3 bucket.
func (p *AWSProvider) CreateS3Bucket(bucketName string) error {
    _, err := p.S3Service.CreateBucket(bucketName)
    return err
}

// SendMessageToSQS sends a message to the specified SQS queue.
func (p *AWSProvider) SendMessageToSQS(queueURL, messageBody string) error {
    return p.SQSService.SendMessage(queueURL, messageBody)
}

// CreateSNSTopic creates a new SNS topic.
func (p *AWSProvider) CreateSNSTopic(topicName string) (string, error) {
    return p.SNSService.CreateTopic(topicName)
}

// CreateUserIAM creates a new IAM user.
func (p *AWSProvider) CreateUserIAM(userName string) error {
    _, err := p.IAMService.CreateUser(userName)
    return err
}

// GetAppConfig retrieves the configuration for a specified AppConfig profile.
func (p *AWSProvider) GetAppConfig(applicationID, environmentID, configurationProfileID, clientID string) (*appconfig.GetConfigurationOutput, error) {
    return p.AppConfigService.GetConfiguration(applicationID, environmentID, configurationProfileID, clientID)
}

func (p *AWSProvider) DeleteEC2Instance(instanceID string) error {
    _, err := p.EC2Service.TerminateInstances([]string{instanceID})
    return err
}

// ListS3Buckets lists all S3 buckets.
func (p *AWSProvider) ListS3Buckets() ([]string, error) {
    return p.S3Service.ListBuckets()
}

// ReceiveMessageFromSQS receives messages from the specified SQS queue.
func (p *AWSProvider) ReceiveMessageFromSQS(queueURL string) ([]*sqs.Message, error) {
    return p.SQSService.ReceiveMessage(queueURL)
}

// DeleteSNSTopic deletes an SNS topic by its ARN.
func (p *AWSProvider) DeleteSNSTopic(topicArn string) error {
    return p.SNSService.DeleteTopic(topicArn)
}

// AttachPolicyToUserIAM attaches a policy to an IAM user.
func (p *AWSProvider) AttachPolicyToUserIAM(userName, policyArn string) error {
    return p.IAMService.AttachUserPolicy(userName, policyArn)
}

// UpdateAppConfigProfile updates a configuration profile in AppConfig.
func (p *AWSProvider) UpdateAppConfigProfile(applicationID, profileID, name, locationURI string) error {
    _, err := p.AppConfigService.UpdateConfigurationProfile(applicationID, profileID, name, locationURI)
    return err
}

// ListIAMUsers lists all IAM users.
func (p *AWSProvider) ListIAMUsers() ([]*iam.User, error) {
    result, err := p.IAMService.ListUsers(&iam.ListUsersInput{})
    if err != nil {
        return nil, err
    }
    return result.Users, nil
}

// PublishMessageToSNS publishes a message to the specified SNS topic.
func (p *AWSProvider) PublishMessageToSNS(topicArn, message string) error {
    _, err := p.SNSService.Publish(topicArn, message)
    return err
}

// StartEC2Instance starts an EC2 instance by its ID.
func (p *AWSProvider) StartEC2Instance(instanceID string) error {
    _, err := p.EC2Service.StartInstances(&ec2.StartInstancesInput{
        InstanceIds: []*string{&instanceID},
    })
    return err
}

// StopEC2Instance stops an EC2 instance by its ID.
func (p *AWSProvider) StopEC2Instance(instanceID string) error {
    _, err := p.EC2Service.StopInstances(&ec2.StopInstancesInput{
        InstanceIds: []*string{&instanceID},
    })
    return err
}

// CreateSQSQueue creates a new SQS queue.
func (p *AWSProvider) CreateSQSQueue(queueName string) (string, error) {
    return p.SQSService.CreateQueue(queueName)
}

// DeleteSQSQueue deletes an SQS queue by its URL.
func (p *AWSProvider) DeleteSQSQueue(queueURL string) error {
    return p.SQSService.DeleteQueue(queueURL)
}

// SubscribeToSNSTopic subscribes an endpoint to an SNS topic.
func (p *AWSProvider) SubscribeToSNSTopic(topicArn, protocol, endpoint string) (string, error) {
    return p.SNSService.Subscribe(topicArn, protocol, endpoint)
}

// UnsubscribeFromSNSTopic unsubscribes an endpoint from an SNS topic.
func (p *AWSProvider) UnsubscribeFromSNSTopic(subscriptionArn string) error {
    return p.SNSService.Unsubscribe(subscriptionArn)
}

// CreateIAMAccessKey creates a new access key for an IAM user.
func (p *AWSProvider) CreateIAMAccessKey(userName string) (*iam.AccessKey, error) {
    return p.IAMService.CreateAccessKey(userName)
}

// DeleteIAMAccessKey deletes an access key for an IAM user.
func (p *AWSProvider) DeleteIAMAccessKey(userName, accessKeyID string) error {
    _, err := p.IAMService.DeleteAccessKey(&iam.DeleteAccessKeyInput{
        UserName:    &userName,
        AccessKeyId: &accessKeyID,
    })
    return err
}

func (p *AWSProvider) DescribeEC2Instances() ([]*ec2.Instance, error) {
    input := &ec2.DescribeInstancesInput{}
    result, err := p.EC2Service.DescribeInstances(input)
    if err != nil {
        return nil, err
    }

    var instances []*ec2.Instance
    for _, reservation := range result.Reservations {
        instances = append(instances, reservation.Instances...)
    }
    return instances, nil
}

// DeleteS3Bucket - Deletes an S3 bucket.
func (p *AWSProvider) DeleteS3Bucket(bucketName string) error {
    _, err := p.S3Service.DeleteBucket(&s3.DeleteBucketInput{
        Bucket: &bucketName,
    })
    return err
}

// ListSQSQueues - Lists SQS queues.
func (p *AWSProvider) ListSQSQueues() ([]string, error) {
    result, err := p.SQSService.ListQueues(&sqs.ListQueuesInput{})
    if err != nil {
        return nil, err
    }
    return result.QueueUrls, nil
}

func (p *AWSProvider) RebootEC2Instance(instanceID string) error {
    input := &ec2.RebootInstancesInput{
        InstanceIds: []*string{&instanceID},
    }
    _, err := p.EC2Service.RebootInstances(input)
    return err
}

// ListSNSTopics - Lists all SNS topics.
func (p *AWSProvider) ListSNSTopics() ([]*sns.Topic, error) {
    input := &sns.ListTopicsInput{}
    result, err := p.SNSService.ListTopics(input)
    if err != nil {
        return nil, err
    }

    return result.Topics, nil
}

// DeleteIAMUser - Deletes an IAM user.
func (p *AWSProvider) DeleteIAMUser(userName string) error {
    input := &iam.DeleteUserInput{
        UserName: &userName,
    }
    _, err := p.IAMService.DeleteUser(input)
    return err
}
