package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type combo struct{ sum, i, j int }

type comboHeap []combo

func (h comboHeap) Len() int           { return len(h) }
func (h comboHeap) Less(i, j int) bool { return h[i].sum > h[j].sum }
func (h comboHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *comboHeap) Push(x any)        { *h = append(*h, x.(combo)) }
func (h *comboHeap) Pop() any {
	old := *h
	v := old[len(old)-1]
	*h = old[:len(old)-1]
	return v
}

// MaxSumCombinations returns the k largest values of a[i]+b[j], best-first.
// Sort both arrays descending; the pair grid is then sorted along both axes,
// so a best-first frontier explored with a max-heap visits exactly k cells.
func MaxSumCombinations(a, b []int, k int) []int {
	if len(a) == 0 || len(b) == 0 || k < 1 {
		return []int{}
	}
	x := append([]int{}, a...)
	y := append([]int{}, b...)
	sort.Sort(sort.Reverse(sort.IntSlice(x)))
	sort.Sort(sort.Reverse(sort.IntSlice(y)))
	if k > len(x)*len(y) {
		k = len(x) * len(y)
	}
	h := &comboHeap{{x[0] + y[0], 0, 0}}
	seen := map[[2]int]bool{{0, 0}: true}
	out := make([]int, 0, k)
	for len(out) < k {
		c := heap.Pop(h).(combo)
		out = append(out, c.sum)
		if c.i+1 < len(x) && !seen[[2]int{c.i + 1, c.j}] {
			seen[[2]int{c.i + 1, c.j}] = true
			heap.Push(h, combo{x[c.i+1] + y[c.j], c.i + 1, c.j})
		}
		if c.j+1 < len(y) && !seen[[2]int{c.i, c.j + 1}] {
			seen[[2]int{c.i, c.j + 1}] = true
			heap.Push(h, combo{x[c.i] + y[c.j+1], c.i, c.j + 1})
		}
	}
	return out
}

func main() {
	fmt.Println(MaxSumCombinations([]int{1, 4, 2, 3}, []int{2, 5, 1, 6}, 4))
}
