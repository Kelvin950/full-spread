package TaskQueue

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/hibiken/asynq"
	"github.com/kelvin950/spread/internals/core/domain"
	"github.com/kelvin950/spread/internals/ports"
	"github.com/kelvin950/spread/mocks"
	"github.com/stretchr/testify/require"
)


func NewTaskTest(ec2port ports.Ec2Client , dynamoport ports.DynamoClient) *Task{
  return &Task{
	Ec2: ec2port,
	Dynamodb: dynamoport,
  }
}


type dummyTask struct{

}


func(d dummyTask)Payload()[]byte{
	p,_:=json.Marshal(&domain.Ec2Task{
		Bucket: "dss",
		Key: "dsd",
	})

	return p

}


func(d dummyTask) ResultWriter() *asynq.ResultWriter{
	return nil
}

func (d dummyTask)Type()string{
	return ""
}


func TestProcesTaskError(t *testing.T){
   
	ec2Client:= mocks.NewEc2Client(t)
	dynamoClient:= mocks.NewDynamoClient(t) 

	ec2Client.On("CreateInstance").Return(nil , errors.New("failed"))
  
	taskQueue:= NewTaskTest(ec2Client , dynamoClient) 
	err:= taskQueue.CreateEc2Instance(context.Background() ,&dummyTask{} )

	require.Error(t , err) 

}