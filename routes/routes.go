package routes

import (
	"HareInteract.WebApp/controllers"
	"net/http"
)

func CarregaRotas() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", controllers.LoginHandler)

	http.HandleFunc("/dashboard", controllers.AuthMiddleware(controllers.DashboardHandler))
	http.HandleFunc("/contas", controllers.AuthMiddleware(controllers.ContasHandler))
	http.HandleFunc("/vendas", controllers.AuthMiddleware(controllers.VendasHandler))
	http.HandleFunc("/marketing", controllers.AuthMiddleware(controllers.MarketingHandler))
	http.HandleFunc("/atendimento", controllers.AuthMiddleware(controllers.AtendimentoHandler))
	http.HandleFunc("/timeline", controllers.AuthMiddleware(controllers.TimelineHandler))
	http.HandleFunc("/integracoes", controllers.AuthMiddleware(controllers.IntegracoesHandler))
	http.HandleFunc("/configuracoes", controllers.AuthMiddleware(controllers.ConfiguracoesHandler))

	http.HandleFunc("/whatsapp", controllers.AuthMiddleware(controllers.WhatsAppHandler))
	http.HandleFunc("/whatsapp/criarInstancia", controllers.AuthMiddleware(controllers.CriarInstanciaHandler))
	http.HandleFunc("/whatsapp/qrcode", controllers.AuthMiddleware(controllers.QrCodeHandler))
}
