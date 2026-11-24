package main

import (
	"reflect"
	"testing"
)

func TestMaxSumCombinations(t *testing.T) {
	cases := []struct {
		a, b []int
		k    int
		want []int
	}{
		{[]int{1, 4, 2, 3}, []int{2, 5, 1, 6}, 4, []int{10, 9, 9, 8}},
		{[]int{3, 2}, []int{1, 4}, 2, []int{7, 6}},
		{[]int{1}, []int{1}, 1, []int{2}},
		{[]int{1}, []int{1}, 5, []int{2}}, // k capped at grid size
		{[]int{}, []int{1}, 3, []int{}},
	}
	for _, c := range cases {
		if got := MaxSumCombinations(c.a, c.b, c.k); !reflect.DeepEqual(got, c.want) {
			t.Errorf("MaxSumCombinations(%v,%v,%d) = %v, want %v", c.a, c.b, c.k, got, c.want)
		}
	}
}
