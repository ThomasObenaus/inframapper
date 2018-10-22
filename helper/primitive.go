package helper

// NewTrue is the ideomatic way to return a bool pointer with value true in a oneliner.
func NewTrue() *bool {
	b := true
	return &b
}

// NewFalse is the ideomatic way to return a bool pointer with value false in a oneliner.
func NewFalse() *bool {
	b := false
	return &b
}
