package IAM

import (
	"fmt"
	"time"

	"HareInteract.WebApp/db"
)

type Organizacao struct {
	Id            int       `json:"id" validate:"required"`
	Nome          string    `json:"nome" validate:"required,min=3,max=100"`
	ResponsavelId int       `json:"responsavelId" validate:"required"`
	Cpfcnpj       string    `json:"cpfcnpj" validate:"required,min=3,max=100"`
	Pais          string    `json:"pais"`
	Cidade        string    `json:"cidade"`
	Estado        string    `json:"estado"`
	Telefone      string    `json:"telefone"`
	DataCadastro  time.Time `json:"dataCadastro" validate:"required"`
}

func CriarOrganizacao(nome string, responsavelId int, cpfcnpj string, pais string, cidade string, estado string, telefone string) {
	db := db.ConectaBD("public")
	defer db.Close()

	dataCadastro := time.Now()

	cadastrarOrganizacao, err := db.Prepare("insert into organizacao(nome, responsavel, cpfcnpj, pais, cidade, estado, telefone, data_cadastro) values($1, $2, $3, $4, $5, $6, $7, $8)")
	if err != nil {
		panic(err.Error())
	}

	cadastrarOrganizacao.Exec(nome, responsavelId, cpfcnpj, pais, cidade, estado, telefone, dataCadastro)
}

func DeletaOrganizacao(id string) {
	db := db.ConectaBD("public")
	defer db.Close()

	deletarOrganizacao, err := db.Prepare("delete from organizacao where id = $1")
	if err != nil {
		panic(err.Error())
	}

	deletarOrganizacao.Exec(id)
}

func ObterOrganizacao(id string) Organizacao {
	db := db.ConectaBD("public")
	defer db.Close()

	organizacao, err := db.Query("select * from organizacao where id = $1", id)
	if err != nil {
		panic(err.Error())
	}
	defer organizacao.Close()

	organizacaoParaEditar := Organizacao{}

	if organizacao.Next() {
		var id, responsavelId int
		var nome, cpfcnpj, pais, cidade, estado, telefone string
		var dataCadastro time.Time

		err = organizacao.Scan(&id, &nome, &responsavelId, &cpfcnpj, &pais, &cidade, &estado, &telefone, &dataCadastro)
		if err != nil {
			panic(err.Error())
		}

		organizacaoParaEditar.Id = id
		organizacaoParaEditar.Nome = nome
		organizacaoParaEditar.ResponsavelId = responsavelId
		organizacaoParaEditar.Cpfcnpj = cpfcnpj
		organizacaoParaEditar.Pais = pais
		organizacaoParaEditar.Cidade = cidade
		organizacaoParaEditar.Estado = estado
		organizacaoParaEditar.Telefone = telefone
		organizacaoParaEditar.DataCadastro = dataCadastro
	}
	return organizacaoParaEditar
}

func AtualizarOrganizacao(id int, nome string, cpfcnpj, pais, cidade, estado, telefone string) {
	db := db.ConectaBD("public")
	defer db.Close()

	OrganizacaoAtualizada, err := db.Prepare("update organizacao set nome=$1, cpfcnpj=$2, pais=$3, cidade=$4, estado=$5, telefone=$6 where id = $7")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Dados recebidos: (AtualizarOrganizacao)", nome, cpfcnpj, pais, cidade, estado, telefone, id)

	OrganizacaoAtualizada.Exec(nome, cpfcnpj, pais, cidade, estado, telefone, id)
}
