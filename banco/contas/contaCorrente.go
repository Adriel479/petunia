package contas

import "petunia/banco/clientes"

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valor float64) string {
	if valor > c.saldo || valor < 0 {
		return "Saldo insuficiente!"
	}
	c.saldo -= valor
	return "Saque realizado com sucesso!"
}

func (c *ContaCorrente) Depositar(valor float64) (string, float64) {
	if valor > 0 {
		c.saldo += valor
		return "Deposito realizado com sucesso!", c.saldo
	}
	return "Valor do depósito menor que zero!", c.saldo
}

func (c *ContaCorrente) Transferir(valor float64, dest *ContaCorrente) bool {
	if c.saldo > valor && valor > 0 {
		c.Sacar(valor)
		dest.Depositar(valor)
		return true
	}
	return false
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
