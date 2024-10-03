package stack

import (
	"fmt"
	"log"
)

const MAX = 10

type Stack struct {
	a   [MAX]int
	top int
}

func NewStack() Stack {
	return Stack{
		a:   [MAX]int{},
		top: -1,
	}
}

func (s *Stack) push(x int) error {
	if s.top >= (MAX - 1) {
		return fmt.Errorf("Stack overflow")
	} else {
		s.top++
		s.a[s.top] = x
		return nil
	}
}

func (s *Stack) pop() error {
	if s.top == -1 {
		return fmt.Errorf("Stack underflow")
	} else {
		x := s.a[s.top]
		fmt.Printf("%d remove from stack \n", x)
		s.top--
		return nil
	}
}

func (s *Stack) peek() int {
	if s.top < 0 {
		fmt.Println("Stack is empty")
		return 0
	} else {
		x := s.a[s.top]
		return x
	}
}

func main() {
	stack := NewStack()
	for {
		var x int
		fmt.Print("Enter a number to push to stack : ")
		fmt.Scanln(&x)

		err := stack.push(x)
		if err != nil {
			log.Println(err)
			break
		}
	}

	for {
		err := stack.pop()
		if err != nil {
			log.Println(err)
			return
		}
	}
}
