package main

import (
	"fmt"
	"runtime"
)

//go build main.go comand para generar ejecutable
//./main comando para correr archivo ejecutable
func main() {
	fmt.Println(runtime.GOOS, runtime.GOARCH)
}
