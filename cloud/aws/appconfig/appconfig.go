package appconfig

import (
    "fmt"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/appconfig"
)

// AppConfigService provides operations for AWS AppConfig resources.
type AppConfigService struct {
    Client *appconfig.AppConfig
}

// NewAppConfigService creates a new AppConfigService.
func NewAppConfigService(sess *session.Session) *AppConfigService {
    return &AppConfigService{
        Client: appconfig.New(sess),
    }
}

// CreateConfigurationProfile creates a new configuration profile in AppConfig.
func (s *AppConfigService) CreateConfigurationProfile(applicationID, name, locationURI string) (*appconfig.CreateConfigurationProfileOutput, error) {
    input := &appconfig.CreateConfigurationProfileInput{
        ApplicationId: aws.String(applicationID),
        Name:          aws.String(name),
        LocationUri:   aws.String(locationURI),
    }

    result, err := s.Client.CreateConfigurationProfile(input)
    if err != nil {
        return nil, fmt.Errorf("failed to create configuration profile: %v", err)
    }
    return result, nil
}

// GetConfiguration retrieves the configuration for the specified profile.
func (s *AppConfigService) GetConfiguration(applicationID, environmentID, configurationProfileID, clientID string) (*appconfig.GetConfigurationOutput, error) {
    input := &appconfig.GetConfigurationInput{
        ApplicationId:         aws.String(applicationID),
        EnvironmentId:         aws.String(environmentID),
        ConfigurationProfileId: aws.String(configurationProfileID),
        ClientId:              aws.String(clientID),
    }

    result, err := s.Client.GetConfiguration(input)
    if err != nil {
        return nil, fmt.Errorf("failed to get configuration: %v", err)
    }
    return result, nil
}

func (s *AppConfigService) UpdateConfigurationProfile(applicationID, profileID, name, locationURI string) (*appconfig.UpdateConfigurationProfileOutput, error) {
    input := &appconfig.UpdateConfigurationProfileInput{
        ApplicationId:         aws.String(applicationID),
        ConfigurationProfileId: aws.String(profileID),
        Name:                  aws.String(name),
        LocationUri:           aws.String(locationURI),
    }

    result, err := s.Client.UpdateConfigurationProfile(input)
    if err != nil {
        return nil, fmt.Errorf("failed to update configuration profile: %v", err)
    }
    return result, nil
}

// CreateEnvironment creates a new environment in AppConfig.
func (s *AppConfigService) CreateEnvironment(applicationID, name, description string) (*appconfig.CreateEnvironmentOutput, error) {
    input := &appconfig.CreateEnvironmentInput{
        ApplicationId: aws.String(applicationID),
        Name:          aws.String(name),
        Description:   aws.String(description),
    }

    result, err := s.Client.CreateEnvironment(input)
    if err != nil {
        return nil, fmt.Errorf("failed to create environment: %v", err)
    }
    return result, nil
}

// ValidateConfiguration validates a configuration against a schema or a set of rules.
func (s *AppConfigService) ValidateConfiguration(applicationID, configurationProfileID, configurationVersion string) (*appconfig.ValidateConfigurationOutput, error) {
    input := &appconfig.ValidateConfigurationInput{
        ApplicationId:         aws.String(applicationID),
        ConfigurationProfileId: aws.String(configurationProfileID),
        ConfigurationVersion:   aws.String(configurationVersion),
    }

    result, err := s.Client.ValidateConfiguration(input)
    if err != nil {
        return nil, fmt.Errorf("failed to validate configuration: %v", err)
    }
    return result, nil
}
