package solution

import (
	"reflect"
	"testing"
)

func TestAddBinary(t *testing.T) {
	type input struct {
		a string
		b string
	}
	type output struct {
		value string
	}

	tests := []struct {
		name string
		in   input
		out  output
	}{
		{"Example 1", input{"11", "1"}, output{"100"}},
		{"Example 2", input{"1010", "1011"}, output{"10101"}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := output{
				value: addBinary(tc.in.a, tc.in.b),
			}

			if !reflect.DeepEqual(got, tc.out) {
				t.Errorf("\nInput:    %+v\nExpected: %+v\nGot:      %+v", tc.in, tc.out, got)
			}
		})
	}
}
