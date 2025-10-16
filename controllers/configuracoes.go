package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"HareInteract.WebApp/models/IAM"
)

func PerfilConfigHandler(r *http.Request) (IAM.Usuario, error) {
	userID, ok := r.Context().Value(userIdKey).(int)

	if !ok {
		return IAM.Usuario{}, errors.New("Informações de sessão não encontradas.")
	}
	data := IAM.ObterUsuario(strconv.Itoa(userID))

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

	user := IAM.ObterUsuario(strconv.Itoa(userId))

	user.Ativo = false

	IAM.AtualizarUsuario(user.Id, user.Nome, user.Email, user.Username, user.Senha, user.Ativo)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

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

func OrganizacaoCarregaHandler(r *http.Request) IAM.Organizacao {
	userId := r.Context().Value(userIdKey).(int)

	userOrg := IAM.ObterUsuarioOrganizacaoPorUsuario(userId)

	org := IAM.ObterOrganizacao(strconv.Itoa(userOrg.Organizacao))

	return org
}
