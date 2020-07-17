package revstr

import "testing"

type testCase struct {
	in   string
	want string
}

func TestReverseWords(t *testing.T) {
	cases := []testCase{
		{"hello world", "world hello"},
		{"the sky is blue", "blue is sky the"},
		{"  hello world! ", "world! hello"},
		{"a good     example", "example good a"},
		{" ", ""},
	}

	for _, c := range cases {
		observed := reverseWords(c.in)
		if observed != c.want {
			msg := `
For string "%s",
expected string "%s",
but observed "%s"`
			t.Fatalf(msg, c.in, c.want, observed)
		}
	}
}
