package main

import "fmt" 

func main(){
	var usuario string
	var senha string
fmt.Println("digite seu usuário")
fmt.Scan(&usuario)
 if usuario == "admin"{
	fmt.Println("digite sua senha")
	fmt.Scan(&senha)
	if senha == "1234"{
		fmt.Println("senha correta")
	}else{
		fmt.Println("senha incorreta, acesso negado! tente novamente")
	}

 } else {
	fmt.Println("usuário incorreto, tente novamente")
 }
	
} 
