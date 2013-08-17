// +build no-build OMIT

package main

import "fmt"

// fibonacci merupakan fungsi yang mengembalikan
// sebuah fungsi dengan nilai balik sebuah int.
func fibonacci() func() int {
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
