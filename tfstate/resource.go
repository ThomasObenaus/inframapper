package tfstate

type Resource interface {
	Id() string
	Type() string
	String() string
}

// FIXME make use of this resource type instead of plain tfstate
