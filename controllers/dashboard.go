package controllers

import (
	"net/http"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "dashboard.html", nil)
}

func ContasHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "contas.html", nil)
}

func VendasHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "vendas.html", nil)
}

func MarketingHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "marketing.html", nil)
}

func AtendimentoHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "atendimento.html", nil)
}
func TimelineHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "timeline.html", nil)
}

func IntegracoesHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "integracoes.html", nil)
}

func ConfiguracoesHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "configuracoes.html", nil)
}
