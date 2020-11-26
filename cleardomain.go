package tapfed

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/bmheenan/taps"
)

// ClearDomain deletes all information within the given `domain`. It returns immediately then executes the
// API call concurrently, returning an error or nil over `chErr` when done.
func ClearDomain(domain string, chErr chan<- error) {
	go func() {
		d, err := json.Marshal(taps.APICleardomPostReq{
			Domain: domain,
		})
		if err != nil {
			chErr <- fmt.Errorf("Could not convert domain to json: %v", err)
			return
		}
		res, err := http.Post(baseURL+"/cleardomain", "application/json", bytes.NewBuffer(d))
		if err != nil {
			chErr <- fmt.Errorf("Could not make API call: %v", err)
			return
		}
		if res.StatusCode != http.StatusOK {
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
