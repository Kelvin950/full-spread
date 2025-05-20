package ports 


type TaskQueue interface{
	DistributeTask(taskName , priority  string , taskpayload interface{})error
}