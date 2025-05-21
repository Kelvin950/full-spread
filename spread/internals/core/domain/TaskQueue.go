package domain

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Ec2Task struct {
	Bucket string
	Key    string
}

type Ec2TaskState struct {
	State      string    `dynamodbav:"state"`
	StartedAt  time.Time `dynamodbav:"started_at"`
	FinishedAt time.Time `dynamodbav:"finished_at"`
	ErrMsg     string    `dynamodbav:"err_msg"`
	TaskID     string    `dynamodbav:"task_id"`
	Ec2Id      string    `dynamodbav:"ec2_id"`
}

func (e Ec2TaskState) GetKey() (map[string]types.AttributeValue, error) {

	ec2Id, err := attributevalue.Marshal(e.Ec2Id)

	if err != nil {
		return nil, err
	}

	TaskID, err := attributevalue.Marshal(e.Ec2Id)

	if err != nil {
		return nil, err
	}

	return map[string]types.AttributeValue{"ec2_id": ec2Id, "task_id": TaskID}, nil

}
