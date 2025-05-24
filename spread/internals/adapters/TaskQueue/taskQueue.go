package TaskQueue

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/hibiken/asynq"
	"github.com/kelvin950/spread/internals/core/domain"
	"github.com/kelvin950/spread/internals/ports"
)

type Itask interface {
	Payload() []byte
	ResultWriter() *asynq.ResultWriter
	Type() string
}

type Task struct {
	Client      *asynq.Client
	RedisClient ports.RedisClient
	Ec2         ports.Ec2Client
	Dynamodb    ports.DynamoClient
}

func NewTask(r ports.RedisClient, ec2port ports.Ec2Client, dynamoport ports.DynamoClient) *Task {

	return &Task{
		Client:   asynq.NewClient(r),
		Ec2:      ec2port,
		Dynamodb: dynamoport,
	}
}

func (t *Task) DistributeTask(taskName, priority string, taskpayload interface{}) error {

	payload, err := json.Marshal(&taskpayload)

	if err != nil {
		return err
	}

	task := asynq.NewTask(taskName, payload)

	info, err := t.Client.Enqueue(task, asynq.MaxRetry(3), asynq.Retention(24*time.Hour), asynq.ProcessIn(10*time.Second), asynq.Queue(priority))

	if err != nil {
		return err
	}

	log.Println(info.ID)
	return nil
}

func (t *Task) CreateEc2Instance(ctx context.Context, ta Itask) error {

	var payload domain.Ec2Task

	err := json.Unmarshal(ta.Payload(), &payload)

	if err != nil {
		return err
	}

	//create ec2 task
	timestarted := time.Now().Format("20060102T150405")
	ec2Output, err := t.Ec2.CreateInstance(timestarted, ta.ResultWriter().TaskID())

	if err != nil {
		return err
	}

	//write to dynamodb
	err = t.Dynamodb.PutItem(domain.Ec2TaskState{
		State:     "started",
		StartedAt: timestarted,
		TaskID:    ta.ResultWriter().TaskID(),
		Ec2Id:     *ec2Output.Instances[0].InstanceId,
	})

	if err != nil {
		return err
	}

	timer := time.NewTicker(2 * time.Second)
	//track state
	errCh := make(chan error, 1)
	succesChan := make(chan struct{}, 1)

label:
	for {

		select {
		case <-timer.C:
			taskState, err := t.Dynamodb.GetItem(timestarted, ta.ResultWriter().TaskID())
			if err != nil {
				log.Print("dser")
				errCh <- err
			} else {

				if taskState.State == "failed" {
					errCh <- errors.New("task failed")
				} else if taskState.State == "finished" {
					succesChan <- struct{}{}

				}
			}

			//do something
		case err := <-errCh:

			// shut instance down return err
			t.Ec2.DestroyInstance(*ec2Output.Instances[0].InstanceId)
			timer.Stop()
			return err
		case <-succesChan:
			//shut instance down break
			t.Ec2.DestroyInstance(*ec2Output.Instances[0].InstanceId)
			timer.Stop()
			break label
		}

	}
	//save to db
	fmt.Println("writing to db")
	return nil
}
