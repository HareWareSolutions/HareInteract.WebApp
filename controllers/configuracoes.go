package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"HareInteract.WebApp/models/IAM"
	"HareInteract.WebApp/models/apperr"
)

// Handlers de Perfil

func PerfilConfigHandler(r *http.Request) (*IAM.Usuario, error) {
	Id_usuario_valor := r.Context().Value(userIdKey)
	Id_usuario, ok := Id_usuario_valor.(int)

	if !ok {
		return &IAM.Usuario{}, &apperr.Erro{
			Mensagem: "ID de usuário não encontrado no contexto.",
			Status:   http.StatusUnauthorized,
		}
	}

	data, err := IAM.ObterUsuario(Id_usuario)
	if err != nil {
		return &IAM.Usuario{}, err
	}

	return data, nil
}

func PerfilConfigAtualizarHandler(w http.ResponseWriter, r *http.Request) {

	usuario := IAM.Usuario{}

	usuario.Id = r.Context().Value(userIdKey).(int)
	usuario.Nome = r.FormValue("nome")
	usuario.Email = r.FormValue("email")
	usuario.Username = r.FormValue("username")
	usuario.Senha = r.FormValue("senha")
	usuario.Ativo = true

	IAM.AtualizarUsuario(&usuario)

	http.Redirect(w, r, "/configuracoes", http.StatusSeeOther)
}

func PerfilConfigExcluirHandler(w http.ResponseWriter, r *http.Request) {
	Id_usuario := r.Context().Value(userIdKey).(int)

	usuario, err := IAM.ObterUsuario(Id_usuario)

	if err != nil {
		err = &apperr.Erro{
			Mensagem: "Falha ao obter usuário",
			Causa:    err,
		}

	}

	usuario.Ativo = false

	IAM.AtualizarUsuario(usuario)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

//Handlers de Mensagem

func MensagemCarregaHandler(r *http.Request) []IAM.Mensagem {
	userId := r.Context().Value(userIdKey).(int)

	data := IAM.ObterMensagens(userId)

	return data
}

func MensagemExcluirHandler(w http.ResponseWriter, r *http.Request) {

	idForm := r.FormValue("id")
	id, _ := strconv.Atoi(idForm)

	IAM.DeletarMensagem(id)

	http.Redirect(w, r, "/configuracoes", http.StatusSeeOther)
}

//func MensagemAceitaConvite(w http.ResponseWriter, r *http.Request) {

//}

// Handlers de Organização

func OrganizacaoCarregaHandler(r *http.Request) (*IAM.Organizacao, error) {
	Id_usuario_valor := r.Context().Value(userIdKey)
	Id_usuario, ok := Id_usuario_valor.(int)

	if !ok {
		return &IAM.Organizacao{}, &apperr.Erro{
			Mensagem: "ID de usuário não encontrado no contexto!",
			Status:   http.StatusUnauthorized,
		}
	}

	usuario_organizacao := IAM.ObterUsuarioOrganizacaoPorUsuario(Id_usuario)

	organizacao, err := IAM.ObterOrganizacao(strconv.Itoa(usuario_organizacao.Organizacao))

	if err != nil {
		return &IAM.Organizacao{}, &apperr.Erro{
			Mensagem: "Falha ao obter organização.",
			Causa:    err,
		}
	}

	return organizacao, nil
}

func OrganizacaoAtualizaHandler(w http.ResponseWriter, r *http.Request) {

	var organizacao *IAM.Organizacao

	organizacao.Id, _ = strconv.Atoi(r.FormValue("id"))
	organizacao.Nome = r.FormValue("nome")
	organizacao.Cpfcnpj = r.FormValue("documento")
	endereco := r.FormValue("endereco")
	listaEndereco := strings.Split(endereco, ",")
	organizacao.Cidade = strings.TrimSpace(listaEndereco[0])
	organizacao.Estado = strings.TrimSpace(listaEndereco[1])
	organizacao.Pais = strings.TrimSpace(listaEndereco[2])
	organizacao.Telefone = r.FormValue("telefone")

	IAM.AtualizarOrganizacao(organizacao)

	http.Redirect(w, r, "/configuracoes", http.StatusSeeOther)
}

// Handlers de Usuarios

func UsuariosCarregaHandler(r *http.Request) []IAM.UsuarioOrganizacaoPublico {
	userId := r.Context().Value(userIdKey).(int)
	userOrg := IAM.ObterUsuarioOrganizacaoPorUsuario(userId)

	idOrg := userOrg.Organizacao

	data := IAM.ObterUsuariosOrgPublicoPorIdOrg(idOrg)

	return data
}

func UsuarioAtualizaHandler(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")
	userId, _ := strconv.Atoi(id)

	userOrg := IAM.ObterUsuarioOrganizacao(userId)

	userOrg.NivelAcesso = r.FormValue("nivelAcesso")

	IAM.AtualizarUsuarioOrganizacao(userOrg.Id, userOrg.Usuario, userOrg.Organizacao, userOrg.NivelAcesso)

	http.Redirect(w, r, "/configuracoes", http.StatusSeeOther)
}

func UsuarioExcluirHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	idUserExcluir, _ := strconv.Atoi(id)

	IAM.DeletaUsuarioOrganizacao(idUserExcluir)

	http.Redirect(w, r, "/configuracoes", http.StatusSeeOther)
}

func UsuarioSairOrganizacao(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(userIdKey).(int)

	usuario := IAM.ObterUsuarioOrganizacaoPorUsuario(id)

	if IAM.ValidarNivelAcesso(usuario.NivelAcesso, "Proprietario") {
		log.Printf("Proprietário não deve sair da organização.")
		http.Error(w, "Proprietário não pode sair da organização. A organização deve ter outro proprietário designado primeiro ou ser excluída.", http.StatusForbidden)
		return
	}

	IAM.DeletaUsuarioOrganizacao(id)
	http.Redirect(w, r, "/configuracoes", http.StatusSeeOther)
}

func UsuarioConvidarOrganizacao(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("searchInput")
	id := r.Context().Value(userIdKey).(int)

	usuarioOrigem := IAM.ObterUsuarioOrgPublicoPorUsuario(id)

	usuarioDestino, err := IAM.ObterUsuarioPorUsername(username)

	nivelAcessoConvidado := r.FormValue("nivelAcesso")

	if err != nil {
		log.Printf("Erro ao buscar usuário: ", err)
		http.Error(w, "Erro ao buscar usuário. Verifique se o Username está correto.", http.StatusForbidden)
		return
	}

	//Validar se o usuário a convidar está participando de uma organização
	fmt.Println(usuarioDestino.Id)
	usuarioDestinoOrg := IAM.ObterUsuarioOrganizacaoPorUsuario(usuarioDestino.Id) //validação errada

	if usuarioDestinoOrg.Usuario == usuarioDestino.Id {
		http.Error(w, "Usuário convidado já pertence a uma organização! ", http.StatusForbidden)
		log.Printf("Usuário convidado já pertence a uma organização! ")
		return
	}

	conteudo_mensagem := fmt.Sprintf("%s convidou você para sua organização!", usuarioOrigem.Nome)

	var mensagem IAM.Mensagem

	mensagem.Id_remetente = usuarioOrigem.Id
	mensagem.Id_destinatario = usuarioDestino.Id
	mensagem.Mensagem_conteudo = conteudo_mensagem
	mensagem.Urgencia = "Alta"
	mensagem.Tipo = "Convite"
	mensagem.IdOrganizacaoConvite = usuarioOrigem.Organizacao
	mensagem.NivelAcessoUsuarioConvidado = nivelAcessoConvidado
	IAM.CriarConvite(mensagem.Id_remetente, mensagem.Id_destinatario, mensagem.Mensagem_conteudo, mensagem.Urgencia, mensagem.Tipo, mensagem.IdOrganizacaoConvite, mensagem.NivelAcessoUsuarioConvidado)

	fmt.Println("Convite enviado!")

	http.Redirect(w, r, "/configuracoes", http.StatusSeeOther)
}
