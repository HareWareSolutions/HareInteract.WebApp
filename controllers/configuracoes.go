package controllers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"HareInteract.WebApp/models/IAM"
)

// Handlers de Perfil

func PerfilConfigHandler(r *http.Request) (IAM.Usuario, error) {
	userID, ok := r.Context().Value(userIdKey).(int)

	if !ok {
		return IAM.Usuario{}, errors.New("Informações de sessão não encontradas.")
	}
	data := IAM.ObterUsuario(userID)

	return data, nil
}

func PerfilConfigAtualizarHandler(w http.ResponseWriter, r *http.Request) {

	user := IAM.Usuario{}

	user.Id = r.Context().Value(userIdKey).(int)
	user.Nome = r.FormValue("nome")
	user.Email = r.FormValue("email")
	user.Username = r.FormValue("username")
	user.Senha = r.FormValue("senha")

	IAM.AtualizarUsuario(user.Id, user.Nome, user.Email, user.Username, user.Senha, true)

	http.Redirect(w, r, "/configuracoes", http.StatusSeeOther)
}

func PerfilConfigExcluirHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(userIdKey).(int)

	user := IAM.ObterUsuario(userId)

	user.Ativo = false

	IAM.AtualizarUsuario(user.Id, user.Nome, user.Email, user.Username, user.Senha, user.Ativo)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

//Handlers de Mensagem

func MensagemCarregaHandler() []IAM.Mensagem {
	data := IAM.ObterMensagens()

	return data
}

func MensagemExcluirHandler(w http.ResponseWriter, r *http.Request) {

	idForm := r.FormValue("id")
	id, _ := strconv.Atoi(idForm)

	IAM.DeletarMensagem(id)

	http.Redirect(w, r, "/configuracoes", http.StatusSeeOther)
}

// Handlers de Organização

func OrganizacaoCarregaHandler(r *http.Request) IAM.Organizacao {
	userId := r.Context().Value(userIdKey).(int)

	userOrg := IAM.ObterUsuarioOrganizacaoPorUsuario(userId)

	org := IAM.ObterOrganizacao(strconv.Itoa(userOrg.Organizacao))

	return org
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

	if err != nil {
		log.Printf("Erro ao buscar usuário")
		http.Error(w, "Erro ao buscar usuário. Verifique se o Username está correto.", http.StatusForbidden)
		return
	}

	conteudo_mensagem := fmt.Sprintf("%s convidou você para sua organização!", &usuarioOrigem.Nome)

	var mensagem IAM.Mensagem

	mensagem.Id_remetente = usuarioOrigem.Id
	mensagem.Id_destinatario = usuarioDestino.Id
	mensagem.Mensagem_conteudo = conteudo_mensagem
	mensagem.Urgencia = "Alta"
	mensagem.Tipo = "Convite"
	IAM.CriarMensagem(mensagem.Id_remetente, mensagem.Id_destinatario, mensagem.Mensagem_conteudo, mensagem.Urgencia, mensagem.Tipo)

	fmt.Println("Mensagem enviada!")
}
