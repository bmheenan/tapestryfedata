package tapfed_test

import (
	"testing"

	"github.com/bmheenan/tapfed"
)

func TestUpdateIter(t *testing.T) {
	_, err := setupTest("1 stk")
	if err != nil {
		t.Fatalf("Could not set up test: %s", err)
	}
	chID := make(chan int64)
	chErr := make(chan error)
	tapfed.NewThread("A", "a@example.com", "2020-11 Nov", 10, []int64{}, []int64{}, chID, chErr)
	var id int64
	select {
	case id = <-chID:
	case err = <-chErr:
		t.Fatalf("Could not make new thread: %v", err)
	}
	tapfed.UpdateThreads(&tapfed.ThreadUpdates{
		Ths:  []int64{id},
		Iter: "2020-12 Dec",
	}, chErr)
	err = <-chErr
	if err != nil {
		t.Fatalf("Could not update iteration: %v", err)
	}
}
