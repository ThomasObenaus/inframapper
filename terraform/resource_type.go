package terraform

// ResourceType represents the type of an aws resource
type ResourceType int

const (
	// Type_unknown represents a unknown (not implemented resource)
	Type_unknown ResourceType = iota

	// Type_aws_vpc represents a AWS VPC
	Type_aws_vpc
)

var typeMap = map[string]ResourceType{
	"aws_vpc": Type_aws_vpc,
}

// StrToType maps a given string into a ResourceType
func StrToType(v string) ResourceType {
	return typeMap[v]
}
