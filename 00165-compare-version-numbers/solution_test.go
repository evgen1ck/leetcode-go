package solution

import (
	"reflect"
	"testing"
)

func TestCompareVersion(t *testing.T) {
	type input struct {
		version1 string
		version2 string
	}
	type output struct {
		value int
	}

	tests := []struct {
		name string
		in   input
		out  output
	}{
		{"Example 1", input{"1.2", "1.10"}, output{-1}},
		{"Example 2", input{"1.01", "1.001"}, output{0}},
		{"Example 3", input{"1.0", "1.0.0.0"}, output{0}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := output{
				value: compareVersion(tc.in.version1, tc.in.version2),
			}

			if !reflect.DeepEqual(got, tc.out) {
				t.Errorf("\nInput:    %+v\nExpected: %+v\nGot:      %+v", tc.in, tc.out, got)
			}
		})
	}
}
