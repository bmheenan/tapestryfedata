package tapestryfedata

import "fmt"

// GetPersonteam returns the info for a person or team, given their ID
func GetPersonteam(id int) Personteam {
	return Personteam{
		ID:   id,
		Name: "Brandon",
	}
}

// GetThreads returns the threads that match the given query
func GetThreads(q Threadquery) ([]Thread, error) {
	if q.ConstrainStakeholder {
		return []Thread{
			{
				ID:   1,
				Name: "Example thread name",
				Cost: 5,
			},
			{
				ID:   2,
				Name: "Another thread name",
				Cost: 8,
			},
			{
				ID:   3,
				Name: "Big thread",
				Cost: 20,
			},
		}, nil
	}
	if q.ConstrainParent {
		return []Thread{
			{
				ID:   1,
				Name: "Child item",
				Cost: 5,
			},
			{
				ID:   2,
				Name: "Another child item",
				Cost: 5,
			},
			{
				ID:   3,
				Name: "Another child item",
				Cost: 5,
			},
		}, nil
	}
	if q.ConstrainChild {
		return []Thread{
			{
				ID:   1,
				Name: "Parent",
				Cost: 80,
			},
		}, nil
	}
	return nil, fmt.Errorf("Could not get threads: %w", ErrBadQuery)
}
