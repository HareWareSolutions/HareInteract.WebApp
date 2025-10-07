package controllers

import (
	"net/http"
)

// Redireciona para página de Atividades
func AtividadesHandler(w http.ResponseWriter, r *http.Request) {
	searchPath, ok := r.Context().Value(orgCpfCnpjKey).(string)

	if !ok {
		http.Error(w, "Informações de sessão não encontradas.", http.StatusUnauthorized)
		return
	}

	data := map[string]interface{}{
		"searchPath": searchPath,
	}
	templates.ExecuteTemplate(w, "atividades.html", data)
}
