package solution

import (
	"reflect"
	"testing"
)

func TestCountPrimeSetBits(t *testing.T) {
	type input struct {
		left  int
		right int
	}
	type output struct {
		value int
	}

	tests := []struct {
		name string
		in   input
		out  output
	}{
		{"Example 1", input{6, 10}, output{4}},
		{"Example 2", input{10, 15}, output{5}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := output{
				value: countPrimeSetBits(tc.in.left, tc.in.right),
			}

			if !reflect.DeepEqual(got, tc.out) {
				t.Errorf("\nInput:    %+v\nExpected: %+v\nGot:      %+v", tc.in, tc.out, got)
			}
		})
	}
}
