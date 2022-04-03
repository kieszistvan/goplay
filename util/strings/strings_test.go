package strings_test

import (
	"testing"

	"github.com/kieszistvan/goplay/util/strings"
)

func TestStrings(t *testing.T) {
	want := "Sup' pista"
	if got := strings.Sup("pista"); got != want {
		t.Errorf("expected '%s', actual '%s'", want, got)
	}
}
