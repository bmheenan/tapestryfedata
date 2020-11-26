package tapfed_test

import (
	"testing"

	"github.com/bmheenan/tapfed"
)

func TestGetItersForStk(t *testing.T) {
	stks, err := setupTest("1 stk")
	if err != nil {
		t.Fatalf("Could not set up test: %s", err)
	}
	chIters := make(chan []string)
	chErr := make(chan error)
	tapfed.GetItersForStk(stks["a"], chIters, chErr)
	select {
	case err = <-chErr:
		t.Fatalf("Could not get iterations: %v", err)
	case iters := <-chIters:
		if x, g := 4, len(iters); x != g {
			t.Fatalf("Expected %d iters, got %d", x, g)
		}
	}
}
