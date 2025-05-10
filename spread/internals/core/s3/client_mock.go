package s3

import (
	"context"
	"fmt"
	"time"

	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/kelvin950/spread/internals/core/domain"
)

type S3_mock struct {
}

func NewS3_mock() *S3_mock {
	return &S3_mock{}
}

func (s S3_mock) CreateMultiPartUpload(ctx context.Context, input domain.CreateMultiPartUpload) (*s3.CreateMultipartUploadOutput, error) {

	return nil, nil
}

func (s S3_mock) CreatePresignMultiPart(ctx context.Context, input domain.UplaodMultiPart) (*v4.PresignedHTTPRequest, error) {


	select{
	case  <-ctx.Done() :
		 return  nil , ctx.Err()
	default :
	  
	

	}


	 
	time.Sleep(2 * time.Second)

	return &v4.PresignedHTTPRequest{
		URL: fmt.Sprintf("http://example.com/%s", *input.Key),
	}, nil
}

func (s S3_mock) CompleteMultiPart(ctx context.Context, input domain.CompleteMultiPart) (*s3.CompleteMultipartUploadOutput, error) {

	return nil, nil
}
