package main

import (
	"fmt"
	"log"
)

const MAX = 10

type DoubleStack struct {
	a    [MAX]int
	top1 int
	top2 int
}

func NewDoubleStack() DoubleStack {
	return DoubleStack{
		a:    [MAX]int{},
		top1: -1,
		top2: MAX,
	}
}

func (s *DoubleStack) push1(x int) error {
	if s.top2 == s.top1+1 {
		return fmt.Errorf("Stack1 overflow ")
	} else {
		s.top1++
		s.a[s.top1] = x
		return nil
	}
}

func (s *DoubleStack) push2(x int) error {
	if s.top2 == s.top1+1 {
		return fmt.Errorf("Stack2 overflow")
	} else {
		s.top2--
		s.a[s.top2] = x
		return nil
	}
}

func (s *DoubleStack) pop1() error {
	if s.top1 <= -1 {
		return fmt.Errorf("Stack underflow")
	} else {
		x := s.a[s.top1]
		log.Printf("%v remove from stack\n", x)
		s.top1--
		return nil
	}
}

func (s *DoubleStack) pop2() error {
	if s.top2 >= MAX {
		return fmt.Errorf("Stack overflow")
	} else {
		x := s.a[s.top2]
		log.Printf("%v remove from stack\n", x)
		s.top2++
		return nil
	}
}

func main() {
	doubleStack := NewDoubleStack()
	doubleStack.push1(1)
	doubleStack.push1(2)
	doubleStack.push1(3)
	doubleStack.push1(4)
	doubleStack.push1(5)

	doubleStack.push2(6)
	doubleStack.push2(7)
	doubleStack.push2(8)
	doubleStack.push2(9)
	doubleStack.push2(10)

	for {
		if err := doubleStack.pop1(); err != nil {
			break
		}

		if err := doubleStack.pop2(); err != nil {
			break
		}

	}

}
