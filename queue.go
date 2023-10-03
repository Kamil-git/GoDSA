package main

import (
	"fmt"
)

// Struktura Queue reprezentująca kolejkę.
type Queue struct {
	items []int
}

// Funkcja do dodawania elementu do kolejki.
func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}
func NewQueue() *Queue {
	return &Queue{}
}

// Funkcja do usuwania elementu z kolejki i zwracania go.
func (q *Queue) Dequeue() int {
	if len(q.items) == 0 {
		fmt.Println("Kolejka jest pusta.")
		return -1 // Możesz dostosować wartość zwracaną w przypadku błędu.
	}
	dequeuedItem := q.items[0]
	q.items = q.items[1:]
	return dequeuedItem
}

// Funkcja do sprawdzania, czy kolejka jest pusta.
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// Funkcja do pobierania liczby elementów w kolejce.
func (q *Queue) Size() int {
	return len(q.items)
}
func (q *Queue) example() {
	// Tworzenie nowej kolejki.
	queue := Queue{}

	// Dodawanie elementów do kolejki.
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	// Wyświetlanie rozmiaru kolejki.
	fmt.Printf("Rozmiar kolejki: %d\n", queue.Size())

	// Usuwanie i wyświetlanie elementów z kolejki.
	fmt.Printf("Usunięty element: %d\n", queue.Dequeue())
	fmt.Printf("Usunięty element: %d\n", queue.Dequeue())

	// Wyświetlanie rozmiaru kolejki po usunięciu elementów.
	fmt.Printf("Rozmiar kolejki po usunięciu: %d\n", queue.Size())

	// Sprawdzanie, czy kolejka jest pusta.
	if queue.IsEmpty() {
		fmt.Println("Kolejka jest pusta.")
	} else {
		fmt.Println("Kolejka nie jest pusta.")
	}
}
