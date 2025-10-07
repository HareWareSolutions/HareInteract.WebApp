package controllers

import "net/http"

func TicketHandler(w http.ResponseWriter, r *http.Request) {
	searchPath, ok := r.Context().Value(orgCpfCnpjKey).(string)

	if !ok {
		http.Error(w, "Informações de sessão não encontradas.", http.StatusUnauthorized)
		return
	}

	data := map[string]interface{}{
		"searchPath": searchPath,
	}
	templates.ExecuteTemplate(w, "tickets.html", data)
}
