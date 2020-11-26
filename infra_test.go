package tapfed_test

import (
	"fmt"

	"github.com/bmheenan/tapfed"
	"github.com/bmheenan/taps"
)

func setupTest(config string) (stks map[string]string, err error) {
	chErr := make(chan error)
	stks = map[string]string{}
	tapfed.ClearDomain("example.com", chErr)
	if err = <-chErr; err != nil {
		return
	}
	for _, stk := range stkConfigs[configs[config].stks] {
		tapfed.NewStk(stk.Email, stk.Name, stk.Abbrev, stk.ColorF, stk.ColorB, stk.Cadence, stk.Pas, chErr)
		if err = <-chErr; err != nil {
			err = fmt.Errorf("Could not create stakeholder: %v", err)
			return
		}
		stks[stk.Name] = stk.Email
	}
	// TODO set up threads
	return stks, nil
}

type cfgmap struct {
	stks string
	ths  string
}

var configs map[string]cfgmap = map[string]cfgmap{
	"1 stk": cfgmap{
		stks: "1 stk",
		ths:  "blank",
	},
}

var stkConfigs map[string]([]taps.APIStksPostReq) = map[string]([]taps.APIStksPostReq){
	"blank": []taps.APIStksPostReq{},
	"1 stk": []taps.APIStksPostReq{
		taps.APIStksPostReq{
			Email:   "a@example.com",
			Name:    "a",
			Abbrev:  "a",
			ColorF:  "#ffffff",
			ColorB:  "#000000",
			Cadence: taps.Monthly,
		},
	},
}
