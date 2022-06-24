package multipleCase

import (
	"fmt"
	"testing"
)

type SubtractionCases struct {
	a   int
	b   int
	res int
}

func TestSubtraction(t *testing.T) {
	cs := []SubtractionCases{
		{
			a:   8,
			b:   1,
			res: 7,
		},
		{
			a:   10,
			b:   9,
			res: 1,
		},
		{
			a:   5,
			b:   8,
			res: -3,
		},
	}

	for _, st := range cs {
		ans := Subtraction(st.a, st.b)
		if ans != st.res {
			t.Error(fmt.Sprintf("expected %d - %d = %d, found %d", st.a, st.b, st.res, ans))
		}
	}
}
