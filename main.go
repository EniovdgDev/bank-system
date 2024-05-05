package main

import (
	"fmt"
	"os"

	"bank/clientes"
	"bank/contas"
)

func main() {

	titular := clientes.Titular{Nome: "Enio Vieira", CPF: "946.063.510-55", Profissao: "Programador"}
	contaExemplo := contas.ContaCorrente{Titular: titular, NumeroAgencia: 123, NumeroConta: 587}
	contaExemplo.Depositar(100)

	contaPraReceber := contas.ContaCorrente{}

	fmt.Println("Bem vindo(a), ", contaExemplo.Titular.Nome)
	fmt.Println()

	for {
		opcoes(contaExemplo)
		input := escolher()

		switch input {
		case 1:
			fmt.Println("Qual o valor do deposito?")
			response, _ := contaExemplo.Depositar(inserirValor())
			fmt.Println(response)
		case 2:
			fmt.Println("Qual o valor do saque?")
			response := contaExemplo.Sacar(inserirValor())
			fmt.Println(response)
		case 3:
			fmt.Println("Qual o valor do boleto?")
			pagarBoleto(&contaExemplo, inserirValor())
		case 4:
			fmt.Println("Qual o valor para transferir?")
			response := contaExemplo.Transferir(inserirValor(), &contaPraReceber)
			fmt.Println(response)
		case 0:
			print("Saindo.")
			os.Exit(0)
		default:
			print("Opção invalida. Saindo...")
			os.Exit(-1)
		}
	}

}

func opcoes(conta contas.ContaCorrente) {
	println("Seu saldo atual: R$", conta.ObterSaldo())
	fmt.Println()
	fmt.Println("O que você deseja fazer?")
	fmt.Println("1- Depositar.")
	fmt.Println("2- Sacar.")
	fmt.Println("3- Pagar Boleto.")
	fmt.Println("4- Transferir dinheiro.")
	fmt.Println("0- Sair.")
}

func escolher() int {
	var escolha int
	fmt.Scan(&escolha)
	return escolha
}

func inserirValor() float64 {
	print("R$")
	var valorDigitado float64
	fmt.Scan(&valorDigitado)
	fmt.Println()
	return valorDigitado
}

func pagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}
