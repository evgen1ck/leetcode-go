package solution

import (
	"reflect"
	"testing"
)

func TestNumSteps(t *testing.T) {
	type input struct {
		s string
	}
	type output struct {
		value int
	}

	tests := []struct {
		name string
		in   input
		out  output
	}{
		{"Example 1", input{"1101"}, output{6}},
		{"Example 2", input{"10"}, output{1}},
		{"Example 3", input{"1"}, output{0}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := output{
				value: numSteps(tc.in.s),
			}

			if !reflect.DeepEqual(got, tc.out) {
				t.Errorf("\nInput:    %+v\nExpected: %+v\nGot:      %+v", tc.in, tc.out, got)
			}
		})
	}
}
