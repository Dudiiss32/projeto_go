package main 
// a linguagem golang funciona com pacotes, como o main

import "fmt"


func teste(){

	// TIPOS DE VARIAVEIS EM GO ===========================================================

	// var text = "Estudar go"
	// text := "Estudar go" // forma reduzida de declarar variaveis
	// isso é um slice 	
	var taskItems1 = []string {"Fazer a tarefa de casa", "Comer"}
	

	// isso é um array, qnd eu limito o tamanho de itens
	// var taskItems = [20]string {"Fazer a tarefa de casa", "Comer"}
	
	// fmt.Println(text)

	// LOOPS EM GO ===========================================================

	// for index, task := range taskItems1 {
	// 	// fmt.Println(index + 1, "-", task)

	// 	// Posso especificar o formato das variáveis usando o Printf, o %d represenda um decimal e %s uma string
	// 	fmt.Printf("%d. %s\n", index+1, task)

	// }

	// se não queremos usar o index, podemos usar o _
	// for _, task := range taskItems {
	// 	fmt.Println(task)
	// }
	taskItems1 = addTasks(taskItems1, "Fazer comida")
	printTasks(taskItems1)
}

func printTasks (arr []string){
	for index, task := range arr {
		// fmt.Println(index + 1, "-", task)

		// Posso especificar o formato das variáveis usando o Printf, o %d represenda um decimal e %s uma string
		fmt.Printf("%d. %s\n", index+1, task)

	}
}

// depois dos parâmetros se quisermos retornar algo na função, precisamos dizer o tipo da var

func addTasks (arr []string, newTask string) []string{
	var newArray = append(arr, newTask)
	return newArray
}