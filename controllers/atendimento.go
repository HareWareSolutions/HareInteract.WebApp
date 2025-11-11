package controllers

import (
	"net/http"

	"HareInteract.WebApp/models/apperr"
)

func TicketHandler(w http.ResponseWriter, r *http.Request) {
	searchPath, ok := r.Context().Value(orgCpfCnpjKey).(string)

	if !ok {
		err := apperr.Erro{
			Mensagem: "Informações de sessão não encontradas.",
			Status:   http.StatusUnauthorized,
		}

		templates.ExecuteTemplate(w, "erro.html", err)
	}

	data := map[string]interface{}{
		"searchPath": searchPath,
	}
	templates.ExecuteTemplate(w, "tickets.html", data)
}

func AgendamentoHandler(w http.ResponseWriter, r *http.Request) {
	searchPath, ok := r.Context().Value(orgCpfCnpjKey).(string)

	if !ok {
		err := apperr.Erro{
			Mensagem: "Informações de sessão não encontradas.",
			Status:   http.StatusUnauthorized,
		}

		templates.ExecuteTemplate(w, "erro.html", err)
	}

	data := map[string]interface{}{
		"searchPath": searchPath,
	}

	templates.ExecuteTemplate(w, "agendamento.html", data)
}
