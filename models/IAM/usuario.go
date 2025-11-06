package IAM

import (
	"database/sql"
	"log"
	"net/http"

	"HareInteract.WebApp/db"
	"HareInteract.WebApp/models/apperr"
)

type Usuario struct {
	Id       int    `json:"id" validate:"required"`
	Nome     string `json:"nome" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"usuario" validate:"required"`
	Senha    string `json:"senha" validate:"required,senha"`
	Ativo    bool   `json:"ativo"`
}

func CriarUsuario(nome, email, username, senha string) error {
	db := db.ConectaBD("public")
	defer db.Close()

	cadastrarUsuario, err := db.Prepare("insert into usuario(nome, email, username, senha, ativo) values($1, $2, $3, $4, $5)")
	if err != nil {
		return &apperr.Erro{
			Mensagem: "Falha ao preparar query de inserção!",
			Causa:    err,
		}
	}

	cadastrarUsuario.Close()

	_, err = cadastrarUsuario.Exec(nome, email, username, senha, true)

	if err != nil {
		return &apperr.Erro{
			Mensagem: "Erro ao executar a inserção do usuário!",
			Causa:    err,
		}
	}

	return nil //Sucesso bb
}

func DeletaUsuario(id string) error {
	db := db.ConectaBD("public")
	defer db.Close()

	deletarUsuario, err := db.Prepare("delete from usuario where id=$1")
	if err != nil {
		return &apperr.Erro{
			Mensagem: "Erro ao preparar query de remoção!",
			Causa:    err,
		}
	}

	_, err = deletarUsuario.Exec(id)
	if err != nil {
		return &apperr.Erro{
			Mensagem: "Erro ao remover usuário!",
			Causa:    err,
		}
	}

	return nil //Sucesso bb
}

func ObterUsuarios() ([]Usuario, error) {
	db := db.ConectaBD("public")
	defer db.Close()

	statement := "select id, nome, username, senha, ativo from usuario"

	rows, err := db.Query(statement)
	if err != nil {
		return nil, &apperr.Erro{
			Mensagem: "Falha ao executar a query de busca de usuários",
			Causa:    err,
		}
	}

	defer rows.Close()

	usuarios := []Usuario{}

	for rows.Next() {

		var usuario Usuario

		err := rows.Scan(&usuario.Id, &usuario.Nome, &usuario.Email, &usuario.Username, &usuario.Senha, &usuario.Ativo)

		if err != nil {
			log.Printf("Erro ao scananear linha: ", err)
			continue
		}

		usuarios = append(usuarios, usuario)

	}

	if err = rows.Err(); err != nil {
		return nil, &apperr.Erro{
			Mensagem: "Erro na leitura dos resultados da busca por usuários",
			Causa:    err,
		}
	}

	return usuarios, nil

}

func ObterUsuario(id int) (*Usuario, error) {
	db := db.ConectaBD("public")
	defer db.Close()

	var usuario Usuario

	row := db.QueryRow("select id, nome, email, username, senha, ativo from usuario where id=$1", id)

	err := row.Scan(&usuario.Id, &usuario.Nome, &usuario.Email, &usuario.Username, &usuario.Senha, &usuario.Ativo)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &apperr.Erro{
				Mensagem: "Nenhum registro encontrado!",
				Status:   http.StatusNotFound,
			}
		} else {
			return nil, &apperr.Erro{
				Mensagem: "Falha ao consultar usuário no banco de dados!",
				Causa:    err,
				Status:   http.StatusInternalServerError,
			}
		}
	}

	return &usuario, nil
}

func ObterUsuarioPorUsername(username string) (*Usuario, error) {
	db := db.ConectaBD("public")
	defer db.Close()

	var usuario Usuario

	row := db.QueryRow("SELECT id, username, senha, ativo FROM usuario WHERE username = $1", username)

	err := row.Scan(&usuario.Id, &usuario.Username, &usuario.Senha, &usuario.Ativo)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &apperr.Erro{
				Mensagem: "Nenhum registro encontrado!",
			}
		} else {
			return nil, &apperr.Erro{
				Mensagem: "Falha ao consultar usuário no banco de dados!",
				Causa:    err,
			}
		}

	}

	return &usuario, nil
}

func LoginUsuario(username string) (*Usuario, error) {
	db := db.ConectaBD("public")
	defer db.Close()

	row := db.QueryRow("SELECT id, username, senha, ativo FROM usuario WHERE username = $1", username)

	var usuario Usuario

	err := row.Scan(&usuario.Id, &usuario.Username, &usuario.Senha, &usuario.Ativo)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &apperr.Erro{
				Mensagem: "Usuário não cadastrado!",
			}
		} else {
			return nil, &apperr.Erro{
				Mensagem: "Falha ao consultar usuário no banco de dados!",
				Causa:    err,
			}
		}

	}

	return &usuario, nil
}

func AtualizarUsuario(usuario *Usuario) error {
	db := db.ConectaBD("public")
	defer db.Close()

	statement, err := db.Prepare("update usuario set nome=$1, email=$2, username=$3, senha=$4, ativo=$5 where id=$6")
	if err != nil {
		return &apperr.Erro{
			Mensagem: "Falha ao preparar query de atualização!",
			Causa:    err,
		}
	}

	statement.Close()

	_, err = statement.Exec(usuario.Nome, usuario.Email, usuario.Username, usuario.Senha, usuario.Ativo, usuario.Id)

	if err != nil {
		return &apperr.Erro{
			Mensagem: "Falha ao executar atualização do usuário.",
			Causa:    err,
		}
	}

	return nil //Suceso bb
}
