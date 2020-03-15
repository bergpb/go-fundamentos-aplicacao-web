package main

import (
	"net/http"

	"./routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":3001", nil)
}
