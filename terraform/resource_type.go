package terraform

// ResourceType represents the type of an aws resource
type ResourceType int

const (
	// TypeUnknown represents a unknown (not implemented resource)
	TypeUnknown ResourceType = iota

	// TypeAwsVpc represents a AWS VPC
	TypeAwsVpc
)

var typeMap = map[string]ResourceType{
	"aws_vpc": TypeAwsVpc,
}

// StrToType maps a given string into a ResourceType
func StrToType(v string) ResourceType {
	return typeMap[v]
}
