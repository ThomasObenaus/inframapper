package terraform

import "strconv"

type Resource interface {
	Id() string
	Name() string
	Type() Type
	Dependencies() []string
	Provider() string
	String() string
}
type resourceImpl struct {
	id        string
	name      string
	rType     Type
	dependsOn []string
	provider  string
}

func (r *resourceImpl) Id() string {
	return r.id
}

func (r *resourceImpl) Name() string {
	return r.name
}

func (r *resourceImpl) Type() Type {
	return r.rType
}

func (r *resourceImpl) Dependencies() []string {
	return r.dependsOn
}

func (r *resourceImpl) Provider() string {
	return r.provider
}

func (r *resourceImpl) String() string {
	return "[" + r.Type().String() + "] id=" + r.Id() + ",n=" + r.Name() + ",#deps=" + strconv.Itoa(len(r.Dependencies()))
}
