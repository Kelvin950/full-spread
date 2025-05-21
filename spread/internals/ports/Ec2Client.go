package ports

type Ec2Client interface {
	CreateInstance() error
	DestroyInstance() error
}
