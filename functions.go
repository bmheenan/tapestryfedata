package tapestryfedata

// GetPersonteam returns the info for a person or team, given their ID
func GetPersonteam(email string) Personteam {
	return Personteam{
		Email: email,
		Name:  "Brandon",
	}
}

// GetIterationsFromUserID takes a user ID and returns which iterations are relevant to display
func GetIterationsFromUserID(email string) []string {
	return []string{"Inbox", "2020 Oct", "Backlog"}
}

// GetIterationsFromParentID takes the ID of a thread and returns all iterations that contain children threads
func GetIterationsFromParentID(id int) []string {
	return []string{"2020 Oct", "2020 Q4"}
}

// GetThreads returns the threads that match the given query. Returned threads will match any one of the values in all
// of the arrays provided in the ThreadQuery.
func GetThreads(q Threadquery) []Thread {
	if len(q.IDs) == 1 {
		return []Thread{
			{
				ID:   1,
				Name: "Example thread name",
				Owner: Personteam{
					Email: "a@example.com",
					Name:  "Brandon",
				},
				Cost: 5,
			},
		}
	}
	if len(q.Stakeholders) > 0 {
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
		}
	}
	if len(q.ParentIDs) > 0 {
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
		}
	}
	if len(q.ChildIDs) > 0 {
		return []Thread{
			{
				ID:   1,
				Name: "Parent",
				Cost: 80,
			},
		}
	}
	return []Thread{}
}
