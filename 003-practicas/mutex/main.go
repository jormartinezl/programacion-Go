package main

import (
	"fmt"
	"sync"
)

//para compilar es con el siguiente comando  para validar si hay
//race condicion go run -race main.go
func main() {
	var wg sync.WaitGroup

	incremento := 0
	gs := 100

	wg.Add(gs) //se declaran los gorutines
	var m sync.Mutex
	for i := 0; i < gs; i++ {
		go func() {
			m.Lock() //se bloquean las variables para eliminar race condicion

			v := incremento
			v++
			incremento = v
			fmt.Println(incremento)

			m.Unlock() //se desbloquea
			wg.Done()  //se acaba la gorutine
		}()
	}
	wg.Wait() //le decimos al programa que espere a que acaben las goroutines
	fmt.Println("el valor final del incremeto es ", incremento)
}
