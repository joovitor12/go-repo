package main

import (
	"fmt"
	"reflect"
)

func main() {
	//essa linguagem tambem consegue inferir tipos de variaveis e tambem pode usar o := para evitar usar o nome var
	nome := "Joao"
	//quando uma variavel nao possui valor (nao atribuida) ela comeca zerada (0, 0.0, "" e afins)
	var idade int
	//eh recomendavel usar um tipo especifico de float pois se nao especificar o go ira escolher o float64
	var versao = 1.3
	fmt.Println("Hello world. oi", idade)
	fmt.Println("Hello world.", nome)
	fmt.Println(versao)
	fmt.Println("tipo variavel", versao, "->", reflect.TypeOf(versao))
}
