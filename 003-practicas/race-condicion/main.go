package main

import (
	"fmt"
	"runtime"
	"sync"
)

//para compilar es con el siguiente comando  para validar si hay
//race condicion go run -race main.go
func main() {
	var wg sync.WaitGroup

	incremento := 0
	gs := 100

	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := incremento
			//se sede el proceso
			runtime.Gosched()
			v++
			incremento = v
			fmt.Println(incremento)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("el valor final del incremeto es ", incremento)
}
