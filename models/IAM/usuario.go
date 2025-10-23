package IAM

import (
	"database/sql"
	"fmt"
	"log"

	"HareInteract.WebApp/db"
)

type Usuario struct {
	Id       int    `json:"id" validate:"required"`
	Nome     string `json:"nome" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"usuario" validate:"required"`
	Senha    string `json:"senha" validate:"required,senha"`
	Ativo    bool   `json:"ativo"`
}

func CriarUsuario(nome, email, username, senha string) {
	db := db.ConectaBD("public")

	cadastrarUsuario, err := db.Prepare("insert into usuario(nome, email, username, senha, ativo) values($1, $2, $3, $4, $5)")
	if err != nil {
		panic(err.Error())
	}

	cadastrarUsuario.Exec(nome, email, username, senha, true)
	defer db.Close()
}

func DeletaUsuario(id string) {
	db := db.ConectaBD("public")

	deletarUsuario, err := db.Prepare("delete from usuario where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deletarUsuario.Exec(id)
	defer db.Close()
}

func ObterUsuarios() []Usuario {
	db := db.ConectaBD("public")

	statement := "select * from usuario"

	rows, err := db.Query(statement)
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	usuarios := []Usuario{}

	for rows.Next() {

		var u Usuario

		err := rows.Scan(&u.Id, &u.Nome, &u.Email, &u.Username, &u.Senha, &u.Ativo)

		if err != nil {
			log.Printf("Erro ao scananear linha: ", err)
			continue
		}

		usuarios = append(usuarios, u)

		if err = rows.Err(); err != nil {
			log.Fatal("Erro na iteração das linhas: %v", err)
		}
	}

	return usuarios

}

func ObterUsuario(id int) Usuario {
	db := db.ConectaBD("public")

	usuario, err := db.Query("select * from usuario where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	usuarioParaEditar := Usuario{}

	for usuario.Next() {
		var id int
		var nome, email, username, senha string
		var ativo bool

		err = usuario.Scan(&id, &nome, &email, &username, &senha, &ativo)
		if err != nil {
			panic(err.Error())
		}

		usuarioParaEditar.Id = id
		usuarioParaEditar.Nome = nome
		usuarioParaEditar.Email = email
		usuarioParaEditar.Username = username
		usuarioParaEditar.Senha = senha
		usuarioParaEditar.Ativo = ativo
	}
	defer db.Close()
	return usuarioParaEditar
}

func ObterUsuarioPorUsername(username string) (Usuario, error) {
	db := db.ConectaBD("public")

	row := db.QueryRow("SELECT id, username, senha, ativo FROM usuario WHERE username = $1", username)

	var usuario Usuario

	err := row.Scan(&usuario.Id, &usuario.Username, &usuario.Senha, &usuario.Ativo)
	if err != nil {
		if err == sql.ErrNoRows {
			return Usuario{}, fmt.Errorf("usuário '%s' não encontrado", username)
		}
		return Usuario{}, fmt.Errorf("erro ao buscar usuário: %v", err)
	}

	defer db.Close()
	return usuario, nil
}

func LoginUsuario(username string) (Usuario, error) {
	db := db.ConectaBD("public")

	row := db.QueryRow("SELECT id, username, senha, ativo FROM usuario WHERE username = $1", username)

	var usuario Usuario

	err := row.Scan(&usuario.Id, &usuario.Username, &usuario.Senha, &usuario.Ativo)
	if err != nil {
		if err == sql.ErrNoRows {
			return Usuario{}, fmt.Errorf("usuário '%s' não encontrado", username)
		}
		return Usuario{}, fmt.Errorf("erro ao buscar usuário: %v", err)
	}

	defer db.Close()
	return usuario, nil
}

func AtualizarUsuario(id int, nome, email, username, senha string, ativo bool) {
	db := db.ConectaBD("public")

	UsuarioAtualizado, err := db.Prepare("update usuario set nome=$1, email=$2, username=$3, senha=$4, ativo=$5 where id=$6")
	if err != nil {
		panic(err.Error())
	}

	UsuarioAtualizado.Exec(nome, email, username, senha, ativo, id)
	defer db.Close()
}
