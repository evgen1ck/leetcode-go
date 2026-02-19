package solution

import (
	"reflect"
	"testing"
)

func TestHasAlternatingBits(t *testing.T) {
	type input struct {
		n int
	}
	type output struct {
		value bool
	}

	tests := []struct {
		name string
		in   input
		out  output
	}{
		{"Example 1", input{5}, output{true}},
		{"Example 2", input{7}, output{false}},
		{"Example 3", input{1}, output{false}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := output{
				value: hasAlternatingBits(tc.in.n),
			}

			if !reflect.DeepEqual(got, tc.out) {
				t.Errorf("\nInput:    %+v\nExpected: %+v\nGot:      %+v", tc.in, tc.out, got)
			}
		})
	}
}
