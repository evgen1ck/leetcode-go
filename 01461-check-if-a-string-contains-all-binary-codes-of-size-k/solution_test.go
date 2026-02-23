package solution

import (
	"reflect"
	"testing"
)

func TestHasAllCodes(t *testing.T) {
	type input struct {
		s string
		k int
	}
	type output struct {
		value bool
	}

	tests := []struct {
		name string
		in   input
		out  output
	}{
		{"Example 1", input{"00110110", 2}, output{true}},
		{"Example 2", input{"0110", 1}, output{true}},
		{"Example 3", input{"0110", 2}, output{false}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := output{
				value: hasAllCodes(tc.in.s, tc.in.k),
			}

			if !reflect.DeepEqual(got, tc.out) {
				t.Errorf("\nInput:    %+v\nExpected: %+v\nGot:      %+v", tc.in, tc.out, got)
			}
		})
	}
}
