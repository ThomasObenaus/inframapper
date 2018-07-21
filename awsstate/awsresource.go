package awsstate

type AWSResource interface {
	Id() string
	Type() string
}
