package collections

type Queue[T any] interface {
	// Enqueue fügt ein Element am Ende der Schlange hinzu
	Enqueue(value T)
	// Deque entfernt das erste Element vom Anfang der Schlange
	Deque() T
}
