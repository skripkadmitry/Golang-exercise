package main

import (
	"fmt"
)

type Leaf struct {
	next  map[int32]*Leaf
	edges []int
	count int
	id    int
}
type Tree struct {
	root *Leaf
}

func CreateLeaf() *Leaf {
	newleaf := new(Leaf)
	newleaf.next = map[int32]*Leaf{}
	newleaf.edges = make([]int, 0, 5)
	return newleaf
}

func (tr *Tree) Add(s string, id int) {
	cur := tr.root
	for _, el := range s {
		cur.count++
		if newleaf, ok := cur.next[el]; ok {
			cur = newleaf
		} else {
			cur.edges = append(cur.edges, int(el))
			newleaf = CreateLeaf()
			cur.next[el] = newleaf
			cur = newleaf
		}
	}
	cur.count++
	cur.edges = append(cur.edges, -1)
	newleaf := CreateLeaf()
	cur.next[-1] = newleaf
	newleaf.count++
	newleaf.id = id
}

func (tr Tree) Find(s string) int {
	//fmt.Println("Search started")
	found := make([]*Leaf, 1, 5)
	found[0] = tr.root

	var (
		//ok   bool
		mask = int32("*"[0])
	)
	for _, el := range s {
		newfound := make([]*Leaf, 0, 5)
		for _, this := range found {
			if el == mask {
				for _, edge := range this.next {
					newfound = append(newfound, edge)
				}
			} else if cur, ok := this.next[el]; ok {
				newfound = append(newfound, cur)
			} else {
				continue
			}
		}
		if len(newfound) == 0 {
			return 0
		} else {
			found = newfound
		}
	}
	result := 0
	for _, this := range found {
		result += this.count
	}
	return result
	//fmt.Println("Search finished. choosing the right option")
	/*for count := 0; count < rank; {
		if cur.count+count < rank {
			return -1
		}
		sort.Ints(cur.edges)
		for _, el := range cur.edges {
			if count+cur.next[int32(el)].count < rank {
				count += cur.next[int32(el)].count
			} else {
				if el == -1 {
					count++
				}
				cur = cur.next[int32(el)]
				break
			}
		}
	}
	//fmt.Println(cur)
	return cur.id*/
}

func Intertwine(s1, s2 string) string {
	var (
		i   = 0
		j   = len(s2) - 1
		res string
	)
	for ; (i < len(s1)) || (j >= 0); i++ {
		if i >= len(s1) {
			res += "*" + string(s2[j])
		} else if j < 0 {
			res += string(s1[i]) + "*"
		} else {
			res += string(s1[i]) + string(s2[j])
		}
		j--
	}
	//fmt.Println(res, "intertwined")
	return res
}
func Intercalate(s string) string {
	var (
		i   = 0
		j   = len(s) - 1
		res string
	)
	for ; i < len(s); i++ {
		res += string(s[i]) + string(s[j])
		j--
	}
	return res
}

func main() {
	var (
		n, m    int
		S, P, Q string
	)
	fmt.Scan(&n, &m)
	tree := new(Tree)
	tree.root = CreateLeaf()
	for i := 0; i < n; i++ {
		fmt.Scan(&S)
		tree.Add(Intercalate(S), i)
	}

	for j := 0; j < m; j++ {
		fmt.Scan(&P, &Q)
		fmt.Println(tree.Find(Intertwine(P, Q)))
	}

}

func __main__() {
	var (
		n, m int
		s    string
	)
	tree := Tree{
		root: CreateLeaf(),
	}

	fmt.Scan(&n, &m)
	for i := 0; i < n; i++ {
		fmt.Scan(&s)
		tree.Add(s, i+1)
		//fmt.Printf("%s added successfully!\n", s)
	}
	for j := 0; j < m; j++ {
		fmt.Scan(&n, &s)
		fmt.Println(tree.Find(s))
	}

}
