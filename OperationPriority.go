package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type State struct {
	id    int
	Links map[string]*State
}

type StackEl struct {
	Val  string
	Next *StackEl
}
type Stack struct {
	Visible *StackEl
}

func (s *Stack) Add(el string) {
	var next = StackEl{el, s.Visible}
	s.Visible = &next
}
func (s *Stack) Pop() string {
	if s.Visible == nil {
		return ""
	}
	res := s.Visible.Val
	s.Visible = s.Visible.Next
	return res
}
func (s *Stack) String() string {
	res := ""
	for el := s.Visible; el != nil; el = el.Next {
		res += el.Val + " "
	}
	return res
}

func OperationPrecedence(op string) int {
	switch op {
	case "*":
		return 2
	case "+":
		return 1
	case "-":
		return 1
	case "/":
		return 2
	case "^":
		return 3 // this is more complicated
	case "(":
		return -1 // less than any other operation
	default:
		return 0
	}
}

type WrongOperation string

func (w WrongOperation) Error() string {
	return fmt.Sprintf("Unknown operation: %s", w)
}

func Operation(a1, b1, op string) (int, error) {
	var (
		a, _ = strconv.Atoi(a1)
		b, _ = strconv.Atoi(b1)
	)
	switch op {
	case "*":
		return a * b, nil
	case "+":
		return a + b, nil
	case "-":
		return b - a, nil
	case "/":
		return b / a, nil
	case "^":
		return int(math.Pow(float64(b), float64(a))), nil // they are extracted from stack in the wrong order
	case ";":
		return a, nil
	default:
		var wo = WrongOperation(op)
		return 0, wo
	}
}
func OperatorPrecedenceParse(expression string) int { //3 * (3 - 1) ^ 5 + 6 * 12 ;
	var (
		Res      = new(Stack)
		Oper     = new(Stack)
		q        string
		new, res int
	)
	expr := strings.Split(expression, " ")
	for _, el := range expr {
		if _, err := strconv.Atoi(el); err != nil {
			if el == ")" {
				for q := Oper.Pop(); q != "("; q = Oper.Pop() {
					new, _ = Operation(Res.Pop(), Res.Pop(), q)
					Res.Add(strconv.Itoa(new))
				}
			} else if el == "(" {
				Oper.Add("(")
			} else {

				q = Oper.Pop()
				for ; q != "" && (OperationPrecedence(q) > OperationPrecedence(el)); q = Oper.Pop() {
					new, _ = Operation(Res.Pop(), Res.Pop(), q)
					Res.Add(strconv.Itoa(new))
				}
				if q != "" {
					Oper.Add(q)
				}
				Oper.Add(el)
			}
		} else {
			Res.Add(el)
		}
		//fmt.Println("Oper:", Oper, "Res:", Res)
	}
	res, _ = strconv.Atoi(Res.Pop())
	return res
}

func main() {
	fmt.Println(OperatorPrecedenceParse("4 ^ 3 / 2 ;"))
}
