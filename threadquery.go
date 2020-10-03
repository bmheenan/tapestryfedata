package tapestryfedata

import "errors"

// ErrBadQuery is returned for queries that are poorly formated
var ErrBadQuery = errors.New("Bad query")

// Threadquery specifies the query to use when looking up threads
//
// OwnerIDs and Iterations cannot contain more than 1 element
type Threadquery struct {
	OwnerIDs       []string
	StakeholderIDs []string
	ParentIDs      []int
	ChildIDs       []int
	Iterations     []string
}
