package main

// a linguagem golang funciona com pacotes, como o main

import "fmt"

func main() {

	// TIPOS DE VARIAVEIS EM GO ===========================================================

	// var text = "Estudar go"
	// text := "Estudar go" // forma reduzida de declarar variaveis
	// isso é um slice
	var taskItems = []string{"Fazer a tarefa de casa", "Comer"}

	// isso é um array, qnd eu limito o tamanho de itens
	// var taskItems = [20]string {"Fazer a tarefa de casa", "Comer"}

	// fmt.Println(text)
	// fmt.Println("Hello other line")
	// fmt.Println(taskItems)

	// LOOPS EM GO ===========================================================

	for index, task := range taskItems {
		fmt.Println(index+1, "-", task)
	}

	// se não queremos usar o index, podemos usar o _
	// for _, task := range taskItems {
	// 	fmt.Println(task)
	// }

	// ======================================================================
	// FUNÇÕES EM GO ========================================================
	// criando uma função simples para somar dois números
	resultado := somar(10, 20)
	fmt.Println("Resultado da soma:", resultado)

	// chamando função que não retorna nada (void)
	mostrarMensagem("Estudando Go todo dia é brabo")

	// ======================================================================
	// STRUCTS EM GO ========================================================
	// struct é tipo um "objeto", onde defino campos e valores
	type Usuario struct {
		nome  string
		idade int
	}

	// criando um usuário
	user := Usuario{nome: "Maria", idade: 18}
	fmt.Println("Usuário:", user.nome, "-", user.idade, "anos")

	// ======================================================================
	// MAPS EM GO ===========================================================
	// map é tipo um objeto de chave-valor

	config := map[string]string{
		"porta":     "8080",
		"ambiente":  "desenvolvimento",
		"mensagem":  "rodando API fake",
	}

	fmt.Println("Configuração da aplicação:", config["mensagem"])

	// ======================================================================
	// SIMULANDO UMA "API" (BEM SIMPLES) ===================================
	// isso aqui não é uma API real ainda, só simulando uma função q recebe dados
	// como se fosse um endpoint
	resposta := endpointSimples("Pegando lista de tasks...")
	fmt.Println(resposta)
}

// criando funções fora do main ==========================================

func somar(a int, b int) int {
	// retorna a soma dos números
	return a + b
}

func mostrarMensagem(msg string) {
	// só imprime uma mensagem
	fmt.Println("Mensagem:", msg)
}

func endpointSimples(route string) string {
	// simulando processamento de requisição
	// tipo um mini endpoint fake
	return "REQUEST OK: " + route
}
