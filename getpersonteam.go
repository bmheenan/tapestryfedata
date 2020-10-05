package tapestryfedata

import "time"

// GetPersonteam takes a person or team's email and returns all the info for them.
func GetPersonteam(email string, callback func(Personteam, error)) {
	time.Sleep(time.Second) // simulate network traffic
	if email == "brandon@ex.com" {
		callback(Personteam{
			Email: "brandon@ex.com",
			Name:  "Brandon",
		}, nil)
	} else if email == "steph@ex.com" {
		callback(Personteam{
			Email: "steph@ex.com",
			Name:  "Steph",
		}, nil)
	} else if email == "mad@ex.com" {
		callback(Personteam{
			Email: "mad@ex.com",
			Name:  "Marc-Antoine",
		}, nil)
	} else if email == "frontend@ex.com" {
		callback(Personteam{
			Email: "frontend@ex.com",
			Name:  "Frontend",
		}, nil)
	}
}
