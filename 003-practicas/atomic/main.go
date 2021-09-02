package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//para compilar es con el siguiente comando  para validar si hay
//race condicion go run -race main.go
func main() {
	var wg sync.WaitGroup

	var incremento int64
	gs := 100

	wg.Add(gs) //se declaran los gorutines

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&incremento, 1)
			fmt.Println(atomic.LoadInt64(&incremento))
			wg.Done() //se acaba la gorutine
		}()
	}
	wg.Wait() //le decimos al programa que espere a que acaben las goroutines
	fmt.Println("el valor final del incremeto es ", incremento)
}
