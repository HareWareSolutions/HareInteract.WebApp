package controllers

import "net/http"

func WppHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "whatsapp.html", nil)
}

func InstanciaHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "instancia.html", nil)
}
