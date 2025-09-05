package contas

import (
	"HareInteract.WebApp/db"
)

type Contato struct {
	Id          int    `json:"id" validate:"required"`
	Empresa     int    `json:"empresa"`
	Nome        string `json:"nome" validate:"required"`
	Sobrenome   string `json:"sobrenome" validate:"required"`
	Cargo       string `json:"cargo" validate:"required"`
	Email       string `json:"email"`
	Telefone    string `json:"telefone" validate:"required"`
	Responsavel int    `json:"responsavel" validate:"required"`
}

func CriarContato(search_path string, empresa int, nome, sobrenome, cargo, email, telefone, responsavel int) {
	db := db.ConectaBD(search_path)

	cadastrarContato, err := db.Prepare("insert into contato (empresa, nome, sobrenome, cargo, email, telefone, responsavel) values($1,$2,$3,$4,$5,$6,$7)")
	if err != nil {
		panic(err.Error())
	}

	cadastrarContato.Exec(empresa, nome, sobrenome, cargo, email, telefone, responsavel)
	defer db.Close()
}

func DeletaContato(search_path, id string) {
	db := db.ConectaBD(search_path)

	deletarContato, err := db.Prepare("delete from contato where id = $1")
	if err != nil {
		panic(err.Error())
	}

	deletarContato.Exec(id)
	defer db.Close()
}

func ObterContato(search_path, id string) Contato {
	db := db.ConectaBD(search_path)

	contato, err := db.Query("select * from contato where id = $1", id)
	if err != nil {
		panic(err.Error())
	}

	contatoParaEditar := Contato{}

	for contato.Next() {
		var id, empresa, responsavel int
		var nome, sobrenome, cargo, email, telefone string

		err = contato.Scan(&id, &empresa, &nome, &sobrenome, &cargo, &email, &telefone, &responsavel)
		if err != nil {
			panic(err.Error())
		}

		contatoParaEditar.Id = id
		contatoParaEditar.Empresa = empresa
		contatoParaEditar.Nome = nome
		contatoParaEditar.Sobrenome = sobrenome
		contatoParaEditar.Cargo = cargo
		contatoParaEditar.Email = email
		contatoParaEditar.Telefone = telefone
		contatoParaEditar.Responsavel = responsavel
	}
	defer db.Close()
	return contatoParaEditar
}

func AtualizarContato(search_path string, id, empresa, responsavel int, nome, sobrenome, cargo, email, telefone string) {
	db := db.ConectaBD(search_path)

	ContatoAtualizado, err := db.Prepare("update contato set empresa=$1, nome=$2, sobrenome=$3, cargo=$4, email=$5, telefone=$6, responsavel=$7 where id = $8")
	if err != nil {
		panic(err.Error())
	}

	ContatoAtualizado.Exec(empresa, nome, sobrenome, cargo, email, telefone, responsavel, id)
	defer db.Close()
}
