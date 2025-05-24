package main

import (
	"context"
	"runtime"

	"log"
	"net/http"
	"os"
	"os/signal"

	"runtime/trace"
	"time"

	awscfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/hibiken/asynq"
	"github.com/kelvin950/spread/config"
	redisclient "github.com/kelvin950/spread/internals/adapters/RedisClient"
	"github.com/kelvin950/spread/internals/adapters/TaskQueue"
	server "github.com/kelvin950/spread/internals/adapters/httpServer"
	dynamoclient "github.com/kelvin950/spread/internals/core/DynamoClient"
	"github.com/kelvin950/spread/internals/core/api"
	"github.com/kelvin950/spread/internals/core/domain"
	ec2client "github.com/kelvin950/spread/internals/core/ec2Client"
	logger "github.com/rs/zerolog/log"
)

func main() {

	f, _ := os.Create("trace.out")
	trace.Start(f)
	runtime.GOMAXPROCS(1)
	defer trace.Stop()

	conf := config.NewConfig()
	port := conf.GetKey("PORT")
	redisAddr := conf.GetKey("REDIS_ADDR")
	redisPassword := conf.GetKey("REDIS_PASSWORD")
	subnetId := conf.GetKey("SUBNET_ID")
	securityGroupId := conf.GetKey("SECURITY_GROUP_ID")
	amiId := conf.GetKey("AMI_ID")
	maxPrice := conf.GetKey("MAX_PRICE")
	awsRole := conf.GetKey("AWS_ROLE")
	tableName := conf.GetKey("TABLE_NAME")
	cfg, err := awscfg.LoadDefaultConfig(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	rClient := redisclient.NewRedisClient(redisAddr, redisPassword)

	ec2Cl := ec2client.NewEc2Client(cfg, subnetId, securityGroupId, amiId, maxPrice, awsRole)
	dynamoCl := dynamoclient.NewDynamoClient(cfg, tableName)

	taskQueue := TaskQueue.NewTask(rClient, ec2Cl, dynamoCl)

	apis := api.NewApi(cfg, taskQueue)

	asynqSrv := asynq.NewServer(rClient, asynq.Config{ // Specify how many concurrent workers to use
		Concurrency: 5,
		// Optionally specify multiple queues with different priority.
		Queues: map[string]int{

			"critical": 6,
			"default":  3,
			"low":      1,
		},
		ErrorHandler: asynq.ErrorHandlerFunc(func(ctx context.Context, task *asynq.Task, err error) {

			logger.Error().Err(err).Str("type", task.Type()).
				Bytes("payload", task.Payload()).Msg("process task failed")
		}),
	})

	asynqMux := asynq.NewServeMux()

	asynqMux.HandleFunc(domain.Transcode, func(ctx context.Context, t *asynq.Task) error {

		return taskQueue.CreateEc2Instance(ctx, t)

	})

	server := server.NewServer(*apis)

	httpServer := http.Server{
		Addr:    port,
		Handler: server.Router,
	}

	done := make(chan error)
	go func() {

		if err := asynqSrv.Run(asynqMux); err != nil && err != http.ErrServerClosed {
			done <- err
		}
		close(done)
	}()

	errCh := make(chan error)
	go func() {

		log.Println("Starting server on port", port)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errCh <- err
		}
		close(errCh)

	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	select {
	case <-sig:
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		log.Println("Shutting down server...")
		if err := httpServer.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}
	case <-done:
		log.Println("asynq server shutting off")
	case err := <-errCh:
		if err != nil {
			log.Fatal(err)
		}

	}

}
