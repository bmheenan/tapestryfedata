package tapfed

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bmheenan/taps"
)

// NewThread creates a new thread with the given info. It returns immediately, then returns the thread's ID via
// `chID` or an error via `chErr`
func NewThread(name, owner, iter string, cost int, pas, chs []int64, chID chan<- int64, chErr chan<- error) {
	go func() {
		reqB, err := json.Marshal(taps.APIThreadsPostReq{
			Name:  name,
			Owner: owner,
			Iter:  iter,
			Cost:  cost,
			Pas:   pas,
			Chs:   chs,
		})
		if err != nil {
			chErr <- fmt.Errorf("Could not convert request to json: %v", err)
			return
		}
		res, err := http.Post(baseURL+"/threads", "application/json", bytes.NewBuffer(reqB))
		if err != nil {
			chErr <- fmt.Errorf("Could not make API call: %v", err)
			return
		}
		if res.StatusCode != http.StatusCreated {
			chErr <- fmt.Errorf("Call returned code %d", res.StatusCode)
			return
		}
		resB := taps.APIThreadsPostRes{}
		err = json.NewDecoder(res.Body).Decode(&resB)
		if err != nil {
			chErr <- fmt.Errorf("Could not decode response: %v", err)
			return
		}
		chID <- resB.ID
	}()
}
