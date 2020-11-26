package tapfed

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/bmheenan/taps"
)

// GetStk returns the stakeholder matching the given `email`, returned in the provided `callback` function
func GetStk(email string, callback func(taps.Stakeholder, error)) {
	res, err := http.Get(baseURL + "/stakeholders/" + email)
	if err != nil {
		callback(taps.Stakeholder{}, fmt.Errorf("Could not make API call: %v", err))
		return
	}
	if res.StatusCode != http.StatusOK {
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			callback(taps.Stakeholder{}, fmt.Errorf("Error %d - Could not read response body", res.StatusCode))
			return
		}
		callback(taps.Stakeholder{}, fmt.Errorf("Error %d - %s", res.StatusCode, string(b)))
		return
	}
	b := taps.APIStksGetRes{}
	err = json.NewDecoder(res.Body).Decode(&b)
	if err != nil {
		callback(taps.Stakeholder{}, errors.New("Response was not in the expected format"))
		return
	}
	callback(b.Stk, nil)
}

// GetDomainStks returns all stakeholders in the given `domain` in their hierarchy. They're returned through the
// provided `callback` function
func GetDomainStks(domain string, callback func([]taps.StkInHier, error)) {
	res, err := http.Get(baseURL + "/stakeholders/?domain=" + domain)
	if err != nil {
		callback(nil, fmt.Errorf("Could not make API call: %v", err))
		return
	}
	if res.StatusCode != http.StatusOK {
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			callback(nil, fmt.Errorf("Error %d - Could not read response body", res.StatusCode))
			return
		}
		callback(nil, fmt.Errorf("Error %d - %s", res.StatusCode, string(b)))
		return
	}
	b := taps.APIStksDomGetRes{}
	err = json.NewDecoder(res.Body).Decode(&b)
	if err != nil {
		callback(nil, errors.New("Response was not in the expected format"))
		return
	}
	callback(b.Stks, nil)
}

// NewStk creates a new stakeholder with the given information. It returns immediately, with any errors or nil returning
// through the provided `chErr` channel.
func NewStk(email, name, abbrev, colorF, colorB string, cadence taps.Cadence, pas []string, chErr chan error) {
	go func() {
		d, err := json.Marshal(taps.APIStksPostReq{
			Email:   email,
			Name:    name,
			Abbrev:  abbrev,
			ColorF:  colorF,
			ColorB:  colorB,
			Cadence: cadence,
			Pas:     pas,
		})
		if err != nil {
			chErr <- fmt.Errorf("Could not convert request to json: %v", err)
			return
		}
		res, err := http.Post(baseURL+"/stakeholders", "application/json", bytes.NewBuffer(d))
		if err != nil {
			chErr <- fmt.Errorf("Could not make API call: %v", err)
			return
		}
		if res.StatusCode != http.StatusCreated {
			b, err := ioutil.ReadAll(res.Body)
			if err != nil {
				chErr <- fmt.Errorf("Error %d - Could not read response body", res.StatusCode)
				return
			}
			chErr <- fmt.Errorf("Error %d - %s", res.StatusCode, string(b))
			return
		}
		chErr <- nil
	}()
}
