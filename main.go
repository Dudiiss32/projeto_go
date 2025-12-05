package main

// a linguagem golang funciona com pacotes, como o main

import (
	"fmt"
	"net/http"
)

// criando um slice global (para a API conseguir acessar)
var taskItems = []string{"Fazer a tarefa de casa", "Comer"}

func main() {

	// ======================================================================
	// ROTAS DA API =========================================================
	// definindo rotas reais com http.HandleFunc

	http.HandleFunc("/", HelloUser)
	http.HandleFunc("/tasks", GetTasks)
	http.HandleFunc("/add-task", AddTask)

	// iniciando o servidor na porta 8080
	fmt.Println("Servidor rodando na porta 8080...")
	http.ListenAndServe(":8080", nil)

	// ======================================================================

	// criando uma função simples para somar dois números
	resultado := somar(10, 20)
	fmt.Println("Resultado da soma:", resultado)

	// chamando função que não retorna nada (void)
	mostrarMensagem("Estudando Go todo dia")

	// ======================================================================
	// STRUCTS EM GO ========================================================
	type Usuario struct {
		nome  string
		idade int
	}

	user := Usuario{nome: "Maria", idade: 18}
	fmt.Println("Usuário:", user.nome, "-", user.idade, "anos")

	// ======================================================================
	// MAPS EM GO ===========================================================
	config := map[string]string{
		"porta":     "8080",
		"ambiente":  "desenvolvimento",
		"mensagem":  "rodando API",
	}

	fmt.Println("Configuração da aplicação:", config["mensagem"])

	// ======================================================================
	// SIMULANDO UMA "API" =================================
	resposta := endpointSimples("Pegando lista de tasks...")
	fmt.Println(resposta)
}

// ======================================================================
// FUNÇÕES DE API =======================================================

// rota principal
func HelloUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bem-vindo à minha API em Go!")
}

// rota para pegar tasks
func GetTasks(w http.ResponseWriter, r *http.Request) {
	// escrevendo tasks na resposta
	fmt.Fprintf(w, "Lista de tasks:\n")

	for i, task := range taskItems {
		fmt.Fprintf(w, "%d - %s\n", i+1, task)
	}
}

// rota para adicionar task via URL
// exemplo: http://localhost:8080/add-task?item=Estudar%20Go
func AddTask(w http.ResponseWriter, r *http.Request) {

	// pegando o valor enviado na URL
	item := r.URL.Query().Get("item")

	if item == "" {
		fmt.Fprintf(w, "Você precisa enviar o parâmetro 'item'")
		return
	}

	// adicionando ao slice global
	taskItems = append(taskItems, item)

	fmt.Fprintf(w, "Item adicionado com sucesso: %s\n", item)
}

// ======================================================================
// OUTRAS FUNÇÕES ===================================

func somar(a int, b int) int {
	return a + b
}

func mostrarMensagem(msg string) {
	fmt.Println("Mensagem:", msg)
}

func endpointSimples(route string) string {
	return "REQUEST OK: " + route
}
