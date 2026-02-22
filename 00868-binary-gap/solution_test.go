package solution

import (
	"reflect"
	"testing"
)

func TestBinaryGap(t *testing.T) {
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
		{"Example 1", input{22}, output{2}},
		{"Example 2", input{8}, output{0}},
		{"Example 3", input{5}, output{2}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := output{
				value: binaryGap(tc.in.n),
			}

			if !reflect.DeepEqual(got, tc.out) {
				t.Errorf("\nInput:    %+v\nExpected: %+v\nGot:      %+v", tc.in, tc.out, got)
			}
		})
	}
}
