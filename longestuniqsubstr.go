package main

func lengthOfLongestSubstring(s string) int {
	var (
		maxl, curl int
		last       = make(map[int32]int)
	)

	for i, el := range s {
		if prev, ok := last[el]; ok {
			if curl < i-prev {
				curl += 1
			} else {
				if curl > maxl {
					maxl = curl
				}
				curl = i - prev
			}

		} else {
			curl += 1
		}
		last[el] = i
	}
	if curl > maxl {
		maxl = curl
	}
	return maxl
}
