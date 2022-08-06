package main

import (
	"fmt"
)

func findMedianSortedArrays0(nums1 []int, nums2 []int) float64 {
	var (
		n        = len(nums1)
		m        = len(nums2)
		i, found int
		j        = n
		cur      = (i + j) / 2
	)
	//if median is mean of two then the largest is found
	for i < j { //This is kinda binsearch of median candidate
		found, _ = BinSearch(nums2, nums1[cur])
		//fmt.Println(i, j, "borders", found)
		if cur+found > (n+m)/2 {
			j = cur
		} else if cur+found < (n+m)/2 {
			i = cur + 1
		} else { //largest of two in median(or the median itself) in nums1
			if (n+m)%2 == 1 {
				return float64(nums1[cur])
			} else if (cur == 0) || ((found > 0) && (nums1[cur-1] < nums2[found-1])) { //find the twin median
				return (float64(nums1[cur]) + float64(nums2[found-1])) / 2
			} else {
				return (float64(nums1[cur]) + float64(nums1[cur-1])) / 2
			}
		}
		cur = (i + j) / 2
	}
	//median is to find in nums2
	f := (n+m)/2 - i //here`s our candidate
	if (n+m)%2 == 1 {
		return float64(nums2[f])
	} else if (f == 0) || ((i > 0) && (nums1[i-1] > nums2[f-1])) { //twin is in nums1
		return (float64(nums2[f]) + float64(nums1[i-1])) / 2
	} else {
		return (float64(nums2[f]) + float64(nums2[f-1])) / 2
	}

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var (
		n             = len(nums1)
		m             = len(nums2)
		M             = (n + m) / 2
		odd           = (n+m)%2 == 1
		cur1, cur2, i int
		j             = n
	)
	if n > m {
		return findMedianSortedArrays(nums2, nums1)
	}
	for i < j {
		cur1 = (i + j) / 2  //take first cur1 els in nums1
		cur2 = M - cur1 - 1 //therefore 0-cur2 els are expected in nums2
		fmt.Println(cur1, cur2)
		if cur2 < 0 { //should never happen as n < m
			panic(cur2)
		} else if cur2 >= m {
			panic(cur2)
		} else if nums1[cur1] < nums2[cur2] {
			if (cur1 == n-1) || (nums1[cur1+1] >= nums2[cur2]) { //median is between nums1[cur1] and nums1[cur1 + 1]
				if odd {
					return float64(nums2[cur2]) //all this staff was tested using paper & pencil
				} else if (cur2 == 0) || (nums1[cur1] > nums2[cur2-1]) {
					return (float64(nums2[cur2] + nums1[cur1])) / 2
				} else {
					return (float64(nums2[cur2] + nums2[cur2-1])) / 2
				}
			} else if (cur2 == 0) || (nums2[cur2-1] <= nums1[cur1]) { //median is instead between num2[cur2 - 1] and num2[cur2]
				if odd {
					return float64(nums1[cur1+1])
				} else {
					return float64(nums1[cur1]+nums1[cur1+1]) / 2
				}
			} else {
				i = cur1
			}
		} else if nums1[cur1] > nums2[cur2] { //this case is analogous
			if (cur1 == 0) || (nums1[cur1-1] <= nums2[cur2]) {
				if (cur2 == m-1) || (nums2[cur2+1] > nums1[cur1]) {
					if odd {
						return float64(nums1[cur1])
					} else {
						return (float64(nums2[cur2]) + float64(nums1[cur1])) / 2
					}
				} else {
					if odd {
						return float64(nums2[cur2+1])
					} else {
						return (float64(nums2[cur2]) + float64(nums2[cur2+1])) / 2
					}
				}
			} else if (cur2 == m-1) || (nums2[cur2+1] >= nums1[cur1]) {
				if odd {
					return float64(nums1[cur1])
				} else {
					return (float64(nums1[cur1]) + float64(nums1[cur1-1])) / 2
				}
			} else {
				j = cur1
			}
		} else {
			return float64(nums1[cur1])
		}

	}
	// if nums1 is empty
	if odd {
		return float64(nums2[m/2])
	} else {
		return float64(nums2[m/2]+nums2[m/2-1]) / 2
	}
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3, 4}, []int{2, 5, 6, 7}))

}
