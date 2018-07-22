package awsstate

// Resource represents an aws resource
type Resource interface {
	Id() string
	Type() Type
}

// Type represents the type of an aws resource
type Type int

const (
	Type_VPC Type = iota
)
