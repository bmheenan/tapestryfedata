package tapfed

// Threadquery specifies the query to use when looking up threads. Mulitple set values are treated as AND
type Threadquery struct {
	ID          int
	Owner       string
	Stakeholder string
	ParentID    int
	ChildID     int
	Iteration   string
}
