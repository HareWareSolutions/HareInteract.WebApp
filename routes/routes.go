package routes

import (
	"HareInteract.WebApp/controllers"
	"net/http"
)

func CarregaRotas() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", controllers.LoginHandler)

	http.HandleFunc("/dashboard", controllers.DashboardHandler)
	http.HandleFunc("/dashboard/contas", controllers.ContasHandler)
	http.HandleFunc("/dashboard/vendas", controllers.VendasHandler)
	http.HandleFunc("/dashboard/marketing", controllers.MarketingHandler)
	http.HandleFunc("/dashboard/atendimento", controllers.AtendimentoHandler)
	http.HandleFunc("/dashboard/timeline", controllers.TimelineHandler)
	http.HandleFunc("/dashboard/integracoes", controllers.IntegracoesHandler)
	http.HandleFunc("/dashboard/configuracoes", controllers.ConfiguracoesHandler)
}
