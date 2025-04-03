package main

import "fmt" 

func main(){
	var ages = [4]int{17, 16, 20, 40}
	nomes := [4]string{"mario", "luige", "deadpool", "superman"}
	fmt.Println(ages)
	fmt.Println(nomes)
	nomes[3] = "clarkman"
	fmt.Println(nomes)
	
} 
