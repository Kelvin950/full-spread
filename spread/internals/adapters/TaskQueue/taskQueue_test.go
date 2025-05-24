package TaskQueue

import (
	"context"
	"encoding/json"
	"errors"

	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/hibiken/asynq"
	"github.com/kelvin950/spread/internals/core/domain"
	"github.com/kelvin950/spread/internals/ports"
	"github.com/kelvin950/spread/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func NewTaskTest(ec2port ports.Ec2Client, dynamoport ports.DynamoClient) *Task {
	return &Task{
		Ec2:      ec2port,
		Dynamodb: dynamoport,
	}
}

type dummyTask struct {
}

func (d dummyTask) Payload() []byte {
	p, _ := json.Marshal(&domain.Ec2Task{
		Bucket: "dss",
		Key:    "dsd",
	})

	return p

}

func (d dummyTask) ResultWriter() *asynq.ResultWriter {

	return &asynq.ResultWriter{}
}

func (d dummyTask) Type() string {
	return ""
}

func TestProcesTaskError(t *testing.T) {

	ec2Client := mocks.NewEc2Client(t)
	dynamoClient := mocks.NewDynamoClient(t)

	ec2Client.On("CreateInstance", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil, errors.New("failed"))

	taskQueue := NewTaskTest(ec2Client, dynamoClient)
	err := taskQueue.CreateEc2Instance(context.Background(), &dummyTask{})

	require.Error(t, err)

}

func TestProcesTaskDynamoError(t *testing.T) {

	ec2Client := mocks.NewEc2Client(t)
	dynamoClient := mocks.NewDynamoClient(t)

	ec2Client.On("CreateInstance", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(&ec2.RunInstancesOutput{
		Instances: []types.Instance{
			{
				InstanceId: aws.String("fdfd"),
			},
		},
	}, nil)

	dynamoClient.On("PutItem", mock.Anything).Return(errors.New("failed"))

	taskQueue := NewTaskTest(ec2Client, dynamoClient)
	err := taskQueue.CreateEc2Instance(context.Background(), &dummyTask{})

	require.Error(t, err)

}

func TestProcessTaskCheckTimerWhenErrorisCalled(t *testing.T) {

	ec2Client := mocks.NewEc2Client(t)
	dynamoClient := mocks.NewDynamoClient(t)

	ec2Client.On("CreateInstance", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(&ec2.RunInstancesOutput{
		Instances: []types.Instance{
			{
				InstanceId: aws.String("fdfd"),
			},
		},
	}, nil)

	dynamoClient.On("PutItem", mock.Anything).Return(nil)

	dynamoClient.On("GetItem", mock.Anything, mock.Anything).Return(domain.Ec2TaskState{}, errors.New("failed"))

	ec2Client.On("DestroyInstance", mock.AnythingOfType("string")).Return(nil)

	taskQueue := NewTaskTest(ec2Client, dynamoClient)
	err := taskQueue.CreateEc2Instance(context.Background(), &dummyTask{})

	require.Error(t, err)

}

func TestProcessTaskCheckTimerWhenTaskFails(t *testing.T) {

	ec2Client := mocks.NewEc2Client(t)
	dynamoClient := mocks.NewDynamoClient(t)

	ec2Client.On("CreateInstance", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(&ec2.RunInstancesOutput{
		Instances: []types.Instance{
			{
				InstanceId: aws.String("fdfd"),
			},
		},
	}, nil)

	dynamoClient.On("PutItem", mock.Anything).Return(nil)

	dynamoClient.On("GetItem", mock.Anything, mock.Anything).Return(domain.Ec2TaskState{

		State: "failed",
	}, nil)

	ec2Client.On("DestroyInstance", mock.AnythingOfType("string")).Return(nil)

	taskQueue := NewTaskTest(ec2Client, dynamoClient)
	err := taskQueue.CreateEc2Instance(context.Background(), &dummyTask{})

	require.Error(t, err)

}

func TestProcessTaskCheckTimerWhenTaskFinished(t *testing.T) {

	ec2Client := mocks.NewEc2Client(t)
	dynamoClient := mocks.NewDynamoClient(t)

	ec2Client.On("CreateInstance", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(&ec2.RunInstancesOutput{
		Instances: []types.Instance{
			{
				InstanceId: aws.String("fdfd"),
			},
		},
	}, nil)

	dynamoClient.On("PutItem", mock.Anything).Return(nil)

	dynamoClient.On("GetItem", mock.Anything, mock.Anything).Return(domain.Ec2TaskState{

		State: "finished",
	}, nil)

	ec2Client.On("DestroyInstance", mock.AnythingOfType("string")).Return(nil)

	taskQueue := NewTaskTest(ec2Client, dynamoClient)

	err := taskQueue.CreateEc2Instance(context.Background(), &dummyTask{})

	require.NoError(t, err)

}
