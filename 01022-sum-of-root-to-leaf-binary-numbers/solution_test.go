package solution

import (
	"reflect"
	"testing"
)

func TestSumRootToLeaf(t *testing.T) {
	type input struct {
		treeVals []any
	}
	type output struct {
		value int
	}

	tests := []struct {
		name string
		in   input
		out  output
	}{
		{"Example 1", input{[]any{1, 0, 1, 0, 1, 0, 1}}, output{22}},
		{"Example 2", input{[]any{0}}, output{0}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			root := buildTree(tc.in.treeVals)

			got := output{
				value: sumRootToLeaf(root),
			}

			if !reflect.DeepEqual(got, tc.out) {
				t.Errorf("\nInput:    %v\nExpected: %v\nGot:      %v", tc.in.treeVals, tc.out.value, got.value)
			}
		})
	}
}

func buildTree(vals []any) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}
	root := &TreeNode{Val: vals[0].(int)}
	queue := []*TreeNode{root}
	i := 1
	for i < len(vals) {
		curr := queue[0]
		queue = queue[1:]
		if i < len(vals) && vals[i] != nil {
			curr.Left = &TreeNode{Val: vals[i].(int)}
			queue = append(queue, curr.Left)
		}
		i++

		if i < len(vals) && vals[i] != nil {
			curr.Right = &TreeNode{Val: vals[i].(int)}
			queue = append(queue, curr.Right)
		}
		i++
	}
	return root
}
