package prison

import "testing"

type testCase struct {
	in   []int
	want []int
}

func TestPrisonAfterNDays(t *testing.T) {
	cases := []testCase{
		{[]int{0, 1, 0, 1, 1, 0, 0, 1}, []int{0, 0, 1, 1, 0, 0, 0, 0}},
	}

	for _, c := range cases {
		observed := prisonAfterNDays(c.in, 7)
		for i, cl := range c.want {
			if observed[i] != cl {
				msg := `
For %v,
expected %v
but observed %v`
				t.Fatalf(msg, c.in, c.want, observed)
			}
		}
	}
}

func TestPrisonAfterNDaysSlow(t *testing.T) {
	cases := []testCase{
		{[]int{0, 1, 0, 1, 1, 0, 0, 1}, []int{0, 0, 1, 1, 0, 0, 0, 0}},
	}

	for _, c := range cases {
		observed := prisonAfterNDaysSlow(c.in, 7)
		for i, cl := range c.want {
			if observed[i] != cl {
				msg := `
For %v,
expected %v
but observed %v`
				t.Fatalf(msg, c.in, c.want, observed)
			}
		}
	}
}

func BenchmarkPrisonAfterNDays(b *testing.B) {
	cells := []int{1, 1, 1, 1, 0, 0, 0, 0}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		prisonAfterNDays(cells, 799341212)
	}
}

func BenchmarkPrisonAfterNDaysSlow(b *testing.B) {
	cells := []int{1, 1, 1, 1, 0, 0, 0, 0}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		prisonAfterNDaysSlow(cells, 799341212)
	}
}
