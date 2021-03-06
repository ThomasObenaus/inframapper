package aws

// Resource represents an aws resource
type Resource interface {
	// ID returns the id of the AWS resource (i.e. 'vpc-f8168d93')
	ID() string

	// Type returns the type of this resource (i.e. aws_vpc)
	Type() ResourceType

	String() string
}
