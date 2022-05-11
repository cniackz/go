package main

import "fmt"

func printArrey(arreglo []int) {
	fmt.Println("")
	fmt.Println("Arrey value:")
	for i := 0; i < len(arreglo); i++ {
		fmt.Println(arreglo[i])
	}
}

func main() {

	arregloA := make([]int, 4)
	arregloA[0] = 1
	arregloA[1] = 2
	arregloA[2] = 3
	arregloA[3] = 4
	arregloB := make([]int, 4)
	counter := 0
	for i := len(arregloA); i > 0; i-- {
		//fmt.Println(i)
		arregloB[counter] = arregloA[i-1]
		counter++
	}
	printArrey(arregloA)
	printArrey(arregloB)
}
