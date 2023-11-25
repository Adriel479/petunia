package main

import (
	"fmt"
	"petunia/banco/clientes"
	"petunia/banco/contas"
)

type Contrato interface {
	Sacar(float64) string
}

func PagarBoleto(conta Contrato, valor float64) {
	conta.Sacar(valor)
}

func main() {
	var (
		c1 = &contas.ContaCorrente{
			Titular: clientes.Titular{
				Nome:      "Bob",
				CPF:       "000.000.000-00",
				Profissao: "Analista de segurança",
			},
			NumeroAgencia: 123,
			NumeroConta:   12345,
		}
		c2 = &contas.ContaCorrente{
			Titular: clientes.Titular{
				Nome:      "Alice",
				CPF:       "123.456.789-12",
				Profissao: "Analista de segurança",
			},
			NumeroAgencia: 123,
			NumeroConta:   54321,
		}
		p1 = &contas.ContaPounpanca{
			Titular: clientes.Titular{
				Nome:      "Ana",
				CPF:       "00.456.444-12",
				Profissao: "Developer",
			},
			NumeroAgencia: 445,
			NumeroConta:   11223344,
		}
	)

	c1.Depositar(100)
	c2.Depositar(10)

	fmt.Println(c1)
	fmt.Println(c2)

	fmt.Println(c1.Sacar(10))
	fmt.Println(c2.Sacar(5))

	fmt.Println(c1)
	fmt.Println(c2)

	fmt.Println(c1.Depositar(100))
	fmt.Println(c2.Depositar(500))

	fmt.Println(c1)
	fmt.Println(c2)

	fmt.Println(c1.Transferir(80.00, c2))
	fmt.Println(c1)
	fmt.Println(c2)

	fmt.Println(c1.Depositar(-100))
	fmt.Println(c2.Depositar(-500))

	fmt.Println(c1.Sacar(-10))
	fmt.Println(c2.Sacar(-5))

	fmt.Println(p1.Depositar(100))

	PagarBoleto(c1, 40)
	PagarBoleto(p1, 56)
	fmt.Println(c1.ObterSaldo())
	fmt.Println(p1.ObterSaldo())

}
