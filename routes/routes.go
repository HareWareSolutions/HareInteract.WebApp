package routes

import (
	"HareInteract.WebApp/controllers"
	"net/http"
)

func CarregaRotas() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", controllers.LoginHandler)
}
