package main

import (
	"fmt"
	"sort"
)

type Enumerate struct {
	Index []int
	Nums  []int
}

func (e Enumerate) Len() int {
	return len(e.Nums)
}

func (e Enumerate) Less(i, j int) bool {
	return e.Nums[i] < e.Nums[j]
}

func (e Enumerate) Swap(i, j int) {
	e.Index[i], e.Index[j] = e.Index[j], e.Index[i]
	e.Nums[i], e.Nums[j] = e.Nums[j], e.Nums[i]
}

func main() {
	nums := []int{2, 7, 11, 15}
	index := make([]int, len(nums), len(nums))
	for i := 0; i < len(index); i++ {
		index[i] = i
	}
	Arr := Enumerate{index, nums}
	target := 9
	sort.Sort(Arr)
	i, j := 0, len(nums)-1
	for i <= j {
		if Arr.Nums[i]+Arr.Nums[j] > target {
			j -= 1
		} else if Arr.Nums[i]+Arr.Nums[j] < target {
			i += 1
		} else {
			fmt.Println(Arr.Index[i], Arr.Index[j])
			return
		}
	}

}
