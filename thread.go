package tapestryfedata

// Thread holds info for chunks of work
type Thread struct {
	ID           int
	Name         string
	Cost         int
	Owner        Personteam
	Stakeholders []Personteam
	Iteration    string
	ParentIDs    []int
	ChildIDs     []int
}
