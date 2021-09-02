package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Printf("Numero de CPUS: %v\n", runtime.NumCPU())
	fmt.Printf("Numero de Goroutines: %v\n\n", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("hola desde la primera gorutine")
		wg.Done()
	}()
	go func() {
		fmt.Println("hola desde la segunda gorutine")
		wg.Done()
	}()

	fmt.Printf("\nNumero de CPUS: %v\n", runtime.NumCPU())
	fmt.Printf("Numero de Goroutines: %v\n", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("a punto de finalizar main.")
	fmt.Printf("\nNumero de CPUS: %v\n", runtime.NumCPU())
	fmt.Printf("Numero de Goroutines: %v\n", runtime.NumGoroutine())
}
