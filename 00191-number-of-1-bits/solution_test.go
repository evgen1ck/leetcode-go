package solution

import (
	"reflect"
	"testing"
)

func TestHammingWeight(t *testing.T) {
	type input struct {
		n int
	}
	type output struct {
		value int
	}

	tests := []struct {
		name string
		in   input
		out  output
	}{
		{"Example 1", input{11}, output{3}},
		{"Example 2", input{128}, output{1}},
		{"Example 3", input{2147483645}, output{30}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := output{
				value: hammingWeight(tc.in.n),
			}

			if !reflect.DeepEqual(got, tc.out) {
				t.Errorf("\nInput:    %+v\nExpected: %+v\nGot:      %+v", tc.in, tc.out, got)
			}
		})
	}
}
