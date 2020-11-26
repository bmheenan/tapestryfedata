package tapfed

import (
	"fmt"

	"github.com/bmheenan/taps"
)

// GetThreadrowsForStkIter returns all threads (as hierarchical threadrows) for the given stakeholder `stk` and
// iteration `iter`. It returns immediately, then returns via the channel `chThs`, or an error via `chErr`
func GetThreadrowsForStkIter(stk, iter string, chThs chan<- []taps.Threadrow, chErr chan<- error) {
	go func() {
		b := taps.APIThreadrowsGetRes{}
		err := decodeGet(fmt.Sprintf("%s/threadrows/?stk=%s&iter=%s", baseURL, stk, iter), &b)
		if err != nil {
			chErr <- err
			return
		}
		chThs <- b.Ths
	}()
}
