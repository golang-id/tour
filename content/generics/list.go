//go:build OMIT

package main

// List merepresentasikan linked-list yang dapat menyimpan nilai bertipe
// apa pun.
type List[T any] struct {
	next *List[T]
	val  T
}

func main() {
}
