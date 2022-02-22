package count_integers_with_event_digit_sum

import (
	"testing"
)

var (
	testCases = []struct {
		Args   int
		Result int
	}{
		{
			Args:   4,
			Result: 2,
		},
		{
			Args:   30,
			Result: 14,
		},
	}
)

func Test_countEven(t *testing.T) {
	const pre string = "count_integers_with_event_digit_sum"
	for _, c := range testCases {
		ans := countEven(c.Args)
		if ans != c.Result {
			t.Errorf("%s result not match, args=%v, expected=%v, ans=%v", pre, c.Args, c.Result, ans)
		}
	}
}
