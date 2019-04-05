// +build no-build OMIT

package main

import "github.com/Go-ID-community/go-tour-id/pic"

type Image struct{}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
