package main

import (
	"fmt"
	"net/http"
)

func main(){
	fmt.Println("Hello world")

	http.HandleFunc("/", HelloUser)

	http.ListenAndServe(":8080", nil)

}

func HelloUser(writer http.ResponseWriter, request *http.Request){
	var boasVindas = "Olá usuário, está funcionando"
	fmt.Fprintln(writer, boasVindas)
}