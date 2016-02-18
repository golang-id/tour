// +build no-build OMIT

package main

import "fmt"

// fibonacci adalah sebuah fungsi yang mengembalikan sebuah fungsi yang
// mengembalikan sebuah integer.
func fibonacci() func() int {
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
