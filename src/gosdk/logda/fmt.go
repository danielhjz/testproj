package logda

import "fmt"

// wont be import by go get github.com/danielhjz/testproj/gosdk
func Print(a interface{}) {
	fmt.Println("hi", a)
}
