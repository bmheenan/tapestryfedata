package tapestryfedata

import (
	"errors"
	"time"
)

// GetIterationsForPersonteam returns all iterations that contain threads owned by the given Personteam's email
func GetIterationsForPersonteam(email string, cb func([]string, error)) {
	time.Sleep(300 * time.Millisecond) // simulate network traffic
	if email == "brandon@ex.com" {
		cb([]string{"Inbox", "2020 Oct", "2020 Nov"}, nil)
	} else if email == "steph@ex.com" {
		cb([]string{"Inbox", "2020 Oct", "Backlog"}, nil)
	} else if email == "frontend@ex.com" {
		cb([]string{"2020 Q4", "Backlog"}, nil)
	} else {
		cb([]string{}, errors.New("Person not found"))
	}
}

// GetIterationsForParentThread returns all iterations that contain child threads of the given parent id
func GetIterationsForParentThread(id int, cb func([]string, error)) {
	time.Sleep(300 * time.Millisecond) // simulate network traffic
	if id == 1 {
		cb([]string{"2020 Oct", "2020 Nov"}, nil)
	} else {
		cb([]string{}, nil)
	}
}
