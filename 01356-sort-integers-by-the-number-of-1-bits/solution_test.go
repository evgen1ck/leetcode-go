package solution

import (
	"reflect"
	"testing"
)

func TestSortByBits(t *testing.T) {
	type input struct {
		arr []int
	}
	type output struct {
		value []int
	}

	tests := []struct {
		name string
		in   input
		out  output
	}{
		{"Example 1", input{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8}}, output{[]int{0, 1, 2, 4, 8, 3, 5, 6, 7}}},
		{"Example 2", input{[]int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1}}, output{[]int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := output{
				value: sortByBits(tc.in.arr),
			}

			if !reflect.DeepEqual(got, tc.out) {
				t.Errorf("\nInput:    %+v\nExpected: %+v\nGot:      %+v", tc.in, tc.out, got)
			}
		})
	}
}
