package solution

import (
	"reflect"
	"testing"
)

func TestConcatenatedBinary(t *testing.T) {
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
		{"Example 1", input{1}, output{1}},
		{"Example 2", input{3}, output{27}},
		{"Example 3", input{12}, output{505379714}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := output{
				value: concatenatedBinary(tc.in.n),
			}

			if !reflect.DeepEqual(got, tc.out) {
				t.Errorf("\nInput:    %+v\nExpected: %+v\nGot:      %+v", tc.in, tc.out, got)
			}
		})
	}
}
