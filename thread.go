package tapfed

// Threadrow holds the summary info for a thread, needed to display it in a table with other threads
type Threadrow struct {
	ID             int
	Name           string
	Cost           int
	Owner          Personteam
	Iteration      string
	Children       []Threadrow // To stop from needing to traverse the tree too deep, ChildrenLoaded is true if Children
	ChildrenLoaded bool        // is populated, or false if that query still needs to be run.
}

// Threaddetail holds detailed information on a thread, needed for the drilldown view
type Threaddetail struct {
	ID           int
	Name         string
	Cost         int
	Owner        Personteam
	Stakeholders []Personteam
	Iteration    string
	// Children and parents aren't required; we can use functions that query using ID as the parent/child to get them
}
