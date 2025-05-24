package api

import (
	"context"

	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/kelvin950/spread/internals/core/domain"
	"github.com/kelvin950/spread/internals/core/s3"
	"github.com/kelvin950/spread/internals/ports"
	"golang.org/x/sync/errgroup"
)

type IApi interface {
}

var (
	transcode = "Transcode_job"
)

type Api struct {
	S3Client  s3.IS3
	TaskQueue ports.TaskQueue
}

func NewApi(config aws.Config, taskQueue ports.TaskQueue) *Api {

	s3Client := s3.NewS3(config, 2*time.Hour)
	return &Api{
		S3Client:  s3Client,
		TaskQueue: taskQueue,
	}
}

func NewApiMock() *Api {

	return &Api{
		S3Client: s3.NewS3_mock(),
	}
}

func NewApItest(s s3.IS3) *Api {

	return &Api{
		S3Client: s,
	}
}

func (a Api) CreateMultiPartUpload(data domain.CreateMultiPartUpload) (string, error) {

	output, err := a.S3Client.CreateMultiPartUpload(context.Background(), data)

	if err != nil {
		return "", err
	}

	if output == nil {
		return "", nil
	}
	return *output.UploadId, nil
}

func (a Api) CreatePresignMultiPart(data []domain.UplaodMultiPart) ([]domain.UplaodMultiPartApiRes, error) {

	resData := make([]domain.UplaodMultiPartApiRes, len(data)+1)

	errGrp, ctx := errgroup.WithContext(context.Background())
	errGrp.SetLimit(2)

	for _, d := range data {

		func(d domain.UplaodMultiPart) {

			errGrp.Go(func() error {

				resp, err := a.S3Client.CreatePresignMultiPart(ctx, d)
				if err == nil {
					resData[*d.PartNumber] = domain.UplaodMultiPartApiRes{
						Url:        resp.URL,
						PartNumber: *d.PartNumber,
					}
				}

				return err

			})
		}(d)
	}

	return resData, errGrp.Wait()
}

func (a Api) CompleteMultiPart(data domain.CompleteMultiPart) (string, error) {

	output, err := a.S3Client.CompleteMultiPart(context.Background(), data)

	if err != nil {
		return "", err
	}

	a.TaskQueue.DistributeTask(transcode, "critical", domain.Ec2Task{
		Bucket: *data.Bucket,
		Key:    *data.Key,
	})
	return *output.Location, nil

}
