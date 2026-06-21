package linkedlists_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/18-linked-lists"
)

func TestSinglyLinkedList(t *testing.T) {
	l := linkedlists.NewSingly()
	l.Prepend(2)
	l.Append(3)
	l.Prepend(1)
	got := l.ToSlice()
	want := []int{1, 2, 3}
	if len(got) != len(want) {
		t.Fatalf("got %v want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got %v want %v", got, want)
		}
	}
	if err := l.Remove(2); err != nil {
		t.Fatal(err)
	}
	if l.Size() != 2 {
		t.Fatalf("size = %d", l.Size())
	}
}

func TestDoublyLinkedList(t *testing.T) {
	l := linkedlists.NewDoubly()
	l.Append(1)
	l.Prepend(0)
	l.Append(2)
	if l.Size() != 3 {
		t.Fatalf("size = %d", l.Size())
	}
	if err := l.Remove(1); err != nil {
		t.Fatal(err)
	}
	if l.Size() != 2 {
		t.Fatalf("size = %d", l.Size())
	}
}

func TestCircularLinkedList(t *testing.T) {
	l := linkedlists.NewCircular()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	got := l.Traverse()
	if len(got) != 3 || got[0] != 1 {
		t.Fatalf("got %v", got)
	}
}

func BenchmarkSinglyAppend(b *testing.B) {
	l := linkedlists.NewSingly()
	for i := 0; i < b.N; i++ {
		l.Append(i)
	}
}
