package tapfed

import (
	"fmt"

	"github.com/bmheenan/taps"
)

// GetItersForStk returns all iterations for the given stakeholder `stk` via the channel `chIters`, or an error via the
// channel `chErr`
func GetItersForStk(stk string, chIters chan<- []string, chErr chan<- error) {
	go func() {
		b := taps.APIItersGetRes{}
		err := decodeGet(fmt.Sprintf("%s/iterations/?stk=%s", baseURL, stk), &b)
		if err != nil {
			chErr <- err
			return
		}
		chIters <- b.Iters
	}()
}

// GetItersForParent returns all iterations for the given parent `pa` via the channel `chIters`, or an error via the
// channel `chErr`
func GetItersForParent(pa int64, chIters chan<- []string, chErr chan<- error) {
	go func() {
		b := taps.APIItersGetRes{}
		err := decodeGet(fmt.Sprintf("%s/iterations/?parent=%d", baseURL, pa), &b)
		if err != nil {
			chErr <- err
			return
		}
		chIters <- b.Iters
	}()
}
