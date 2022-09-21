package mascotpckg_test

import (
	"testing"

	mascotpckg "example.com/go-demo-1/mascot-pckg"
)

func TestMascot(t *testing.T) {
	if mascotpckg.BestMascot() != "Suzi" {
		t.Fatal("Wrong mascot")
	}
}
