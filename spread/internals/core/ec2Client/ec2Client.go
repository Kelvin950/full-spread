package ec2client

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

type Ec2Client struct {
	SubnetId        string
	SecurityGroupId string
	AmiId           string
	MaxPrice        string
	Role            string
	Client          *ec2.Client
}

func NewEc2Client(cfg aws.Config, SubnetId, SecurityGroupId, amiid, maxprice, role string) *Ec2Client {

	client := ec2.NewFromConfig(cfg)
	return &Ec2Client{
		SubnetId:        SubnetId,
		SecurityGroupId: SecurityGroupId,
		AmiId:           amiid,
		MaxPrice:        maxprice,
		Role:            role,
		Client:          client,
	}
}

func (e Ec2Client) CreateInstance() (*ec2.RunInstancesOutput, error) {

	return e.Client.RunInstances(context.Background(), &ec2.RunInstancesInput{
		ImageId:          aws.String(e.AmiId),
		SubnetId:         aws.String(e.SubnetId),
		SecurityGroupIds: []string{e.SecurityGroupId},
		InstanceType:     types.InstanceTypeC6a16xlarge,
		InstanceMarketOptions: &types.InstanceMarketOptionsRequest{
			SpotOptions: &types.SpotMarketOptions{
				MaxPrice:                     aws.String(e.MaxPrice),
				InstanceInterruptionBehavior: types.InstanceInterruptionBehaviorTerminate,
				SpotInstanceType:             types.SpotInstanceTypeOneTime,
			},
		},
		BlockDeviceMappings: []types.BlockDeviceMapping{
			{
				DeviceName: aws.String("dev/sda1"),
				Ebs: &types.EbsBlockDevice{
					DeleteOnTermination: aws.Bool(true),
					VolumeType:          types.VolumeTypeGp3,
					VolumeSize:          aws.Int32(100),
					Encrypted:           aws.Bool(true),
				},
			},
		},
		IamInstanceProfile: &types.IamInstanceProfileSpecification{
			Name: aws.String(e.Role),
		},
	})
}

func (e Ec2Client) DestroyInstance(instanceId string) error {
	_, err := e.Client.TerminateInstances(context.Background(), &ec2.TerminateInstancesInput{
		InstanceIds: []string{instanceId},
	})

	return err
}
