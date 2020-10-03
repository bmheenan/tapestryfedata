package tapestryfedata

// Threadquery specifies the query to use when looking up threads
type Threadquery struct {
	IDs          []int
	Owners       []string
	Stakeholders []string
	ParentIDs    []int
	ChildIDs     []int
	Iterations   []string
}
