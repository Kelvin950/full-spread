package ports

import "github.com/kelvin950/spread/internals/core/domain"

type DynamoClient interface {
	PutItem(item domain.Ec2Task) error
	GetItem(ec2Id, TaskId string) (domain.Ec2Task, error)
}
