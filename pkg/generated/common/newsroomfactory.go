// Code generated by 'gen/eventhandlergen.go'  DO NOT EDIT.
// IT SHOULD NOT BE EDITED BY HAND AS ANY CHANGES MAY BE OVERWRITTEN
// Please reference 'gen/filterergen_template.go' for more details
// File was generated at 2019-10-31 19:47:09.019015 +0000 UTC
package common

var eventTypesNewsroomFactory = []string{
	"ContractInstantiation",
}

func EventTypesNewsroomFactory() []string {
	tmp := make([]string, len(eventTypesNewsroomFactory))
	copy(tmp, eventTypesNewsroomFactory)
	return tmp
}