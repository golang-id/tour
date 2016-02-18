// +build OMIT

package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // menunjuk ke i
	fmt.Println(*p) // baca i lewat pointer
	*p = 21         // set i lewat pointer
	fmt.Println(i)  // lihat nilai terbaru dari i

	p = &j         // p menunjuk ke j
	*p = *p / 37   // bagi nilai j lewat pointer
	fmt.Println(j) // lihat nilai terbaru dari j
}
