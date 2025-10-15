package routes

import (
	"net/http"

	"HareInteract.WebApp/controllers"
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
	http.HandleFunc("/configuracoes/atualizar", controllers.AuthMiddleware(controllers.PerfilConfigAtualizarHandler))
	http.HandleFunc("/configuracoes/excluir/usuario", controllers.AuthMiddleware(controllers.PerfilConfigExcluirHandler))
	http.HandleFunc("/configuracoes/excluir/mensagem", controllers.AuthMiddleware(controllers.MensagemExcluirHandler))
	http.HandleFunc("/contatos", controllers.AuthMiddleware(controllers.ContatosHandler))
	http.HandleFunc("/empresas", controllers.AuthMiddleware(controllers.EmpresasHandler))
	http.HandleFunc("/atividades", controllers.AuthMiddleware(controllers.AtividadesHandler))
	http.HandleFunc("/leads", controllers.AuthMiddleware(controllers.LeadHandler))
	http.HandleFunc("/oportunidades", controllers.AuthMiddleware(controllers.OportunidadeHandler))
	http.HandleFunc("/tickets", controllers.AuthMiddleware(controllers.TicketHandler))
	http.HandleFunc("/campanhas", controllers.AuthMiddleware(controllers.CampanhaHandler))
	http.HandleFunc("/campanhaMembros", controllers.AuthMiddleware(controllers.CampanhaMembrosHandler))
	http.HandleFunc("/whatsapp", controllers.AuthMiddleware(controllers.WhatsAppHandler))
	http.HandleFunc("/whatsapp/criarInstancia", controllers.AuthMiddleware(controllers.CriarInstanciaHandler))
	http.HandleFunc("/whatsapp/qrcode", controllers.AuthMiddleware(controllers.QrCodeHandler))
}
