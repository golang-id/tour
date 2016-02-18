// +build OMIT

package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	// append bekerja pada slice yang nil.
	s = append(s, 0)
	printSlice(s)

	// Slice bertambah seperlunya.
	s = append(s, 1)
	printSlice(s)

	// Kita juga bisa menambahkan lebih dari satu elemen sekaligus.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
