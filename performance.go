package main

import (
	"container/heap"
	"sort"
)

type Sortable struct {
	data   []int
	second []int
}

func (s Sortable) Len() int { return len(s.data) }
func (s Sortable) Less(x, y int) bool {
	if s.data[x] == s.data[y] {
		return s.second[x] > s.second[y]
	} else {
		return s.data[x] > s.data[y]
	}
}
func (s *Sortable) Swap(x, y int) {
	s.data[x], s.data[y] = s.data[y], s.data[x]
	s.second[x], s.second[y] = s.second[y], s.second[x]
}

const Deg = 1000000007

type kBest struct {
	sum     Extreme
	leaders []int
}

func (k kBest) Len() int { return len(k.leaders) }
func (k kBest) Less(x, y int) bool {
	return k.leaders[x] < k.leaders[y]
}
func (k *kBest) Swap(x, y int) {
	k.leaders[x], k.leaders[y] = k.leaders[y], k.leaders[x]
}
func (k *kBest) Push(x interface{}) {
	if y, ok := x.(int); ok == false {
		panic(ok)
	} else {
		k.sum = k.sum.Add(Extreme{0, y})
		k.leaders = append(k.leaders, y)
	}
}
func (k *kBest) Pop() interface{} {
	n := k.Len()
	res := k.leaders[n-1]
	k.sum = k.sum.Substract(Extreme{0, res})
	k.leaders = k.leaders[:n-1]
	return res
}

type Extreme struct {
	times, mod int
}

func (e Extreme) Add(e1 Extreme) Extreme {
	res := Extreme{
		mod:   e.mod + e1.mod,
		times: e.times + e1.times,
	}
	if res.mod >= Deg {
		res.mod -= Deg
		res.times++
	}
	return res
}
func (e Extreme) Substract(e1 Extreme) Extreme {
	res := Extreme{
		mod:   e.mod - e1.mod,
		times: e.times - e1.times,
	}
	if res.mod < 0 {
		res.mod += Deg
		res.times--
	}
	if res.times < 0 {
		panic("Extremes are unsigned")
	}
	return res
}
func (e Extreme) Multiply(e1 Extreme) Extreme {
	res := Extreme{
		mod:   e.mod * e1.mod,
		times: e.times*(e1.times*Deg+e1.mod) + e1.times*e.mod,
	}
	res.times += res.mod / Deg
	res.mod = res.mod % Deg
	return res
}

func (e Extreme) Less(e1 Extreme) bool {
	if e.times == e1.times {
		return e.mod < e1.mod
	} else {
		return e.times < e1.times
	}
}

func maxPerformance(n int, speed []int, efficiency []int, k int) int {
	proggers := new(Sortable)
	//proggers.index = []int{0: n}
	proggers.data = efficiency
	proggers.second = speed
	sort.Sort(proggers)
	leaders := new(kBest)
	heap.Init(leaders)
	for i := 0; i < k; i++ {
		heap.Push(leaders, 0)
	}
	max := Extreme{0, 0}
	for i, el := range proggers.data {
		heap.Push(leaders, proggers.second[i])
		if leaders.Len() > k {
			(heap.Pop(leaders))
		}
		curr := Extreme{0, el}.Multiply(leaders.sum)
		if max.Less(curr) {
			max = curr
		}
	}
	return max.mod
}

func main() {

	res := maxPerformance(100000, []int{2, 10, 3, 1, 5, 8}, []int{5, 4, 3, 9, 7, 2}, 2)
	print(res)
}
