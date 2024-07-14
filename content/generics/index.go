//go:build OMIT

package main

import "fmt"

// Index mengembalikan indeks pada slice s yang bernilai x, atau -1 bila x
// tidak ditemukan.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v dan x bertipe T, yang memiliki syarat "comparable",
		// sehingga dapat menggunakan operator "==".
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	// Index dapat digunakan untuk slice bertipe int.
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Index juga dapat digunakan untuk slice bertipe string.
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))
}
