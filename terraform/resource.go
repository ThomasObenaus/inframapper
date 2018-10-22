package terraform

import "strconv"

// Resource represents a resource that was described in terraform.
type Resource interface {
	// Id returns the id of the real resource (i.e. an AWS ID like 'vpc-f8168d93')
	ID() string

	// Name returns the name of the resource that is used in terraform code (i.e. 'aws_vpc.vpc_main')
	Name() string

	// Type returns the type of this resource (i.e. aws_vpc)
	Type() ResourceType

	// Dependencies returns a list of dependant resources
	Dependencies() []string

	// Returns the name of the provider as it was defined in the terraform code.
	Provider() string

	String() string
}

type resourceImpl struct {
	id        string
	name      string
	rType     ResourceType
	dependsOn []string
	provider  string
}

func (r *resourceImpl) ID() string {
	return r.id
}

func (r *resourceImpl) Name() string {
	return r.name
}

func (r *resourceImpl) Type() ResourceType {
	return r.rType
}

func (r *resourceImpl) Dependencies() []string {
	return r.dependsOn
}

func (r *resourceImpl) Provider() string {
	return r.provider
}

func (r *resourceImpl) String() string {
	return "[" + r.Type().String() + "] id=" + r.ID() + ",n=" + r.Name() + ",#deps=" + strconv.Itoa(len(r.Dependencies()))
}
