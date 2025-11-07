package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"HareInteract.WebApp/models/IAM"
	"HareInteract.WebApp/models/apperr"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		usuario := r.FormValue("usuario")
		senha := r.FormValue("senha")

		usuario_correto, err := IAM.LoginUsuario(usuario)
		if err != nil {
			log.Println("Erro ao buscar o usuário: ", err)
			http.Error(w, "Ocorreu um erro, tente novamente.", http.StatusInternalServerError)
			return
		}

		if !usuario_correto.Ativo {
			log.Println("O seu usuário não está mais ativo no sistema.")
			http.Error(w, "Sua conta não está ativa.", http.StatusUnauthorized)
			return
		}

		if usuario == usuario_correto.Username && senha == usuario_correto.Senha {
			usuarioOrganizacao, err := IAM.ObterUsuarioOrganizacaoPorUsuario(usuario_correto.Id)
			if err != nil {
				statusCode := http.StatusInternalServerError

				if appErr, isCustom := err.(*apperr.Erro); isCustom {
					if appErr.Status != 0 {
						statusCode = appErr.Status
					}
				}

				w.WriteHeader(statusCode)

				templates.ExecuteTemplate(w, "erro.html", err)
				return
			}

			usuarioOrganizacao_IdOrganizacao := strconv.Itoa(usuarioOrganizacao.Organizacao)

			organizacao, err := IAM.ObterOrganizacao(usuarioOrganizacao_IdOrganizacao)

			if err != nil {
				statusCode := http.StatusInternalServerError

				if appErr, isCustom := err.(*apperr.Erro); isCustom {
					if appErr.Status != 0 {
						statusCode = appErr.Status
					}
				}

				w.WriteHeader(statusCode)

				templates.ExecuteTemplate(w, "erro.html", err)
				return
			}

			cpfCnpj := organizacao.Cpfcnpj

			session, err := GetSession(r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			session.Values["authenticated"] = true
			session.Values["orgCpfCnpj"] = cpfCnpj
			session.Values["userId"] = usuarioOrganizacao.Usuario
			session.Values["accessLevel"] = usuarioOrganizacao.NivelAcesso

			err = SaveSession(w, r, session)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			fmt.Println(session.Values["accessLevel"])

			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
			return
		} else {
			log.Println("Credenciais inválidas para o usuário:", usuario)
			http.Error(w, "Usuário ou senha incorretos.", http.StatusUnauthorized)
			return
		}
	}
	templates.ExecuteTemplate(w, "index.html", nil)
}
