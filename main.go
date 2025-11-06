package main 
// a linguagem golang funciona com pacotes, como o main

import "fmt"

func main(){

	// TIPOS DE VARIAVEIS EM GO ===========================================================

	// var text = "Estudar go"
	// text := "Estudar go" // forma reduzida de declarar variaveis
	// isso é um slice 
	var taskItems = []string {"Fazer a tarefa de casa", "Comer"}

	// isso é um array, qnd eu limito o tamanho de itens
	// var taskItems = [20]string {"Fazer a tarefa de casa", "Comer"}
	

	// fmt.Println(text)
	// fmt.Println("Hello other line")
	// fmt.Println(taskItems)


	// LOOPS EM GO ===========================================================

	for index, task := range taskItems {
		fmt.Println(index + 1, "-", task)
	}

	// se não queremos usaro index, podemos usar o _
	// for _, task := range taskItems {
	// 	fmt.Println(task)
	// }
}