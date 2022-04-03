package ints_test

import (
	"testing"

	"github.com/kieszistvan/goplay/util/ints"
)

func TestMaxInt(t *testing.T) {

	t.Run("test maximum int selection", func(t *testing.T) {
		arr := []int{10, 1000, 1, 100}
		want := arr[1]
		if got := ints.MaxInt(arr...); got != want {
			t.Errorf("failed to pick max int, want %d, got %d", want, got)
		}
	})
}
