package IAM

import (
	"HareInteract.WebApp/db"
)

type Usuario struct {
	Id    int    `json:"id" validate:"required"`
	Nome  string `json:"nome" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	User  string `json:"usuario" validate:"required"`
	Senha string `json:"senha" validate:"required,senha"`
	Ativo bool   `json:"ativo"`
}

func CriarUsuario(nome, email, user, senha string) {
	db := db.ConectaBD("public")

	cadastrarUsuario, err := db.Prepare("insert into usuario(nome, email, user, senha, ativo) values($1, $2, $3, $4, $5)")
	if err != nil {
		panic(err.Error())
	}

	cadastrarUsuario.Exec(nome, email, user, senha, true)
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

func ObterUsuario(id string) Usuario {
	db := db.ConectaBD("public")

	usuario, err := db.Query("select * from usuario where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	usuarioParaEditar := Usuario{}

	for usuario.Next() {
		var id int
		var nome, email, user, senha string
		var ativo bool

		err = usuario.Scan(&id, &nome, &email, &user, &senha, &ativo)
		if err != nil {
			panic(err.Error())
		}

		usuarioParaEditar.Id = id
		usuarioParaEditar.Nome = nome
		usuarioParaEditar.Email = email
		usuarioParaEditar.User = user
		usuarioParaEditar.Senha = senha
		usuarioParaEditar.Ativo = ativo
	}
	defer db.Close()
	return usuarioParaEditar
}

func AtualizarUsuario(id int, nome, email, user, senha string, ativo bool) {
	db := db.ConectaBD("public")

	UsuarioAtualizado, err := db.Prepare("update usuario set nome=$1, email=$2, user=$3, senha=$4, ativo=$5 where id=$6")
	if err != nil {
		panic(err.Error())
	}

	UsuarioAtualizado.Exec(nome, email, user, senha, ativo, id)
	defer db.Close()
}
