package awsresource

type Infra interface {
	FindById(id string) *Resource

	FindVPC(id string) *Vpc
	Vpcs() []*Vpc
}
