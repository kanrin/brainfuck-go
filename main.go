package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

var debug bool
var stackSize int

func init() {
	flag.BoolVar(&debug, "d", false, "open debug mode")
	flag.IntVar(&stackSize, "s", 128, "set run stack size")
	flag.Parse()
}

type Stack struct {
	i []int
}

func (q *Stack) Push(i int) {
	q.i = append(q.i, i)
}

func (q *Stack) Pop() int {
	if len(q.i) == 0 {
		panic("no stack item")
	}

	index := len(q.i) - 1
	item := q.i[index]
	q.i = q.i[:index]

	return item
}

func brainfuck(in []byte) []uint8 {
	memStack := make([]uint8, stackSize)
	loopStack := new(Stack)
	ptr := 0
	for i := 0; i < len(in); i++ {
		switch in[i] {
		case '.':
			fmt.Printf("%s", string(memStack[ptr]))
		case ',':
			i++
			if i >= len(in) {
				panic("not input")
			}
			memStack[ptr] = in[i]
		case '>':
			ptr++
			if ptr > len(memStack) {
				panic("stackoverflow")
			}
		case '<':
			ptr--
			if ptr < 0 {
				panic("stackoverflow")
			}
		case '+':
			memStack[ptr] = memStack[ptr] + 1
		case '-':
			memStack[ptr] = memStack[ptr] - 1
		case '[':
			if memStack[ptr] == 0 {
				for in[i] != ']' {
					i++
				}
				i++
			} else {
				loopStack.Push(i)
			}
		case ']':
			if memStack[ptr] != 0 {
				i = loopStack.Pop() - 1
			} else {
				loopStack.Pop()
			}
		default:
		}
	}
	fmt.Println("")
	return memStack
}

func main() {
	in, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	out := brainfuck(in)
	if debug {
		fmt.Println(out)
	}
}
