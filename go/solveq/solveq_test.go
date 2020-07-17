package solveq

import "testing"

type testCase struct {
	in   string
	want string
}

func TestSolveEquation(t *testing.T) {
	cases := []testCase{
		{in: "2x+1x-5x+3-6=7x+5", want: "No solution"},
		{in: "2x=2", want: "x=1"},
		{in: "x=x", want: "Infinite solutions"},
		{in: "2x+3x-6x=x+2", want: "x=-1"},
		{in: "12x-2+33x+6-50x=17x-16-12x", want: "x=2"},
		{in: "100x+13=98x-3", want: "x=-8"},
	}
	for _, c := range cases {
		observed := solveEquation(c.in)
		if observed != c.want {
			msg := `
For %s,
Expected: %s
Observed: %s`
			t.Fatalf(msg, c.in, c.want, observed)
		}
	}
}
