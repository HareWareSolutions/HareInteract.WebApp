package IAM

import (
	"HareInteract.WebApp/db"
	"time"
)

type Organizacao struct {
	Id            int       `json:"id" validate:"required"`
	Nome          string    `json:"nome" validate:"required,min=3,max=100"`
	ResponsavelId int       `json:"responsavelId" validate:"required"`
	Pais          string    `json:"pais"`
	Cidade        string    `json:"cidade"`
	Estado        string    `json:"estado"`
	Telefone      string    `json:"telefone"`
	DataCadastro  time.Time `json:"dataCadastro" validate:"required"`
}

func CriarOrganizacao(nome string, responsavelId int, pais string, cidade string, estado string, telefone string) {
	db := db.ConectaBD("public")

	dataCadastro := time.Now()

	cadastrarOrganizacao, err := db.Prepare("insert into organizacao(nome, responsavelId, pais, cidade, estado, telefone, dataCadastro) values($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		panic(err.Error())
	}

	cadastrarOrganizacao.Exec(nome, responsavelId, pais, cidade, estado, telefone, dataCadastro)
	defer db.Close()
}

func DeletaOrganizacao(id string) {
	db := db.ConectaBD("public")

	deletarOrganizacao, err := db.Prepare("delete from organizacao where id = $1")
	if err != nil {
		panic(err.Error())
	}

	deletarOrganizacao.Exec(id)
	defer db.Close()
}

func ObterOrganizacao(id string) Organizacao {
	db := db.ConectaBD("public")

	organizacao, err := db.Query("select * from organizacao where id = $1", id)
	if err != nil {
		panic(err.Error())
	}

	organizacaoParaEditar := Organizacao{}

	for organizacao.Next() {
		var id, responsavelId int
		var nome, pais, cidade, estado, telefone string

		err = organizacao.Scan(&id, &nome, &responsavelId, &pais, &cidade, &estado, &telefone)
		if err != nil {
			panic(err.Error())
		}

		organizacaoParaEditar.Id = id
		organizacaoParaEditar.Nome = nome
		organizacaoParaEditar.ResponsavelId = responsavelId
		organizacaoParaEditar.Pais = pais
		organizacaoParaEditar.Cidade = cidade
		organizacaoParaEditar.Estado = estado
		organizacaoParaEditar.Telefone = telefone
	}
	defer db.Close()
	return organizacaoParaEditar
}

func AtualizarOrganizacao(id int, nome string, responsavelId int, pais, cidade, estado, telefone string) {
	db := db.ConectaBD("public")

	OrganizacaoAtualizada, err := db.Prepare("update organizacao set nome=$1, responsavelId=$2, pais=$3, cidade=$4, estado=$5, telefone=$6 where id = $7")
	if err != nil {
		panic(err.Error())
	}

	OrganizacaoAtualizada.Exec(nome, responsavelId, pais, cidade, estado, telefone, id)
	defer db.Close()
}
