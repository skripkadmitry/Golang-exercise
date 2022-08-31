package main

import (
	"fmt"
	"strconv"
)

type Sudoku struct {
	Filled   [][]byte
	empty    int
	possible [9][9]map[byte]struct{}
}
type Backup struct {
	Last *Sudoku
	next *Backup
	len  int
}

func (b *Backup) Save(sudok *Sudoku) {
	res := new(Backup)
	BackedSudoku := new(Sudoku)
	BackedSudoku.Filled = [][]byte{[]byte("........."),
		[]byte("........."),
		[]byte("........."),
		[]byte("........."),
		[]byte("........."),
		[]byte("........."),
		[]byte("........."),
		[]byte("........."),
		[]byte("........."),
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			BackedSudoku.Filled[i][j] = sudok.Filled[i][j]
			BackedSudoku.possible[i][j] = map[byte]struct{}{}
			for key, el := range sudok.possible[i][j] {
				BackedSudoku.possible[i][j][key] = el
			}
		}
	}
	BackedSudoku.empty = sudok.empty
	res.Last = b.Last
	res.next = b.next
	b.Last = BackedSudoku
	b.next = res
	res.len = b.len
	b.len++
}
func (b *Backup) Retrieve() (*Sudoku, *Backup) {
	if b == nil {
		panic("Nothing is backuped!")
	}
	res := b.Last
	b = b.next
	if res == nil {
		panic("Nothing is backuped!")
	}
	return res, b
}

func (s *Sudoku) String() string {
	res := ""
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			res += strconv.Itoa(int(s.Filled[i][j]))
		}
		res += "\n"
	}
	return res + "\n" + strconv.Itoa(s.empty)
}

type Err string

func (e Err) Error() string {
	return string(e)
}
func (s *Sudoku) AddNumber(x, y, Val byte) error {
	//fmt.Println("Trying to add", x, y, Val)
	s.empty -= 1
	s.Filled[x][y] = Val                   //add value
	s.possible[x][y] = map[byte]struct{}{} //no possible values for cell with number
	Impossible := func(a, b byte) error {
		if a == x && b == y {
			return nil
		}
		if s.Filled[a][b] == Val {
			return Err("Broken Sudoku")
		}
		delete(s.possible[a][b], Val)
		if len(s.possible[a][b]) == 1 { //the only number remaining. 0 is technically impossible
			for key, _ := range s.possible[a][b] {
				if err := s.AddNumber(a, b, key); err != nil {
					return err
				}
			}
		}
		return nil
	}
	for i := 0; i < 9; i++ {
		if err := Impossible(byte(i), y); err != nil { //fill the row
			return err
		}
		if err := Impossible(x, byte(i)); err != nil { //fill the col
			return err
		}
		if err := Impossible(x-x%3+byte(i)%3, y-y%3+byte(i)/3); err != nil { //fil the cage
			return err
		}
	}
	return nil
}

func CreateGrid(board [][]byte) *Sudoku {
	res := new(Sudoku)
	res.empty = 81
	Filled := make([][]byte, 9)
	for k := 0; k < 9; k++ {
		Filled[k] = make([]byte, 9, 9)
		for l := 0; l < 9; l++ {
			res.possible[k][l] = map[byte]struct{}{
				1: {},
				2: {},
				3: {},
				4: {},
				5: {},
				6: {},
				7: {},
				8: {},
				9: {},
			}
		}
	}
	res.Filled = Filled
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if rune(board[i][j]) != '.' {
				V := board[i][j] - 48
				if err := res.AddNumber(byte(i), byte(j), V); err != nil {
					fmt.Println(res, "check da grid plz")
					panic("Wrong grid")
					return nil
				}
			}
		}
	}
	return res
}

func solveSudoku(board [][]byte) {
	Sudok := CreateGrid(board)
	Backups := new(Backup)
	for Sudok.empty > 0 {
		i := 0
		j := 0
		for ; i < 9; i++ {
			for ; (j < 9) && (Sudok.Filled[i][j] != 0); j++ {
			}
			if j == 9 {
				j = 0
			} else {
				break
			}
		}
		if i == 9 {
			Sudok, Backups = Backups.Retrieve()
		} else {
			for key, _ := range Sudok.possible[i][j] {
				delete(Sudok.possible[i][j], key)
				if len(Sudok.possible[i][j]) > 0 {
					Backups.Save(Sudok)
				}
				if Sudok.AddNumber(byte(i), byte(j), key) != nil {
					Sudok, Backups = Backups.Retrieve()
				}
				break
			}
		}
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			board[i][j] = Sudok.Filled[i][j] + 48
		}
	}
}

func trySudoku() {
	for j := 0; j < 1000; j++ {
		input := [][]byte{[]byte("........."),
			[]byte("........."),
			[]byte("........."),
			[]byte("........."),
			[]byte("........."),
			[]byte("........."),
			[]byte("........."),
			[]byte("........."),
			[]byte("........."),
		}
		for i := 0; i < 10; i++ {
			//fmt.Println("Lets go", i, j)
			solveSudoku(input)
		}

		fmt.Println(input)
	}
}
