package controllers

import (
	"net/http"
)

func ContatosHandler(w http.ResponseWriter, r *http.Request) {
	searchPath := r.Context().Value(orgCpfCnpjKey).(string)
	data := map[string]interface{}{
		"searchPath": searchPath,
	}
	templates.ExecuteTemplate(w, "contatos.html", data)
}

func EmpresasHandler(w http.ResponseWriter, r *http.Request) {
	searchPath := r.Context().Value(orgCpfCnpjKey).(string)
	data := map[string]interface{}{
		"searchPath": searchPath,
	}
	templates.ExecuteTemplate(w, "empresas.html", data)
}
