// Code generated by "stringer -type=ResourceType mappedInfra/resource_type.go"; DO NOT EDIT.

package mappedInfra

import "strconv"

const _ResourceType_name = "TypeVPC"

var _ResourceType_index = [...]uint8{0, 7}

func (i ResourceType) String() string {
	if i < 0 || i >= ResourceType(len(_ResourceType_index)-1) {
		return "ResourceType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ResourceType_name[_ResourceType_index[i]:_ResourceType_index[i+1]]
}
