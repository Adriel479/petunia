package models

import (
	"log"
	"petunia/loja/db"
)

type Produto struct {
	ID         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func (p *Produto) Carregar() error {
	conexao := db.ConectaComBancoDeDados()
	defer conexao.Close()
	result := conexao.QueryRow("SELECT nome, descricao, preco, quantidade FROM produtos WHERE id = $1;", p.ID)
	if err := result.Scan(
		&p.Nome,
		&p.Descricao,
		&p.Preco,
		&p.Quantidade,
	); err != nil {
		return err
	}
	return nil
}

func (p *Produto) Atualizar() error {
	conexao := db.ConectaComBancoDeDados()
	defer conexao.Close()
	if _, err := conexao.Exec("UPDATE produtos SET nome = $1, descricao = $2, preco = $3, quantidade = $4 WHERE id = $5;",
		p.Nome,
		p.Descricao,
		p.Preco,
		p.Quantidade,
		p.ID,
	); err != nil {
		return err
	}
	return nil
}

func (p *Produto) Deletar() error {
	conexao := db.ConectaComBancoDeDados()
	defer conexao.Close()
	if _, err := conexao.Exec("DELETE FROM produtos WHERE id = $1;", p.ID); err != nil {
		return err
	}
	return nil
}

func (p *Produto) Salvar() error {
	conexao := db.ConectaComBancoDeDados()
	defer conexao.Close()
	_, err := conexao.Exec(
		"INSERT INTO produtos (nome, descricao, preco, quantidade) VALUES ($1, $2, $3, $4);",
		p.Nome, p.Descricao, p.Preco, p.Quantidade,
	)
	return err
}

func BuscarTodosOsProdutos() []Produto {
	conexao := db.ConectaComBancoDeDados()
	defer conexao.Close()
	produtos := make([]Produto, 0)
	query, err := conexao.Query("SELECT * FROM produtos;")
	if err != nil {
		log.Fatal(err)
	}
	defer query.Close()
	for query.Next() {
		var p Produto
		if err = query.Scan(
			&p.ID,
			&p.Nome,
			&p.Descricao,
			&p.Preco,
			&p.Quantidade,
		); err != nil {
			log.Fatal(err)
		}
		produtos = append(produtos, p)
	}
	return produtos
}
