package main

import (
	"net/http"

	"HareInteract.WebApp/routes"
)

func main() {
	routes.CarregaRotas()

	http.ListenAndServe(":8000", nil)
}
