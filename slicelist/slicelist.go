package slicelist

import (
	"fmt"
	"strings"
)

type SliceList[E any] struct {
	slice []E
}

func (s *SliceList[E]) Push(value E) {
	s.slice = append(s.slice, value)
}

func (s *SliceList[E]) Pop() E {
	var value E
	var lastIndex = len(s.slice) - 1
	if lastIndex >= 0 {
		value = s.slice[lastIndex]
		s.slice = s.slice[:lastIndex]
	}
	return value
}

func (s *SliceList[E]) Enqueue(value E) {
	s.Push(value)
}

func (s *SliceList[E]) Deque() E {
	var value E
	if len(s.slice) > 0 {
		value = s.slice[0]
		s.slice = s.slice[1:]
	}
	return value
}

func (s *SliceList[E]) Delete(index int) {
	if index < 0 || index > len(s.slice)-1 {
		return
	}
	s.slice = append(s.slice[:index], s.slice[index+1:]...)
}

func (s *SliceList[E]) Get(index int) E {
	var value E
	if index > 0 && index < len(s.slice) {
		value = s.slice[index]
	}
	return value
}

func (s *SliceList[E]) Insert(index int, value E) {
	if index < 0 || index > len(s.slice)-1 {
		return
	}
	s.slice = append(s.slice[:index], append([]E{value}, s.slice[index:]...)...)
}

func (s SliceList[E]) Len() int {
	return len(s.slice)
}

func (s SliceList[E]) String() string {
	var sb strings.Builder
	sb.WriteRune('[')
	for index, element := range s.slice {
		if index > 0 {
			sb.WriteString(", ")
		}
		fmt.Fprint(&sb, element)
	}
	sb.WriteRune(']')
	return sb.String()
}

func New[T any]() *SliceList[T] {
	return &SliceList[T]{}
}
