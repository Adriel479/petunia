package contas

import "petunia/banco/clientes"

type ContaPounpanca struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaPounpanca) Sacar(valor float64) string {
	if valor > c.saldo || valor < 0 {
		return "Saldo insuficiente!"
	}
	c.saldo -= valor
	return "Saque realizado com sucesso!"
}

func (c *ContaPounpanca) Depositar(valor float64) (string, float64) {
	if valor > 0 {
		c.saldo += valor
		return "Deposito realizado com sucesso!", c.saldo
	}
	return "Valor do depósito menor que zero!", c.saldo
}

func (c *ContaPounpanca) ObterSaldo() float64 {
	return c.saldo
}
