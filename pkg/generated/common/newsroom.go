// Code generated by 'gen/eventhandlergen.go'  DO NOT EDIT.
// IT SHOULD NOT BE EDITED BY HAND AS ANY CHANGES MAY BE OVERWRITTEN
// Please reference 'gen/filterergen_template.go' for more details
// File was generated at 2018-12-19 19:02:58.654898 +0000 UTC
package common

var eventTypesNewsroomContract = []string{
	"ContentPublished",
	"NameChanged",
	"OwnershipRenounced",
	"OwnershipTransferred",
	"RevisionSigned",
	"RevisionUpdated",
	"RoleAdded",
	"RoleRemoved",
}

func EventTypesNewsroomContract() []string {
	tmp := make([]string, len(eventTypesNewsroomContract))
	copy(tmp, eventTypesNewsroomContract)
	return tmp
}
