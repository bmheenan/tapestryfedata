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

// GetThreadrowsForParentIter returns all threads (as hierarchical threadrows) for the given parent thread `parent`
// and iteration `iter`. It returns immediately, then sends results via the channel `chThs`, or an error via `chErr`
func GetThreadrowsForParentIter(parent int64, iter string, chThs chan<- []taps.Threadrow, chErr chan<- error) {
	go func() {
		b := taps.APIThreadrowsGetRes{}
		err := decodeGet(fmt.Sprintf("%s/threadrows/?parent=%d&iter=%s", baseURL, parent, iter), &b)
		if err != nil {
			chErr <- err
			return
		}
		chThs <- b.Ths
	}()
}

// GetThreadrowsForChild returns all threads (as flat threadrows) for the given child thread `child`. It
// returns immediately, then sends results via the channel `chThs`, or an error via `chErr`
func GetThreadrowsForChild(child int64, chThs chan<- []taps.Threadrow, chErr chan<- error) {
	go func() {
		b := taps.APIThreadrowsGetRes{}
		err := decodeGet(fmt.Sprintf("%s/threadrows/?child=%d", baseURL, child), &b)
		if err != nil {
			chErr <- err
			return
		}
		chThs <- b.Ths
	}()
}
