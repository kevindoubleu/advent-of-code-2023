package day01

import "testing"

func Test_getCombinedDigit(t *testing.T) {
	testCases := []struct {
		desc     string
		input    string
		expected int
	}{
		{
			desc:     "happy",
			input:    "1abc2",
			expected: 12,
		}, {
			desc:     "preceding and succeeding letters",
			input:    "pqr3stu8vwx",
			expected: 38,
		}, {
			desc:     "more than 2 digits",
			input:    "a1b2c3d4e5f",
			expected: 15,
		}, {
			desc:     "1 digit",
			input:    "treb7uchet",
			expected: 77,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := getCombinedDigit(tC.input)
			if tC.expected != actual {
				t.Errorf("expected: %d actual: %d for input: %s", tC.expected, actual, tC.input)
			}
		})
	}
}
