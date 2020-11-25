package tapfed

import (
	"fmt"

	"github.com/bmheenan/taps"
)

// GetThreadrowsForStkIter returns all threads (as hierarchical threadrows) for the given stakeholder `stk` and iteration
// `iter` via the channel `chThs`, or an error via the channel `chErr`
func GetThreadrowsForStkIter(stk, iter string, chThs chan<- []taps.Threadrow, chErr chan<- error) {
	b := taps.APIThreadrowsGetRes{}
	if ok := getDecodedRes(fmt.Sprintf("%s/threadrows/?stk=%s&iter=%s", baseURL, stk, iter), b, chErr); !ok {
		return
	}
	chThs <- b.Ths
}
