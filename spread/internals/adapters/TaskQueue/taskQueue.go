package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/hibiken/asynq"
)

type Itask  interface{

}

type Task  struct {

  client *asynq.Client

}


// func NewTask() *Task {
// 	return &Task{
// 		client: asynq.NewClient(),
// 	}
// }


func(t *Task) DistributeTask(taskName , priority  string , taskpayload interface{})error {

	payload , err:= json.Marshal(&taskpayload) 

	 if err!=nil{
		return err
	 }


	task :=asynq.NewTask(taskName , payload)
	

	info , err:= t.client.Enqueue(task   ,asynq.MaxRetry(3) , asynq.Retention(24 * time.Hour) ,  asynq.ProcessIn(10 * time.Second) , asynq.Queue(priority) )

	if err!=nil{
		return err 
	}

	log.Println(info.ID)
	return  nil 
}

