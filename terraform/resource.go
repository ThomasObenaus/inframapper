package terraform

type Resource interface {
	Id() string
	Type() Type
	String() string
}

// Type represents the type of an aws resource
type Type int

const (
	Type_aws_vpc Type = iota
)
