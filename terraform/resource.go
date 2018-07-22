package terraform

type Resource struct {
	Id        string
	Type      Type
	DependsOn []string
}

func (r *Resource) String() string {
	return "id=" + r.Id + ", type=" + r.Type.String()
}

// Type represents the type of an aws resource
type Type int

const (
	Type_aws_vpc Type = iota
)
