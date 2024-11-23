package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack []int

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func (s *Stack) Pop() {
	if len(*s) > 0 {
		*s = (*s)[:len(*s)-1]
	}
}

func (s *Stack) Size() {
	fmt.Println(len(*s))
}

func (s *Stack) Get(order int) {
	if len(*s) >= order {
		fmt.Println((*s)[order-1])
	}
}

type Command struct {
	Name  string
	Value int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var n int
	fmt.Scanln(&n)

	cmds := make([]Command, n)
	for i := 0; i < n; i++ {
		if !scanner.Scan() {
			break
		}

		line := scanner.Text()
		parts := strings.Fields(line)
		cmd := parts[0]

		value := -1
		if len(parts) > 1 {
			value, _ = strconv.Atoi(parts[1])
		}
		cmds[i] = Command{Name: cmd, Value: value}
	}

	stack := Stack{}
	for _, c := range cmds {
		if c.Name == "push_back" {
			stack.Push(c.Value)
		} else if c.Name == "pop_back" {
			stack.Pop()
		} else if c.Name == "size" {
			stack.Size()
		} else {
			stack.Get(c.Value)
		}
	}
}
