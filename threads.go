package tapfed

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/bmheenan/taps"
)

// GetThread returns the thread corresponding to the given `id`. It returns immediately, then sends the thread via
// the chTh channel, or an error via chErr
func GetThread(id int64, chTh chan<- taps.Thread, chErr chan<- error) {
	go func() {
		res, err := http.Get(fmt.Sprintf("%s/threads/%d", baseURL, id))
		if err != nil {
			chErr <- fmt.Errorf("Could not call API: %v", err)
			return
		}
		b := taps.APIThreadsGetRes{}
		err = json.NewDecoder(res.Body).Decode(&b)
		if err != nil {
			chErr <- fmt.Errorf("Could not decode response: %v", err)
			return
		}
		chTh <- b.Th
	}()
}

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

// ThreadUpdates is a struct describing which threads to update and what to update for them. Each change will be made
// to each member of `Ths`. Empty arrays and zero-values won't be modified. Within 'Ord`, only `Pa` or `Stk` can be set
type ThreadUpdates struct {
	Ths  []int64
	Name string
	Desc struct {
		New   bool
		Value string
	}
	Cost struct {
		New   bool
		Value int
	}
	AddParents []int64
	RmParents  []int64
	Iter       string
	Ord        struct {
		Pa  int64
		Stk string
		Val int64
	}
}

// UpdateThreads takes a struct `u` and makes the updates specified to the thread(s) specified. It returns an error if
// it receives one from the server, via `chErr`
func UpdateThreads(u *ThreadUpdates, chErr chan<- error) {
	go func() {
		if len(u.Ths) == 0 {
			chErr <- errors.New("There must be at least 1 thread to update")
			return
		}
		req := taps.APIThreadsPutReq{}
		req.Name = u.Name
		if u.Desc.New {
			req.Desc.New = true
			req.Desc.Value = u.Desc.Value
		}
		if u.Cost.New {
			req.Cost.New = true
			req.Cost.Value = u.Cost.Value
		}
		req.AddParents = u.AddParents
		req.RmParents = u.RmParents
		req.Iter = u.Iter
		req.Ord.Pa = u.Ord.Pa
		req.Ord.Stk = u.Ord.Stk
		req.Ord.Val = u.Ord.Val
		reqB, err := json.Marshal(req)
		if err != nil {
			chErr <- fmt.Errorf("Could not translate request to json: %v", err)
			return
		}
		chErrs := make(chan error, len(u.Ths))
		for _, th := range u.Ths {
			go func(th int64) {
				client := http.Client{}
				req, err := http.NewRequest(
					http.MethodPut,
					fmt.Sprintf("%s/threads/%d", baseURL, th),
					bytes.NewBuffer(reqB),
				)
				if err != nil {
					chErrs <- fmt.Errorf("Could not make PUT request for %d: %v", th, err)
					return
				}
				req.Header.Set("Content-Type", "application/json")
				res, err := client.Do(req)
				if err != nil {
					chErrs <- fmt.Errorf("Could not PUT %d: %v", th, err)
					return
				}
				if res.StatusCode != http.StatusOK {
					chErrs <- fmt.Errorf("API returned status %d", res.StatusCode)
					return
				}
				chErrs <- nil
			}(th)
		}
		go func() {
			errs := []error{}
			for i := 0; i < len(u.Ths); i++ {
				err = <-chErrs
				if err != nil {
					errs = append(errs, err)
				}
			}
			if len(errs) > 0 {
				s := "Errors sending PUT request:"
				for _, e := range errs {
					s = s + " [" + e.Error() + "]"
				}
				chErr <- errors.New(s)
				return
			}
			chErr <- nil
		}()
	}()
}
