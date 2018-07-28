package mappedInfra

type Infra interface {
	NumResources() int
	Resources() []MappedResource
	ResourcesByType() []MappedResource
	ResourceByName() MappedResource
	ResourceById() MappedResource
}

type infraImpl struct {
	mappedResources []MappedResource
}

func (in *infraImpl) NumResources() int {
	return len(in.mappedResources)
}

func (in *infraImpl) Resources() []MappedResource {
	return nil
}

func (in *infraImpl) ResourcesByType() []MappedResource {
	return nil
}

func (in *infraImpl) ResourceByName() MappedResource {
	return nil
}

func (in *infraImpl) ResourceById() MappedResource {
	return nil
}
