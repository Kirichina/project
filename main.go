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

func abreCadastro(resposta http.ResponseWriter, requisicao *http.Request){
	page, erro := template.ParseFiles("template/cadastro.html")

	if erro != nil {
		fmt.Println("Erro em cadastro", erro)
		return
	}
	page.Execute(resposta, nil) 
}

func salvaDoador(resposta http.ResponseWriter, requisicao *http.Request){
	page, erro := template.ParseFiles("template/cadastro.html")
	
	if erro != nil {
		fmt.Println("Erro em cadastro", erro)
		return
	}
	

	if requisicao.Method == http.MethodPost{
	nomeDoador := requisicao.FormValue("nome")
	telefoneDoador := requisicao.FormValue("telefone")
	emailDoador := requisicao.FormValue("email")
	tipoDoador := requisicao.FormValue("tipo")

	fmt.Println("O nome do Doador é: ", nomeDoador)
	fmt.Println("O telefone do Doador é: ", telefoneDoador)
	fmt.Println("O email do Doador é: ", emailDoador)
	fmt.Println("O tipo do Doador é: ", tipoDoador)

	}
	page.Execute(resposta, nil)
}


func main() {
	//Criando um EndPoint
	http.HandleFunc("/inicial", abreIndex)
	http.HandleFunc("/cadastro", abreCadastro)
	http.HandleFunc("/salvadoador", salvaDoador)

	fmt.Print("Iniciando Servidor!!!!")
	//Subindo o servidor na porta 8080
	erro := http.ListenAndServe(":8080", nil)
	if erro != nil {
		fmt.Print("Servidor com Problemas")
	}
}