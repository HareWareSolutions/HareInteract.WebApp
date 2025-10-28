// controllers/middleware.go

package controllers

import (
	"context"
	"fmt"
	"net/http"
)

// Defina uma chave para o context. É uma boa prática usar um tipo customizado.
type contextKey string

const (
	orgCpfCnpjKey      contextKey = "orgCpfCnpj"
	userIdKey          contextKey = "userID"
	nivelAcessoUserKey contextKey = "accessLevel"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := GetSession(r)
		if err != nil {
			// Lidar com erro na sessão (ex: erro de decriptação)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		authenticated, ok := session.Values["authenticated"].(bool)
		if !ok || !authenticated {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		// Recuperar o CPF/CNPJ da sessão e injetá-lo no contexto da requisição
		orgCpfCnpj, ok := session.Values["orgCpfCnpj"].(string)
		if !ok {
			// Se o valor não estiver na sessão, redirecionar
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		userID, ok := session.Values["userId"].(int)
		if !ok {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		nivelAcessoUsuario, ok := session.Values["accessLevel"].(string)
		fmt.Println(nivelAcessoUsuario)
		if !ok {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		// Criar um novo contexto com o valor do CPF/CNPJ
		ctx := context.WithValue(r.Context(), orgCpfCnpjKey, orgCpfCnpj)
		ctx = context.WithValue(ctx, userIdKey, userID)
		ctx = context.WithValue(ctx, nivelAcessoUserKey, nivelAcessoUsuario)

		fmt.Println(ctx)
		// Chamar o próximo handler com o novo contexto
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
