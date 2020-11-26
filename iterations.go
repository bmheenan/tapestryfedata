package tapfed

import (
	"fmt"
	"log"

	"github.com/bmheenan/taps"
)

// GetItersForStk returns all iterations for the given stakeholder `stk` via the channel `chIters`, or an error via the
// channel `chErr`
func GetItersForStk(stk string, chIters chan<- []string, chErr chan<- error) {
	b := taps.APIItersGetRes{}
	if ok := getDecodedRes(fmt.Sprintf("%s/iterations/?stk=%s", baseURL, stk), b, chErr); !ok {
		log.Println("getDecodedRes: not ok")
		return
	}
	log.Println("getDecodedRes: ok")
	chIters <- b.Iters
}
