package ports

import "github.com/aws/aws-sdk-go-v2/service/ec2"

type Ec2Client interface {
	CreateInstance() (*ec2.RunInstancesOutput, error) 
	DestroyInstance(instanceId string) error
}
