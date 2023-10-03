package main

import (
	"errors"
	"fmt"
)

const stackSize = 10

type Stack struct {
	container [stackSize]int
	top       int
}

func NewStack() *Stack {
	return &Stack{top: -1}
}

func (s *Stack) Push(item int) error {
	if s.top == stackSize-1 {
		return errors.New("StackOverflow")
	}
	s.top++
	s.container[s.top] = item
	return nil
}

func (s *Stack) Pop() (int, error) {
	if s.top == -1 {
		return 0, errors.New("Stack is empty")
	}
	item := s.container[s.top]
	s.container[s.top] = 0 // Wyczyszczenie wartości w tablicy
	s.top--
	return item, nil
}
func (s *Stack) example() {
	//Przykłady użycia
	stack := NewStack()
	for i := 1; i <= stackSize; i++ {
		if err := stack.Push(i); err != nil {
			fmt.Println("Błąd push:", err)
		}
	}

	for i := 0; i < stackSize; i++ {
		item, err := stack.Pop()
		if err != nil {
			fmt.Println("Błąd pop:", err)
			break
		}
		fmt.Printf("%d ", item)
	}

	fmt.Println()
}
