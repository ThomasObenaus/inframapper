package mappedInfra

type Infra interface {
	NumResources() int
	Resources() map[string]MappedResource
	ResourceById(id string) MappedResource
}

type infraImpl struct {
	mappedResources map[string]MappedResource
}

func (in *infraImpl) NumResources() int {
	return len(in.mappedResources)
}

func (in *infraImpl) Resources() map[string]MappedResource {
	return in.mappedResources
}

func (in *infraImpl) ResourceById(id string) MappedResource {
	return in.mappedResources[id]
}
