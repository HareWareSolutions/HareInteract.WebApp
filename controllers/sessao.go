package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"HareInteract.WebApp/models/IAM"
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

		if usuario_correto.Ativo == false {
			log.Println("O seu usuário não está mais ativo no sistema.")
			http.Error(w, "Sua conta não está ativa.", http.StatusUnauthorized)
			return
		}

		if usuario == usuario_correto.Username && senha == usuario_correto.Senha {
			usuarioOrganizacao := IAM.ObterUsuarioOrganizacaoPorUsuario(usuario_correto.Id)

			organizacaoDados := IAM.ObterOrganizacao(strconv.Itoa(usuarioOrganizacao.Organizacao))
			cpfCnpj := organizacaoDados.Cpfcnpj

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
