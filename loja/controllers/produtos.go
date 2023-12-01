package controllers

import (
	"log"
	"net/http"
	"petunia/loja/models"
	"strconv"
	"text/template"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.BuscarTodosOsProdutos()
	templates.ExecuteTemplate(w, "index", produtos)
}

func New(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "new", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		produto, err := parseProduto(r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if err = produto.Salvar(); err != nil {
			log.Println("[ERROR] erro ao salvar produto:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		log.Println("[INFO] Produto cadastrado...")
		http.Redirect(w, r, "/index", http.StatusMovedPermanently)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	var (
		produto = models.Produto{}
		err     error
	)
	if produto.ID, err = strconv.Atoi(r.URL.Query().Get("id")); err != nil {
		log.Println("[ERROR] erro ao converter ID do produto:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err = produto.Deletar(); err != nil {
		log.Println("[ERROR] erro ao deletar produto:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/index", http.StatusMovedPermanently)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	var (
		produto models.Produto
		err     error
	)
	if produto.ID, err = strconv.Atoi(r.URL.Query().Get("id")); err != nil {
		log.Println("[ERROR] erro ao converter ID do produto:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err = produto.Carregar(); err != nil {
		log.Println("[ERROR] erro ao carregar produto:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	templates.ExecuteTemplate(w, "edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		produto, err := parseProduto(r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if err = produto.Atualizar(); err != nil {
			log.Println("[ERROR] erro ao atualizar o produto:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		log.Println("[INFO] Produto atualizado...")
		http.Redirect(w, r, "/index", http.StatusMovedPermanently)
	}
}

func parseProduto(r *http.Request) (*models.Produto, error) {
	var (
		produto models.Produto
		err     error
	)
	produto.Nome = r.FormValue("nome")
	produto.Descricao = r.FormValue("descricao")
	if produto.Preco, err = strconv.ParseFloat(r.FormValue("preco"), 64); err != nil {
		log.Println("[ERROR] erro ao converter preço do produto:", err)
		return nil, err
	}
	if produto.Quantidade, err = strconv.Atoi(r.FormValue("quantidade")); err != nil {
		log.Println("[ERROR] erro ao converter quantidade do produto:", err)
		return nil, err
	}
	if r.FormValue("id") != "" {
		if produto.ID, err = strconv.Atoi(r.FormValue("id")); err != nil {
			log.Println("[ERROR] erro ao converter ID do produto:", err)
			return nil, err
		}
	}
	if err = produto.Atualizar(); err != nil {
		log.Println("[ERROR] erro ao atualizar produto:", err)
		return nil, err
	}
	return &produto, nil
}
