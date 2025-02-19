package main

import ("fmt"
		"html/template"
		"net/http")

func abreIndex(resposta http.ResponseWriter, requisicao *http.Request){
	page, erro := template.ParseFiles("template/index.html")
	
	if erro != nil {
		fmt.Println("o erro foi", erro)
		return
	}
	page.Execute(resposta, nil)
}

func main() {
	//Criando um EndPoint
	http.HandleFunc("/", abreIndex)

	//Subindo o servidor na porta 8080
	erro := http.ListenAndServe(":8080", nil)
	if erro != nil {
		fmt.Print("Servidor com Problemas")
	}
	fmt.Print("ohayou sekai good morning world");
}