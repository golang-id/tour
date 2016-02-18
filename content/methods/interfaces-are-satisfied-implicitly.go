// +build OMIT

package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// Method berikut berarti type T mengimplementasikan interface I,
// tapi kita tidak perlu secara eksplisit mendeklarasikannya.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}
