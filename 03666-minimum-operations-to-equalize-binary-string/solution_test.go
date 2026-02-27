package solution

import (
	"reflect"
	"testing"
)

func TestMinOperations(t *testing.T) {
	type input struct {
		s string
		k int
	}
	type output struct {
		value int
	}

	tests := []struct {
		name string
		in   input
		out  output
	}{
		{"Example 1", input{"110", 1}, output{1}},
		{"Example 2", input{"0101", 3}, output{2}},
		{"Example 3", input{"101", 2}, output{-1}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := output{
				value: minOperations(tc.in.s, tc.in.k),
			}

			if !reflect.DeepEqual(got, tc.out) {
				t.Errorf("\nInput:    %+v\nExpected: %+v\nGot:      %+v", tc.in, tc.out, got)
			}
		})
	}
}
