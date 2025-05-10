package s3

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/kelvin950/spread/internals/core/domain"
)

type IS3 interface {
	CreateMultiPartUpload(ctx context.Context, input domain.CreateMultiPartUpload) (*s3.CreateMultipartUploadOutput, error)
	CreatePresignMultiPart(ctx context.Context, input domain.UplaodMultiPart) (*v4.PresignedHTTPRequest, error)
	CompleteMultiPart(ctx context.Context, input domain.CompleteMultiPart) (*s3.CompleteMultipartUploadOutput, error)
}

type S3 struct {
	S3client     *s3.Client
	PresignCient *s3.PresignClient
}

func NewS3(cfg aws.Config, expiresin time.Duration) *S3 {

	
	s3Client := s3.NewFromConfig(cfg , func(o *s3.Options) {
		o.UseAccelerate= true
	})
	presignClient := s3.NewPresignClient(s3Client, func(po *s3.PresignOptions) {
		po.Expires = expiresin
	})
	return &S3{
		S3client:     s3Client,
		PresignCient: presignClient,
	}
}

func (s S3) CreateMultiPartUpload(ctx context.Context, input domain.CreateMultiPartUpload) (*s3.CreateMultipartUploadOutput, error) {

	ouput, err := s.S3client.CreateMultipartUpload(ctx, &s3.CreateMultipartUploadInput{
		Key:    input.Key,
		Bucket: input.BucketName,
	})

	return ouput, err
}

func (s S3) CreatePresignMultiPart(ctx context.Context, input domain.UplaodMultiPart) (*v4.PresignedHTTPRequest, error) {

	req, err := s.PresignCient.PresignUploadPart(ctx, &s3.UploadPartInput{
		Bucket:     input.Bucket,
		UploadId:   input.UploadId,
		Key:        input.Key,
		PartNumber: input.PartNumber,
	})

	return req, err

}

func (s S3) CompleteMultiPart(ctx context.Context, input domain.CompleteMultiPart) (*s3.CompleteMultipartUploadOutput, error) {

	parts := []types.CompletedPart{}

	for _, r := range input.MultipartUpload.Part {
		parts = append(parts, types.CompletedPart{
			ETag:       r.Etag,
			PartNumber: r.PartNumber,
		})
	}
	output, err := s.S3client.CompleteMultipartUpload(ctx, &s3.CompleteMultipartUploadInput{
		Bucket:   input.Bucket,
		Key:      input.Key,
		UploadId: input.UploadId,
		MultipartUpload: &types.CompletedMultipartUpload{
			Parts: parts,
		},
	})

	return output, err
}
