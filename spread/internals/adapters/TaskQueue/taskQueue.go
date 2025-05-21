package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/hibiken/asynq"
	"github.com/kelvin950/spread/internals/core/domain"
	"github.com/kelvin950/spread/internals/ports"
)

type Itask interface {
}

type Task struct {
	Client      *asynq.Client
	RedisClient ports.RedisClient
}

func NewTask(r ports.RedisClient) *Task {
	return &Task{
		Client: asynq.NewClient(r),
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

func (t *Task) CreateEc2Instance(ctx context.Context, ta *asynq.Task) error {

	var payload domain.Ec2Task

	err := json.Unmarshal(ta.Payload(), &payload)

	return err

	//create ec2 task

	//write to dynamodb

	//track state

	//save to db
}
