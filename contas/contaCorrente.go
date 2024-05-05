package contas

import (
	"bank/clientes"
	"fmt"
)

type ContaCorrente struct {
	Titular                    clientes.Titular
	NumeroAgencia, NumeroConta int
	saldo                      float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, string) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso.", c.ObterSaldo()
	} else {
		return "Valor do deposito é menor que zero.", c.ObterSaldo()
	}

}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) string {
	if valorDaTransferencia <= c.saldo && valorDaTransferencia > 0 {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return "Transferencia realizada com sucesso"
	} else {
		return "Saldo insuficiente."
	}
}

func (c *ContaCorrente) ObterSaldo() string {
	saldoFormatado := fmt.Sprintf("%.2f", c.saldo)
	return saldoFormatado
}
