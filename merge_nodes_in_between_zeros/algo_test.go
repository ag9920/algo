package merge_nodes_in_between_zeros

import (
	"reflect"
	"testing"
)

var (
	testCases = []struct {
		Args   []int
		Result []int
	}{
		{
			Args:   []int{0, 3, 1, 0, 4, 5, 2, 0},
			Result: []int{4, 11},
		},
		{
			Args:   []int{0, 1, 0, 3, 0, 2, 2, 0},
			Result: []int{1, 3, 4},
		},
	}
)

func constructListNodes(args []int) *ListNode {
	var top *ListNode
	if len(args) == 0 {
		return top
	}
	length := len(args)
	front := &ListNode{}
	top = front
	for index, arg := range args {
		front.Val = arg
		if index == length-1 {
			break
		}
		front.Next = &ListNode{}
		front = front.Next
	}
	return top
}

func unpackListNodes(head *ListNode) []int {
	var array []int
	for head != nil {
		array = append(array, head.Val)
		head = head.Next
	}
	return array
}

func Test_mergeNodes(t *testing.T) {
	const pre string = "merge_nodes_in_between_zeros"
	for _, c := range testCases {
		node := constructListNodes(c.Args)
		result := mergeNodes(node)
		ans := unpackListNodes(result)
		if !reflect.DeepEqual(result, c.Result) {
			t.Errorf("%s result not match, args=%v, expected=%v, ans=%v", pre, c.Args, c.Result, ans)
		}
	}
}
