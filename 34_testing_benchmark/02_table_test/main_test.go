package main

import "testing"

func TestMySum(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}
	tests := []test{
		test{[]int{2, 4}, 6},
		test{[]int{8, 10}, 18},
		test{[]int{21, 21}, 42},
		test{[]int{1, -2}, -1},
	}
	for _, v := range tests {
		x := mySum(v.data...)
		if x != v.answer {
			t.Error("Expected", v.answer, " but got", x)
		}
	}
}
