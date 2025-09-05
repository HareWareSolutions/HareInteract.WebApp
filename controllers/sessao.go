package controllers

import (
	"HareInteract.WebApp/models/IAM"
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.ParseFiles("templates/index.html"))

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		usuario := r.FormValue("usuario")
		senha := r.FormValue("senha")

		usuario_correto, err := IAM.LoginUsuario(usuario)
		if err != nil {
			log.Println("Erro ao buscar o usuário: ", err)
			return
		}

		if usuario_correto.Ativo == false {
			log.Println("O seu usuário não está mais ativo no sistema.")
			return
		}
		if usuario == usuario_correto.Username && senha == usuario_correto.Senha {
			templates.ExecuteTemplate(w, "dashboard.html", nil)
			return
		} else {
			log.Println("Credenciais inválidas para o usuário:", usuario)

			http.Error(w, "Usuário ou senha incorretos.", http.StatusUnauthorized)
			return
		}
	}
	templates.ExecuteTemplate(w, "index.html", nil)
}
