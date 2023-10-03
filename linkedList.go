package main

import "fmt"

// Struktura reprezentująca pojedynczy element listy.
type Node struct {
	Data int   // Przechowywane dane.
	Next *Node // Wskaźnik do następnego elementu.
}

// Struktura reprezentująca listę.
type LinkedList struct {
	Head *Node // Wskaźnik do głowy listy (pierwszego elementu).
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

// Funkcja do dodawania nowego elementu na koniec listy.
func (ll *LinkedList) Append(data int) {
	newNode := &Node{Data: data, Next: nil}

	// Jeśli lista jest pusta, ustaw nowy element jako głowę.
	if ll.Head == nil {
		ll.Head = newNode
		return
	}

	// W przeciwnym razie przeszukujemy listę, aż dotrzemy do ostatniego elementu.
	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}

	// Dodajemy nowy element na koniec listy.
	current.Next = newNode
}

// Funkcja do wyświetlania zawartości listy.
func (ll *LinkedList) Display() {
	current := ll.Head
	for current != nil {
		fmt.Printf("%d -> ", current.Data)
		current = current.Next
	}
	fmt.Println("nil")
}

func (l *LinkedList) example() {
	// Tworzenie nowej listy.
	list := LinkedList{}

	// Dodawanie elementów do listy.
	list.Append(1)
	list.Append(2)
	list.Append(3)

	// Wyświetlanie zawartości listy.
	list.Display()
}
