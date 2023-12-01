package main

import (
	"net/http"
	"petunia/loja/routes"
)

func main() {
	routes.CarregarRotas()
	http.ListenAndServe(":8000", nil)
}
