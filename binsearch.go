package main

func BinSearch(array []int, q int) (int, bool) {
	i, j := 0, len(array)
	med := 0
	for i < j {
		med = array[(i+j)/2]
		if med < q {
			i = (i+j)/2 + 1
		} else if med > q {
			j = (i + j) / 2
		} else {
			return (i + j) / 2, true //query is found
		}
	}
	return j, false //least el that is greater than query
}
