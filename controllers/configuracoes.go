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

//func PerfilAtualizacaoHandler() {}
