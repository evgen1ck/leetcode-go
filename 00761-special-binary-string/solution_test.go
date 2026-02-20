package solution

import (
	"reflect"
	"testing"
)

func TestMakeLargestSpecial(t *testing.T) {
	type input struct {
		s string
	}
	type output struct {
		value string
	}

	tests := []struct {
		name string
		in   input
		out  output
	}{
		{"Example 1", input{"11011000"}, output{"11100100"}},
		{"Example 2", input{"10"}, output{"10"}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := output{
				value: makeLargestSpecial(tc.in.s),
			}

			if !reflect.DeepEqual(got, tc.out) {
				t.Errorf("\nInput:    %+v\nExpected: %+v\nGot:      %+v", tc.in, tc.out, got)
			}
		})
	}
}
