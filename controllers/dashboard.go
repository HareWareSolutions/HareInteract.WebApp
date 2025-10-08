package controllers

import (
	"fmt"
	"net/http"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	searchPath, ok := r.Context().Value(orgCpfCnpjKey).(string)
	if !ok {
		http.Error(w, "Informações de sessão não encontradas.", http.StatusUnauthorized)
		return
	}

	data := map[string]interface{}{
		"searchPath": searchPath,
	}
	templates.ExecuteTemplate(w, "dashboard.html", data)
}

func ContasHandler(w http.ResponseWriter, r *http.Request) {
	searchPath, ok := r.Context().Value(orgCpfCnpjKey).(string)

	if !ok {
		http.Error(w, "Informações de sessão não encontradas.", http.StatusUnauthorized)
		return
	}

	data := map[string]interface{}{
		"searchPath": searchPath,
	}
	templates.ExecuteTemplate(w, "contas.html", data)
}

func VendasHandler(w http.ResponseWriter, r *http.Request) {
	searchPath, ok := r.Context().Value(orgCpfCnpjKey).(string)
	if !ok {
		http.Error(w, "Informações de sessão não encontradas.", http.StatusUnauthorized)
		return
	}

	data := map[string]interface{}{
		"searchPath": searchPath,
	}
	templates.ExecuteTemplate(w, "vendas.html", data)
}

func MarketingHandler(w http.ResponseWriter, r *http.Request) {
	searchPath, ok := r.Context().Value(orgCpfCnpjKey).(string)
	if !ok {
		http.Error(w, "Informações de sessão não encontradas.", http.StatusUnauthorized)
		return
	}

	data := map[string]interface{}{
		"searchPath": searchPath,
	}
	templates.ExecuteTemplate(w, "marketing.html", data)
}

func AtendimentoHandler(w http.ResponseWriter, r *http.Request) {
	searchPath, ok := r.Context().Value(orgCpfCnpjKey).(string)
	if !ok {
		http.Error(w, "Informações de sessão não encontradas.", http.StatusUnauthorized)
		return
	}

	data := map[string]interface{}{
		"searchPath": searchPath,
	}
	templates.ExecuteTemplate(w, "atendimento.html", data)
}

func TimelineHandler(w http.ResponseWriter, r *http.Request) {
	searchPath, ok := r.Context().Value(orgCpfCnpjKey).(string)
	if !ok {
		http.Error(w, "Informações de sessão não encontradas.", http.StatusUnauthorized)
		return
	}

	data := map[string]interface{}{
		"searchPath": searchPath,
	}
	templates.ExecuteTemplate(w, "timeline.html", data)
}

func IntegracoesHandler(w http.ResponseWriter, r *http.Request) {
	searchPath, ok := r.Context().Value(orgCpfCnpjKey).(string)
	if !ok {
		http.Error(w, "Informações de sessão não encontradas.", http.StatusUnauthorized)
		return
	}

	data := map[string]interface{}{
		"searchPath": searchPath,
	}
	templates.ExecuteTemplate(w, "integracoes.html", data)
}

func ConfiguracoesHandler(w http.ResponseWriter, r *http.Request) {
	searchPath, ok := r.Context().Value(orgCpfCnpjKey).(string)
	if !ok {
		http.Error(w, "Informações de sessão não encontradas.", http.StatusUnauthorized)
		return
	}

	user, err := PerfilConfigHandler(r)

	fmt.Println(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	data := map[string]interface{}{
		"searchPath": searchPath,
		"user":       user,
	}

	templates.ExecuteTemplate(w, "configuracoes.html", data)
}
