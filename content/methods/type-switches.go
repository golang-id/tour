// +build OMIT

package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Dua kali %v adalah %v\n", v, v*2)
	case string:
		fmt.Printf("%q adalah %v bytes panjangnya\n", v, len(v))
	default:
		fmt.Printf("Saya tidak tipe %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
