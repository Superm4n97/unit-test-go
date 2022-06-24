package add

import "testing"

func TestAdd2(t *testing.T) {
	if Add2(5, 5) != 11 {
		t.Error("5+5 expected 10")
	}
}
