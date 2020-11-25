package tapfed

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getDecodedRes(path string, b interface{}, chErr chan<- error) (ok bool) {
	res, err := http.Get(path)
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
	err = json.NewDecoder(res.Body).Decode(&b)
	if err != nil {
		chErr <- errors.New("Response was not in the expected format")
		return
	}
	ok = true
	return
}
