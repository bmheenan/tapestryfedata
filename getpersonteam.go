package tapfed

import (
	"time"

	"github.com/bmheenan/taps"
)

// GetPersonteam takes a person or team's email and returns all the info for them.
func GetPersonteam(email string, cb func(taps.Stakeholder, error)) {
	time.Sleep(300 * time.Millisecond) // simulate network traffic
	if email == "brandon@ex.com" {
		cb(taps.Stakeholder{
			Email: "brandon@ex.com",
			Name:  "Brandon",
		}, nil)
	} else if email == "steph@ex.com" {
		cb(taps.Stakeholder{
			Email: "steph@ex.com",
			Name:  "Steph",
		}, nil)
	} else if email == "frontend@ex.com" {
		cb(taps.Stakeholder{
			Email: "frontend@ex.com",
			Name:  "Frontend",
		}, nil)
	}
}
