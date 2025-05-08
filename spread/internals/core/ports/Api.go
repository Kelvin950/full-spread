package ports

import "github.com/kelvin950/spread/internals/core/domain"

type Api interface {
	CreateMultiPartUpload(data domain.CreateMultiPartUpload) (string, error)
	CreatePresignMultiPart(data []domain.UplaodMultiPart) ([]domain.UplaodMultiPartApiRes, error)
	CompleteMultiPart(data domain.CompleteMultiPart) (string, error)
}
