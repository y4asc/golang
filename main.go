package main

import "fmt" 

func main(){
	a, b := 10, 3
	fmt.Println("a soma é:", a + b)
	fmt.Println("a subtração é:", a - b)
	fmt.Println("a divisão é:", a / b)
	fmt.Println("a multiplicação é:", a * b)
	fmt.Println("o resto da divisão:", a % b)

	a+=1
	fmt.Println("incrementa a", a)

	if a > 0 && b > 0 {
		fmt.Println("numeros positivos")
} }
