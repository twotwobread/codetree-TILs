package main

import "fmt"

const (
	MaxCmdCnt  = 1000
	MaxMoveCnt = 100
	Offset     = MaxCmdCnt * MaxMoveCnt
)

type Command struct {
	MoveCount int
	Direction string
}

func (c Command) IsLeft() bool {
	return c.Direction == "L"
}

type Tile struct {
	LeftCnt  int
	RightCnt int
	Color    string
}

func (t *Tile) GoLeft() {
	t.LeftCnt += 1
	if t.IsGray() {
		t.Color = "G"
	} else {
		t.Color = "W"
	}
}

func (t *Tile) GoRight() {
	t.RightCnt += 1
	if t.IsGray() {
		t.Color = "G"
	} else {
		t.Color = "B"
	}
}

func (t Tile) IsGray() bool {
	return t.LeftCnt >= 2 && t.RightCnt >= 2
}

type Tiles struct {
	Now   int
	Tiles []Tile
}

func NewTiles() *Tiles {
	return &Tiles{
		Now:   Offset,
		Tiles: make([]Tile, MaxCmdCnt*MaxCmdCnt+Offset),
	}
}

func (t *Tiles) MoveTiles(c Command) {
	if c.IsLeft() {
		t.MoveLeft()
	} else {
		t.MoveRight()
	}
	for range c.MoveCount - 1 {
		if c.IsLeft() {
			t.Now -= 1
			t.MoveLeft()
		} else {
			t.Now += 1
			t.MoveRight()
		}
	}
}

func (t *Tiles) MoveLeft() {
	t.Tiles[t.Now].GoLeft()
}

func (t *Tiles) MoveRight() {
	t.Tiles[t.Now].GoRight()
}

func (td Tiles) GetColorCnt() (int, int, int) {
	w, b, g := 0, 0, 0
	for _, t := range td.Tiles {
		if t.Color == "W" {
			w += 1
		} else if t.Color == "B" {
			b += 1
		} else if t.Color == "G" {
			g += 1
		}
	}
	return w, b, g
}

func main() {
	var n int
	fmt.Scanln(&n)

	cmds := make([]Command, n)
	for i := range n {
		var m int
		var d string
		fmt.Scanf("%d %s", &m, &d)

		cmds[i] = Command{MoveCount: m, Direction: d}
	}

	tiles := Move(cmds)
	w, b, g := tiles.GetColorCnt()
	fmt.Printf("%d %d %d", w, b, g)
}

func Move(cmds []Command) *Tiles {
	tiles := NewTiles()
	for _, c := range cmds {
		tiles.MoveTiles(c)
	}
	return tiles
}
