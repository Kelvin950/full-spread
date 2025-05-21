package dynamoclient

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/kelvin950/spread/internals/core/domain"
)

type DynamoClient struct {
	Table  string
	Client *dynamodb.Client
}

func NewDynamoClient(cfg aws.Config, tablename string) *DynamoClient {

	client := dynamodb.NewFromConfig(cfg)

	return &DynamoClient{
		Table:  tablename,
		Client: client,
	}
}

func (d DynamoClient) PutItem(item domain.Ec2Task) error {

	av, err := attributevalue.MarshalMap(item)

	if err != nil {
		return err
	}

	_, err = d.Client.PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: aws.String(d.Table), Item: av,
	})
	return err
}

func (d DynamoClient) GetItem(ec2Id, TaskId string) (domain.Ec2TaskState, error) {

	task := domain.Ec2TaskState{TaskID: TaskId, Ec2Id: ec2Id}

	taskmap, err := task.GetKey()

	if err != nil {
		return domain.Ec2TaskState{}, err
	}

	output, err := d.Client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(d.Table),
		Key:       taskmap,
	})

	if err != nil {
		return domain.Ec2TaskState{}, err
	}

	var ec2TaskState domain.Ec2TaskState

	err = attributevalue.UnmarshalMap(output.Item, &ec2TaskState)

	return ec2TaskState, err

}
