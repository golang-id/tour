module github.com/golang-id/tour

go 1.24.0

require (
	git.sr.ht/~shulhan/pakakeh.go v0.60.3-0.20260202092811-1e3bb9be8444
	golang.org/x/tools v0.40.0
	golang.org/x/tools/godoc v0.1.0-deprecated
)

require (
	github.com/yuin/goldmark v1.7.13 // indirect
	golang.org/x/mod v0.31.0 // indirect
	golang.org/x/net v0.48.0 // indirect
	golang.org/x/sync v0.19.0 // indirect
	golang.org/x/sys v0.39.0 // indirect
)

//replace golang.org/x/tools => ../../../golang.org/x/tools

//replace golang.org/x/net => ../../../golang.org/x/net
