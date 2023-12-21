# c2loud
The goal is to make a common library to use any cloud service


sameple use file

main.go ->

package main

import (
    "fmt"
    "github.com/Akshay-Verma-CS/c2loud/cloud/aws"
)

func main() {
    // Configuration for AWSProvider
    config := &aws.AWSConfig{
        UseIAMRole:  false, // or true if you are using an IAM role
        AccessKey:   "YOUR_ACCESS_KEY",
        SecretKey:   "YOUR_SECRET_KEY",
        Region:      "us-west-2",
        // If using IAM role, set IAMRoleARN and SessionName
        // IAMRoleARN:  "arn:aws:iam::123456789012:role/your-role",
        // SessionName: "your-session-name",
    }

    // Initialize AWSProvider
    awsProvider, err := aws.GetInstance(config)
    if err != nil {
        fmt.Println("Error initializing AWSProvider:", err)
        return
    }

    // Example: Using S3 service to list buckets
    buckets, err := awsProvider.S3Service.ListBuckets()
    if err != nil {
        fmt.Println("Error listing S3 buckets:", err)
        return
    }

    fmt.Println("S3 Buckets:", buckets)
}

