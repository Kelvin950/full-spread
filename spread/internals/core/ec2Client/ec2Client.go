package ec2client

import (
	"context"
	"fmt"

	"encoding/base64"

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

func (e Ec2Client) CreateInstance(timestarted, taskid string) (*ec2.RunInstancesOutput, error) {

	script := fmt.Sprintf(`#!/bin/bash
docker run -d -p 80:80 nginx

# Run your custom transcode container
{
  docker run -d -p 80:80 nginx

  docker run \
    -e key=video.mp4 \
    -e bucket=testbucketkab \
    -e path=/usr/bin/ffmpeg \
    -e AWS_REGION=us-east-1 \
    -e taskid=%s \
    -e timestarted=%s \
    letsgo21/transcode:cpu

} > /var/log/user-data.log 2>&1
`, taskid, timestarted)
	return e.Client.RunInstances(context.Background(), &ec2.RunInstancesInput{
		ImageId: aws.String(e.AmiId),

		SecurityGroupIds: []string{e.SecurityGroupId},
		InstanceType:     types.InstanceTypeC6a16xlarge,
		MaxCount:         aws.Int32(1),
		MinCount:         aws.Int32(1),
		KeyName:          aws.String("test-key"),
		InstanceMarketOptions: &types.InstanceMarketOptionsRequest{
			MarketType: types.MarketTypeSpot,
			SpotOptions: &types.SpotMarketOptions{
				MaxPrice:                     aws.String(e.MaxPrice),
				InstanceInterruptionBehavior: types.InstanceInterruptionBehaviorTerminate,
				SpotInstanceType:             types.SpotInstanceTypeOneTime,
			},
		},
		BlockDeviceMappings: []types.BlockDeviceMapping{
			{
				DeviceName: aws.String("/dev/sda1"),
				Ebs: &types.EbsBlockDevice{
					DeleteOnTermination: aws.Bool(true),
					VolumeType:          types.VolumeTypeIo2,
					Iops:                aws.Int32(25000),  
					VolumeSize:          aws.Int32(100),
					Encrypted:           aws.Bool(true),
				},
			},
		},
		IamInstanceProfile: &types.IamInstanceProfileSpecification{
			Name: aws.String(e.Role),
		},
		UserData: aws.String(base64.StdEncoding.EncodeToString([]byte(script))),
	})
}

func (e Ec2Client) DestroyInstance(instanceId string) error {
	_, err := e.Client.TerminateInstances(context.Background(), &ec2.TerminateInstancesInput{
		InstanceIds: []string{instanceId},
	})

	return err
}
