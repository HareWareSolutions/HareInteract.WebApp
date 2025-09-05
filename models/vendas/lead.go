package vendas

import (
	"HareInteract.WebApp/db"
)

type Lead struct {
	Id          int    `json:"id" validate:"required"`
	Nome        string `json:"nome" validate:"required"`
	Email       string `json:"email"`
	Telefone    string `json:"telefone" validate:"required"`
	Empresa     string `json:"empresa"`
	Origem      string `json:"origem"`
	Status      string `json:"status" validate:"required"`
	Responsavel int    `json:"responsavel" validate:"required"`
}

func CriarLead(search_path string, id, responsavel int, nome, email, telefone, empresa, origem, status string) {
	db := db.ConectaBD(search_path)

	cadastrarLead, err := db.Prepare("insert into lead(nome, email, telefone, empresa, origem, status, responsavel) values($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		panic(err.Error())
	}

	cadastrarLead.Exec(nome, email, telefone, empresa, origem, status, responsavel)
	defer db.Close()
}

func DeletaLead(search_path string, id int) {
	db := db.ConectaBD(search_path)

	deletarLead, err := db.Prepare("delete from lead where id = $1")
	if err != nil {
		panic(err.Error())
	}

	deletarLead.Exec(id)
	defer db.Close()
}

func ObterLead(serch_path string, id int) Lead {
	db := db.ConectaBD(serch_path)

	lead, err := db.Query("select * from lead where id = $1", id)
	if err != nil {
		panic(err.Error())
	}

	leadParaEditar := Lead{}

	for lead.Next() {
		var id, responsavel int
		var nome, email, telefone, empresa, origem, status string

		err = lead.Scan(&id, &nome, &email, &telefone, &empresa, &origem, &status, &responsavel)
		if err != nil {
			panic(err.Error())
		}

		leadParaEditar.Id = id
		leadParaEditar.Nome = nome
		leadParaEditar.Email = email
		leadParaEditar.Telefone = telefone
		leadParaEditar.Empresa = empresa
		leadParaEditar.Origem = origem
		leadParaEditar.Status = status
		leadParaEditar.Responsavel = responsavel
	}
	defer db.Close()
	return leadParaEditar
}

func AtualizarLead(search_path string, id, responsavel int, nome, email, telefone, empresa, origem, status string) {
	db := db.ConectaBD(search_path)

	leadAtualizado, err := db.Prepare("update lead set nome=$1, email=$2, telefone=$3, empresa=$4, origem=$5, status=$6, responsavel=$7 where id=$8")
	if err != nil {
		panic(err.Error())
	}

	leadAtualizado.Exec(nome, email, telefone, empresa, origem, status, responsavel, id)
	defer db.Close()
}
