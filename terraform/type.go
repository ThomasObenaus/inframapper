package terraform

// Type represents the type of an aws resource
type Type int

const (
	Type_unkown Type = iota
	Type_aws_vpc
)

var typeMap = map[string]Type{
	"aws_vpc": Type_aws_vpc,
}

func StrToType(v string) Type {
	return typeMap[v]
}
