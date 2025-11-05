package main 
// a linguagem golang funciona com pacotes, como o main

import "fmt"

func main(){

	// var text = "Duda é linda"
	text := "Duda é linda"
	// isso é um slice 
	var taskItems = []string {"Fazer a tarefa de casa", "Comer"}

	// isso é um array, qnd eu limito o tamanho de itens
	// var taskItems = [20]string {"Fazer a tarefa de casa", "Comer"}
	


	fmt.Println(text)
	fmt.Println("Hello other line")
	fmt.Println(taskItems)
}