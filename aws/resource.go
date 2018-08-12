package aws

// Resource represents an aws resource
type Resource interface {
	Id() string
	Type() ResourceType
	String() string
}
