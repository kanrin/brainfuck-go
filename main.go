package main

import (
	"flag"
	"io"
	"os"
)

var stackSize int
var file string

func init() {
	flag.StringVar(&file, "f", "", "input file")
	flag.IntVar(&stackSize, "s", 128, "run stack size")
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

func brainfuck(in []byte) []byte {
	memStack := make([]byte, stackSize)
	loopStack := new(Stack)
	ptr := 0
	for i := 0; i < len(in); i++ {
		switch in[i] {
		case '.':
			print(string(memStack[ptr]))
		case ',':
			i++
			if i >= len(in) {
				panic("not input")
			}
			memStack[ptr] = in[i]
		case '>':
			ptr++
			if ptr >= stackSize {
				ptr = ptr - stackSize
			}
		case '<':
			ptr--
			if ptr < 0 {
				ptr = ptr + stackSize
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
	println("")
	return memStack
}

func main() {
	var in []byte
	var err error
	if file == "" {
		in, err = io.ReadAll(os.Stdin)
	} else {
		in, err = os.ReadFile(file)
	}
	if err != nil {
		panic(err)
	}
	brainfuck(in)
}
