package terraform

// Type represents the type of an aws resource
type ResourceType int

const (
	Type_unkown ResourceType = iota
	Type_aws_vpc
)

var typeMap = map[string]ResourceType{
	"aws_vpc": Type_aws_vpc,
}

func StrToType(v string) ResourceType {
	return typeMap[v]
}
