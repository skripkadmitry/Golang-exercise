package main

import "fmt"

func mainish() {
	nums := []int{3, 2, 4}
	target := 6
	//index := make([]int, len(nums), len(nums))
	dict := map[int]int{}
	for i, el := range nums {
		if res, in := dict[target-el]; in {
			fmt.Println(i, res)
			return
		} else {
			dict[el] = i
		}

	}
}
