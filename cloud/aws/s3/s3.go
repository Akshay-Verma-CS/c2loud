package s3

import (
    "fmt"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
)

// S3Service provides operations for S3 resources.
type S3Service struct {
    Client *s3.S3
}

// NewS3Service creates a new S3Service.
func NewS3Service(sess *session.Session) *S3Service {
    return &S3Service{
        Client: s3.New(sess),
    }
}

// CreateBucket creates a new S3 bucket.
func (s *S3Service) CreateBucket(bucketName string) error {
    _, err := s.Client.CreateBucket(&s3.CreateBucketInput{
        Bucket: aws.String(bucketName),
    })
    if err != nil {
        return fmt.Errorf("failed to create bucket: %v", err)
    }
    return nil
}

// DeleteBucket deletes an S3 bucket.
func (s *S3Service) DeleteBucket(bucketName string) error {
    _, err := s.Client.DeleteBucket(&s3.DeleteBucketInput{
        Bucket: aws.String(bucketName),
    })
    if err != nil {
        return fmt.Errorf("failed to delete bucket: %v", err)
    }
    return nil
}

// ListBuckets lists all S3 buckets.
func (s *S3Service) ListBuckets() ([]string, error) {
    result, err := s.Client.ListBuckets(nil)
    if err != nil {
        return nil, fmt.Errorf("failed to list buckets: %v", err)
    }

    var buckets []string
    for _, b := range result.Buckets {
        buckets = append(buckets, *b.Name)
    }
    return buckets, nil
}

// GetBucketInfo retrieves information about an S3 bucket.
func (s *S3Service) GetBucketInfo(bucketName string) (*BucketInfo, error) {
    // Get basic bucket information (creation date)
    result, err := s.Client.ListBuckets(nil)
    if err != nil {
        return nil, fmt.Errorf("failed to list buckets: %v", err)
    }

    var bucketInfo BucketInfo
    for _, b := range result.Buckets {
        if *b.Name == bucketName {
            bucketInfo.Name = bucketName
            bucketInfo.CreationDate = b.CreationDate
            break
        }
    }

    // Get bucket location
    loc, err := s.Client.GetBucketLocation(&s3.GetBucketLocationInput{Bucket: aws.String(bucketName)})
    if err != nil {
        return nil, fmt.Errorf("failed to get bucket location for %s: %v", bucketName, err)
    }
    bucketInfo.Location = aws.StringValue(loc.LocationConstraint)

    // Get versioning information
    versioning, err := s.Client.GetBucketVersioning(&s3.GetBucketVersioningInput{Bucket: aws.String(bucketName)})
    if err != nil {
        return nil, fmt.Errorf("failed to get bucket versioning for %s: %v", bucketName, err)
    }
    bucketInfo.Versioning = aws.StringValue(versioning.Status)

    // Add more API calls as needed to fetch other information

    return &bucketInfo, nil
}

func (s *S3Service) UploadFile(bucketName, key, filePath string) error {
    file, err := os.Open(filePath)
    if err != nil {
        return fmt.Errorf("failed to open file %q, %v", filePath, err)
    }
    defer file.Close()

    _, err = s.Client.PutObject(&s3.PutObjectInput{
        Bucket: aws.String(bucketName),
        Key:    aws.String(key),
        Body:   file,
    })
    if err != nil {
        return fmt.Errorf("failed to put file %q in bucket %q, %v", key, bucketName, err)
    }
    return nil
}

// DownloadFile downloads a file from an S3 bucket.
func (s *S3Service) DownloadFile(bucketName, key, filePath string) error {
    outputFile, err := os.Create(filePath)
    if err != nil {
        return fmt.Errorf("failed to create file %q, %v", filePath, err)
    }
    defer outputFile.Close()

    resp, err := s.Client.GetObject(&s3.GetObjectInput{
        Bucket: aws.String(bucketName),
        Key:    aws.String(key),
    })
    if err != nil {
        return fmt.Errorf("failed to get file %q from bucket %q, %v", key, bucketName, err)
    }
    defer resp.Body.Close()

    _, err = io.Copy(outputFile, resp.Body)
    if err != nil {
        return fmt.Errorf("failed to copy file %q, %v", filePath, err)
    }
    return nil
}

// ListObjects lists the objects in an S3 bucket.
func (s *S3Service) ListObjects(bucketName string) ([]string, error) {
    resp, err := s.Client.ListObjectsV2(&s3.ListObjectsV2Input{
        Bucket: aws.String(bucketName),
    })
    if err != nil {
        return nil, fmt.Errorf("failed to list objects in bucket %q, %v", bucketName, err)
    }

    var objects []string
    for _, item := range resp.Contents {
        objects = append(objects, *item.Key)
    }
    return objects, nil
}

// DeleteObject deletes an object from an S3 bucket.
func (s *S3Service) DeleteObject(bucketName, key string) error {
    _, err := s.Client.DeleteObject(&s3.DeleteObjectInput{
        Bucket: aws.String(bucketName),
        Key:    aws.String(key),
    })
    if err != nil {
        return fmt.Errorf("failed to delete object %q from bucket %q, %v", key, bucketName, err)
    }
    return nil
}
