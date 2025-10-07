package controllers

import "net/http"

// Redireciona para página Leads
func LeadHandler(w http.ResponseWriter, r *http.Request) {
	searchPath, ok := r.Context().Value(orgCpfCnpjKey).(string)

	if !ok {
		http.Error(w, "Informações de sessão não encontradas.", http.StatusUnauthorized)
		return
	}

	data := map[string]interface{}{
		"searchPath": searchPath,
	}
	templates.ExecuteTemplate(w, "leads.html", data)
}

// Redireciona para página Oportunidades
func OportunidadeHandler(w http.ResponseWriter, r *http.Request) {
	searchPath, ok := r.Context().Value(orgCpfCnpjKey).(string)

	if !ok {
		http.Error(w, "Informações de sessão não encontradas.", http.StatusUnauthorized)
		return
	}

	data := map[string]interface{}{
		"searchPath": searchPath,
	}
	templates.ExecuteTemplate(w, "oportunidades.html", data)
}
