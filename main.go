package main

import (
	"net/http"

	"go-fundamentos-aplicacao-web/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":3001", nil)
}
