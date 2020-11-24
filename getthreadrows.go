package tapfed

import (
	"time"

	"github.com/bmheenan/taps"
)

// GetThreadrows provides a list of threadrows for listing threads
func GetThreadrows(tq Threadquery, cb func([]taps.Threadrow, error)) {
	time.Sleep(200 * time.Millisecond) // simulate network traffic
	/*if tq.Stakeholder == "frontend@ex.com" && tq.Iteration == "2020 Q4" {*/
	cb([]taps.Threadrow{
		taps.Threadrow{
			ID:    1,
			Name:  "Big Q4 project",
			Cost:  38,
			Iter:  "2020 Q4",
			Owner: taps.Stakeholder{Email: "frontend@ex.com", Name: "Frontend"},
			Children: []taps.Threadrow{
				taps.Threadrow{
					ID:       2,
					Name:     "Design big Q4 project",
					Cost:     10,
					Iter:     "2020 Oct",
					Owner:    taps.Stakeholder{Email: "brandon@ex.com", Name: "Brandon"},
					Children: []taps.Threadrow{},
				},
				taps.Threadrow{
					ID:       3,
					Name:     "Implement big Q4 project",
					Cost:     20,
					Iter:     "2020 Oct",
					Owner:    taps.Stakeholder{Email: "steph@ex.com", Name: "Steph"},
					Children: []taps.Threadrow{},
				},
				taps.Threadrow{
					ID:       4,
					Name:     "Test big Q4 project",
					Cost:     8,
					Iter:     "2020 Nov",
					Owner:    taps.Stakeholder{Email: "brandon@ex.com", Name: "Brandon"},
					Children: []taps.Threadrow{},
				},
			},
		},
	}, nil)
	/*} else if tq.Stakeholder == "frontend@ex.com" && tq.Iteration == "Backlog" {
		cb([]Threadrow{
			Threadrow{
				ID:             5,
				Name:           "Future project",
				Cost:           20,
				Iteration:      "Backlog",
				Owner:          Personteam{Email: "frontend@ex.com", Name: "Frontend"},
				ChildrenLoaded: true,
				Children:       []Threadrow{},
			},
			Threadrow{
				ID:             6,
				Name:           "Another future project",
				Cost:           30,
				Iteration:      "Backlog",
				Owner:          Personteam{Email: "frontend@ex.com", Name: "Frontend"},
				ChildrenLoaded: true,
				Children:       []Threadrow{},
			},
		}, nil)
	} else if tq.Stakeholder == "brandon@ex.com" && tq.Iteration == "Inbox" {
		cb([]Threadrow{
			Threadrow{
				ID:             7,
				Name:           "Untriaged item",
				Cost:           8,
				Iteration:      "Inbox",
				Owner:          Personteam{Email: "brandon@ex.com", Name: "Brandon"},
				ChildrenLoaded: true,
				Children:       []Threadrow{},
			},
		}, nil)
	} else if tq.Stakeholder == "brandon@ex.com" && tq.Iteration == "2020 Oct" {
		cb([]Threadrow{
			Threadrow{
				ID:             2,
				Name:           "Design big Q4 project",
				Cost:           10,
				Iteration:      "2020 Oct",
				Owner:          Personteam{Email: "brandon@ex.com", Name: "Brandon"},
				Children:       []Threadrow{},
				ChildrenLoaded: false,
			},
			Threadrow{
				ID:             8,
				Name:           "Personal October project",
				Cost:           30,
				Iteration:      "2020 Oct",
				Owner:          Personteam{Email: "brandon@ex.com", Name: "Brandon"},
				ChildrenLoaded: true,
				Children:       []Threadrow{},
			},
			Threadrow{
				ID:             9,
				Name:           "Random todo",
				Cost:           1,
				Iteration:      "2020 Oct",
				Owner:          Personteam{Email: "brandon@ex.com", Name: "Brandon"},
				ChildrenLoaded: true,
				Children:       []Threadrow{},
			},
		}, nil)
	} else if tq.Stakeholder == "brandon@ex.com" && tq.Iteration == "2020 Nov" {
		cb([]Threadrow{
			Threadrow{
				ID:             4,
				Name:           "Test big Q4 project",
				Cost:           8,
				Iteration:      "2020 Nov",
				Owner:          Personteam{Email: "brandon@ex.com", Name: "Brandon"},
				Children:       []Threadrow{},
				ChildrenLoaded: true,
			},
		}, nil)
	} else if tq.Stakeholder == "steph@ex.com" && tq.Iteration == "Inbox" {
		cb([]Threadrow{
			Threadrow{
				ID:             10,
				Name:           "An item not yet triaged",
				Cost:           10,
				Iteration:      "Inbox",
				Owner:          Personteam{Email: "steph@ex.com", Name: "Steph"},
				Children:       []Threadrow{},
				ChildrenLoaded: true,
			},
		}, nil)
	} else if tq.Stakeholder == "steph@ex.com" && tq.Iteration == "2020 Oct" {
		cb([]Threadrow{
			Threadrow{
				ID:             3,
				Name:           "Implement big Q4 project",
				Cost:           20,
				Iteration:      "2020 Oct",
				Owner:          Personteam{Email: "steph@ex.com", Name: "Steph"},
				Children:       []Threadrow{},
				ChildrenLoaded: false,
			},
			Threadrow{
				ID:             11,
				Name:           "Cool project thing",
				Cost:           30,
				Iteration:      "2020 Oct",
				Owner:          Personteam{Email: "steph@ex.com", Name: "Steph"},
				Children:       []Threadrow{},
				ChildrenLoaded: false,
			},
			Threadrow{
				ID:             12,
				Name:           "A little task by itself",
				Cost:           3,
				Iteration:      "2020 Oct",
				Owner:          Personteam{Email: "steph@ex.com", Name: "Steph"},
				Children:       []Threadrow{},
				ChildrenLoaded: false,
			},
		}, nil)
	} else if tq.Stakeholder == "steph@ex.com" && tq.Iteration == "Backlog" {
		cb([]Threadrow{
			Threadrow{
				ID:             13,
				Name:           "Future work",
				Cost:           20,
				Iteration:      "Backlog",
				Owner:          Personteam{Email: "steph@ex.com", Name: "Steph"},
				Children:       []Threadrow{},
				ChildrenLoaded: false,
			},
			Threadrow{
				ID:             14,
				Name:           "More future work",
				Cost:           10,
				Iteration:      "Backlog",
				Owner:          Personteam{Email: "steph@ex.com", Name: "Steph"},
				Children:       []Threadrow{},
				ChildrenLoaded: false,
			},
			Threadrow{
				ID:             15,
				Name:           "A totally out there future project idea",
				Cost:           80,
				Iteration:      "Backlog",
				Owner:          Personteam{Email: "steph@ex.com", Name: "Steph"},
				Children:       []Threadrow{},
				ChildrenLoaded: false,
			},
		}, nil)
	} else if tq.ChildID == 2 || tq.ChildID == 3 || tq.ChildID == 4 {
		cb([]Threadrow{
			Threadrow{
				ID:             1,
				Name:           "Big Q4 project",
				Cost:           38,
				Iteration:      "2020 Q4",
				Owner:          Personteam{Email: "frontend@ex.com", Name: "Frontend"},
				ChildrenLoaded: true,
				Children: []Threadrow{
					Threadrow{
						ID:             2,
						Name:           "Design big Q4 project",
						Cost:           10,
						Iteration:      "2020 Oct",
						Owner:          Personteam{Email: "brandon@ex.com", Name: "Brandon"},
						Children:       []Threadrow{},
						ChildrenLoaded: false,
					},
					Threadrow{
						ID:             3,
						Name:           "Implement big Q4 project",
						Cost:           20,
						Iteration:      "2020 Oct",
						Owner:          Personteam{Email: "steph@ex.com", Name: "Steph"},
						Children:       []Threadrow{},
						ChildrenLoaded: false,
					},
					Threadrow{
						ID:             4,
						Name:           "Test big Q4 project",
						Cost:           8,
						Iteration:      "2020 Nov",
						Owner:          Personteam{Email: "brandon@ex.com", Name: "Brandon"},
						Children:       []Threadrow{},
						ChildrenLoaded: false,
					},
				},
			},
		}, nil)
	} else if tq.ParentID == 1 && tq.Iteration == "2020 Oct" {
		cb([]Threadrow{
			Threadrow{
				ID:             2,
				Name:           "Design big Q4 project",
				Cost:           10,
				Iteration:      "2020 Oct",
				Owner:          Personteam{Email: "brandon@ex.com", Name: "Brandon"},
				Children:       []Threadrow{},
				ChildrenLoaded: false,
			},
			Threadrow{
				ID:             3,
				Name:           "Implement big Q4 project",
				Cost:           20,
				Iteration:      "2020 Oct",
				Owner:          Personteam{Email: "steph@ex.com", Name: "Steph"},
				Children:       []Threadrow{},
				ChildrenLoaded: false,
			},
		}, nil)
	} else if tq.ParentID == 1 && tq.Iteration == "2020 Nov" {
		cb([]Threadrow{
			Threadrow{
				ID:             4,
				Name:           "Test big Q4 project",
				Cost:           8,
				Iteration:      "2020 Nov",
				Owner:          Personteam{Email: "brandon@ex.com", Name: "Brandon"},
				Children:       []Threadrow{},
				ChildrenLoaded: false,
			},
		}, nil)
	} else {
		cb([]Threadrow{}, nil)
	}*/
}
