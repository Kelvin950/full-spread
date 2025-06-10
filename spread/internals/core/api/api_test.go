package api

import (
	"errors"

	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/kelvin950/spread/internals/core/domain"
	"github.com/kelvin950/spread/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreatePresignMultiPart(t *testing.T) {

	is3Mock := mocks.NewIS3(t)

	is3Mock.On("CreatePresignMultiPart", mock.Anything, mock.Anything).Return(&v4.PresignedHTTPRequest{
		URL: "ddsd",
	}, nil)

	api := NewApItest(is3Mock, nil, nil, nil)

	res, err := api.CreatePresignMultiPart([]domain.UplaodMultiPart{
		{
			PartNumber: aws.Int32(1),
			UploadId:   aws.String("ds"),
			Key:        aws.String("ds"),
			Bucket:     aws.String("ds"),
		},
	})

	require.NoError(t, err)

	require.NotEqual(t, res, nil)

}

func TestCreateMultiPartUpload(t *testing.T) {

	is3 := mocks.NewIS3(t)

	is3.On("CreateMultiPartUpload", mock.Anything, mock.AnythingOfType("domain.CreateMultiPartUpload")).Return(nil, errors.New("failed"))

	api := NewApItest(is3, nil, nil, nil)

	_, err := api.CreateMultiPartUpload(domain.CreateMultiPartUpload{})

	require.Error(t, err)

}

func TestCompleteMultiPart(t *testing.T) {

	is3 := mocks.NewIS3(t)

	is3.On("CompleteMultiPart", mock.Anything, mock.AnythingOfType("domain.CompleteMultiPart")).Return(&s3.CompleteMultipartUploadOutput{
		Location: aws.String("dsds"),
	}, nil)

	api := NewApItest(is3, nil, nil, nil)

	ares, err := api.CompleteMultiPart(domain.CompleteMultiPart{})

	require.NoError(t, err)

	require.IsType(t, "", ares)
}
