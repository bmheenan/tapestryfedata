package tapfed

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func decodeGet(path string, b interface{}) (err error) {
	res, err := http.Get(path)
	if err != nil {
		return fmt.Errorf("Could not make API call: %v", err)
	}
	if res.StatusCode != http.StatusOK {
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("Error %d - Could not read response body", res.StatusCode)
		}
		return fmt.Errorf("Error %d - %s", res.StatusCode, string(b))
	}
	err = json.NewDecoder(res.Body).Decode(b)
	if err != nil {
		return errors.New("Response was not in the expected format")
	}
	return nil
}
