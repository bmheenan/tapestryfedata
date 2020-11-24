package tapfed

// Threadquery specifies the query to use when looking up threads. Mulitple set values are treated as AND
type Threadquery struct {
	ID          int64
	Owner       string
	Stakeholder string
	ParentID    int64
	ChildID     int64
	Iteration   string
}
