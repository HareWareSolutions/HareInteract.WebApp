package controllers

import (
	"html/template"
	"io"
	"net/http"
	"strings"
)

var templates = template.Must(template.ParseFiles("templates/index.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		err := templates.ExecuteTemplate(w, "index.html", nil)
		if err != nil {
			http.Error(w, "Erro ao renderizar template", http.StatusInternalServerError)
		}
		return
	}

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Erro ao processar o formulário", http.StatusBadRequest)
			return
		}

		username := r.FormValue("usuario")
		password := r.FormValue("senha")

		payload := strings.NewReader("grant_type=password&username=" + username + "&password=" + password + "&scope=&client_id=string&client_secret=string")

		req, err := http.NewRequest("POST", "https://backend.hareblast.com.br/token", payload)
		if err != nil {
			http.Error(w, "Erro ao criar requisição externa", http.StatusInternalServerError)
			return
		}

		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("accept", "application/json")

		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			http.Error(w, "Erro ao enviar requisição externa", http.StatusInternalServerError)
			return
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			http.Error(w, "Erro ao ler resposta da API", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
		return
	}

	http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
}
