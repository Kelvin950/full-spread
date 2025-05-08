package api

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/kelvin950/spread/internals/core/domain"
	"github.com/kelvin950/spread/internals/core/s3"
)

type IApi interface {
}

type Api struct {
	S3Client s3.IS3
}

func NewApi(config aws.Config) *Api {

	s3Client := s3.NewS3(config, 2*time.Hour)
	return &Api{
		S3Client: s3Client,
	}
}

func NewApiMock() *Api {

	return &Api{
		S3Client: s3.NewS3_mock(),
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

	var wg sync.WaitGroup
	errCh := make(chan error, len(data))
	resCh := make(chan domain.UplaodMultiPartApiRes, len(data))
	resData := make([]domain.UplaodMultiPartApiRes, 0, len(data))
	concurrencyLimit := 20
	semaphore := make(chan struct{}, concurrencyLimit)

	for _, d := range data {
		d := d
		wg.Add(1)

		go func(u domain.UplaodMultiPart) {
			defer wg.Done()

			semaphore <- struct{}{}

			output, err := a.S3Client.CreatePresignMultiPart(context.Background(), u)
			fmt.Println(err)
			if err != nil {

				errCh <- err
				<-semaphore
				return
			}

			resCh <- domain.UplaodMultiPartApiRes{
				PartNumber: *u.PartNumber,
				Url:        output.URL,
			}

			<-semaphore
		}(d)

	}

	go func() {
		wg.Wait()
		close(errCh)
		close(resCh)
	}()

	for err := range errCh {

		if err != nil {

			return nil, err
		}
	}

	for resp := range resCh {

		resData = append(resData, resp)

	}
	return resData, nil

}

func (a Api) CompleteMultiPart(data domain.CompleteMultiPart) (string, error) {

	output, err := a.S3Client.CompleteMultiPart(context.Background(), data)

	if err != nil {
		return "", err
	}

	return *output.Location, nil

}
