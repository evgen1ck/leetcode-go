package solution

import (
	"reflect"
	"testing"
)

func TestReadBinaryWatch(t *testing.T) {
	type input struct {
		turnedOn int
	}
	type output struct {
		value []string
	}

	tests := []struct {
		name string
		in   input
		out  output
	}{
		{"Example 1", input{1}, output{[]string{"0:01", "0:02", "0:04", "0:08", "0:16", "0:32", "1:00", "2:00", "4:00", "8:00"}}},
		{"Example 2", input{9}, output{nil}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := output{
				value: readBinaryWatch(tc.in.turnedOn),
			}

			if !reflect.DeepEqual(got, tc.out) {
				t.Errorf("\nInput:    %+v\nExpected: %+v\nGot:      %+v", tc.in, tc.out, got)
			}
		})
	}
}
