package add

import "testing"

func TestAdd(t *testing.T) {
	if Add(2, 3) != 5 {
		t.Error("expected 2+3 = 5")
	}
}
