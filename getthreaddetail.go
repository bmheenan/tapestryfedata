package tapfed

import (
	"time"

	"github.com/bmheenan/taps"
)

// GetThreaddetail returns the details for the given thread
func GetThreaddetail(id int, cb func(taps.Thread, error)) {
	time.Sleep(300 * time.Millisecond) // simulate network traffic
	if id == 1 {
		cb(taps.Thread{
			ID:      1,
			Name:    "Big Q4 project",
			CostDir: 38,
			Owner:   taps.Stakeholder{Email: "frontend@ex.com", Name: "Frontend"},
			Iter:    "2020 Q4",
		}, nil)
	} else /*if id == 2*/ {
		cb(taps.Thread{
			ID:      2,
			Name:    "Design big Q4 project",
			CostDir: 10,
			Owner:   taps.Stakeholder{Email: "brandon@ex.com", Name: "Brandon"},
			Iter:    "2020 Oct",
		}, nil)
		/*} else if id == 3 {
			cb(Threaddetail{
				ID:        3,
				Name:      "Implement big Q4 project",
				Cost:      20,
				Owner:     Personteam{Email: "steph@ex.com", Name: "Steph"},
				Iteration: "2020 Oct",
			}, nil)
		} else if id == 4 {
			cb(Threaddetail{
				ID:        4,
				Name:      "Test big Q4 project",
				Cost:      8,
				Owner:     Personteam{Email: "brandon@ex.com", Name: "Brandon"},
				Iteration: "2020 Nov",
			}, nil)
		} else if id == 5 {
			cb(Threaddetail{
				ID:        5,
				Name:      "Future project",
				Cost:      20,
				Owner:     Personteam{Email: "frontend@ex.com", Name: "Frontend"},
				Iteration: "Backlog",
			}, nil)
		} else if id == 6 {
			cb(Threaddetail{
				ID:        6,
				Name:      "Another future project",
				Cost:      30,
				Owner:     Personteam{Email: "frontend@ex.com", Name: "Frontend"},
				Iteration: "Backlog",
			}, nil)
		} else if id == 7 {
			cb(Threaddetail{
				ID:        7,
				Name:      "Untriaged item",
				Cost:      8,
				Owner:     Personteam{Email: "brandon@ex.com", Name: "Brandon"},
				Iteration: "Inbox",
			}, nil)
		} else if id == 8 {
			cb(Threaddetail{
				ID:        8,
				Name:      "Personal October project",
				Cost:      30,
				Owner:     Personteam{Email: "brandon@ex.com", Name: "Brandon"},
				Iteration: "2020 Oct",
			}, nil)
		} else if id == 9 {
			cb(Threaddetail{
				ID:        9,
				Name:      "Random todo",
				Cost:      1,
				Owner:     Personteam{Email: "brandon@ex.com", Name: "Brandon"},
				Iteration: "2020 Oct",
			}, nil)
		} else if id == 10 {
			cb(Threaddetail{
				ID:        10,
				Name:      "An item not yet triaged",
				Cost:      10,
				Owner:     Personteam{Email: "steph@ex.com", Name: "Steph"},
				Iteration: "Inbox",
			}, nil)
		} else if id == 11 {
			cb(Threaddetail{
				ID:        11,
				Name:      "Cool project thing",
				Cost:      30,
				Owner:     Personteam{Email: "steph@ex.com", Name: "Steph"},
				Iteration: "2020 Oct",
			}, nil)
		} else if id == 12 {
			cb(Threaddetail{
				ID:        12,
				Name:      "A little task by itself",
				Cost:      3,
				Owner:     Personteam{Email: "steph@ex.com", Name: "Steph"},
				Iteration: "2020 Oct",
			}, nil)
		} else if id == 13 {
			cb(Threaddetail{
				ID:        13,
				Name:      "Future work",
				Cost:      20,
				Owner:     Personteam{Email: "steph@ex.com", Name: "Steph"},
				Iteration: "Backlog",
			}, nil)
		} else if id == 14 {
			cb(Threaddetail{
				ID:        14,
				Name:      "More future work",
				Cost:      10,
				Owner:     Personteam{Email: "steph@ex.com", Name: "Steph"},
				Iteration: "Backlog",
			}, nil)
		} else if id == 15 {
			cb(Threaddetail{
				ID:        15,
				Name:      "A totally out there future project idea",
				Cost:      80,
				Owner:     Personteam{Email: "steph@ex.com", Name: "Steph"},
				Iteration: "Backlog",
			}, nil)*/
	}
}
