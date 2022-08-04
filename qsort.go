package main

import (
	"fmt"
	"math/rand"
)

const N = 45

func main() {

	//test for random slice
	sorting := make([]int, N, N)
	for i := 0; i < N; i++ {
		sorting[i] = i
	}
	rand.Shuffle(len(sorting), func(i, j int) {
		sorting[i], sorting[j] = sorting[j], sorting[i]
	})
	fmt.Println(sorting)
	QSort(&sorting, 0, len(sorting))
	fmt.Println(sorting)
}

func partition(sorting *[]int, st, en int) int {
	pivot := rand.Intn(en-st) + st
	pivotValue := (*sorting)[pivot]
	//fmt.Println("Partition with pivot", pivotValue)
	(*sorting)[en-1], (*sorting)[pivot] = (*sorting)[pivot], (*sorting)[en-1] //move pivot to the end
	j := st - 1                                                               //number of partitioned element from start
	for i := st; i < en; i++ {
		if (*sorting)[i] <= pivotValue { //if smaller than pivot
			j += 1                                                      //increase count
			(*sorting)[i], (*sorting)[j] = (*sorting)[j], (*sorting)[i] //move to other smaller elements
		}
	}
	return j //returns position of pivot element
}

type Pair struct {
	St int
	En int
}

func QSort(sorting *[]int, st, en int) { //st included, en not included
	rand.Seed(239)
	Storage := make([]Pair, 0, en) //create query to store slices to sort
	pair := Pair{
		En: en,
		St: st,
	}
	Storage = append(Storage, pair) //add whole slice to query
	var p int
	for j := 0; j < len(Storage); j++ {
		//extract segment from storage
		st = Storage[j].St
		en = Storage[j].En
		if en-st <= 1 { // slice of length 1 or less is already sorted
			continue
		}
		p = partition(sorting, Storage[j].St, Storage[j].En)
		//add new pairs
		pair = Pair{En: en, St: p + 1}
		Storage = append(Storage, pair)
		pair = Pair{St: st, En: p}
		Storage = append(Storage, pair)
	}
}
