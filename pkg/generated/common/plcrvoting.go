// Code generated by 'gen/eventhandlergen.go'  DO NOT EDIT.
// IT SHOULD NOT BE EDITED BY HAND AS ANY CHANGES MAY BE OVERWRITTEN
// Please reference 'gen/filterergen_template.go' for more details
// File was generated at 2018-10-11 16:19:15.178676999 +0000 UTC
package common

var eventTypesPLCRVotingContract = []string{
	"PollCreated",
	"TokensRescued",
	"VoteCommitted",
	"VoteRevealed",
	"VotingRightsGranted",
	"VotingRightsWithdrawn",
}

func EventTypesPLCRVotingContract() []string {
	tmp := make([]string, len(eventTypesPLCRVotingContract))
	copy(tmp, eventTypesPLCRVotingContract)
	return tmp
}