package IAM

import (
	"database/sql"
	"fmt"
	"time"

	"HareInteract.WebApp/db"
	"HareInteract.WebApp/models/apperr"
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

func CriarOrganizacao(organizacao *Organizacao) error {
	db := db.ConectaBD("public")
	defer db.Close()

	dataCadastro := time.Now()

	statement, err := db.Prepare("insert into organizacao(nome, responsavel, cpfcnpj, pais, cidade, estado, telefone, data_cadastro) values($1, $2, $3, $4, $5, $6, $7, $8)")

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(organizacao.Nome, organizacao.ResponsavelId, organizacao.Cpfcnpj, organizacao.Pais, organizacao.Cidade, organizacao.Estado, organizacao.Telefone, dataCadastro)

	if err != nil {
		return err

	}

	return nil
}

func DeletaOrganizacao(id string) error {
	db := db.ConectaBD("public")
	defer db.Close()

	deletarOrganizacao, err := db.Prepare("delete from organizacao where id = $1")

	if err != nil {
		return err
	}

	defer deletarOrganizacao.Close()

	_, err = deletarOrganizacao.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func ObterOrganizacao(id string) (*Organizacao, error) {
	db := db.ConectaBD("public")
	defer db.Close()

	var organizacao Organizacao

	row := db.QueryRow("select id, nome, responsavel, cpfcnpj, pais, cidade, estado, telefone, data_cadastro from organizacao where id = $1", id)

	err := row.Scan(&organizacao.Id, &organizacao.Nome, &organizacao.ResponsavelId, &organizacao.Cpfcnpj,
		&organizacao.Pais, &organizacao.Cidade, &organizacao.Estado, &organizacao.Telefone, &organizacao.DataCadastro)

	if err != nil {

		if err == sql.ErrNoRows {
			return nil, err
		} else {
			return nil, err
		}

	}
	fmt.Println("ObterOrganizacao resultou em: ", organizacao)
	return &organizacao, nil

}

func AtualizarOrganizacao(organizacao *Organizacao) error {
	db := db.ConectaBD("public")
	defer db.Close()

	statement, err := db.Prepare("update organizacao set nome=$1, cpfcnpj=$2, pais=$3, cidade=$4, estado=$5, telefone=$6 where id = $7")

	if err != nil {
		return &apperr.Erro{
			Mensagem: "Falha ao preparar query de atualização!",
			Causa:    err,
		}
	}

	defer statement.Close()

	_, err = statement.Exec(organizacao.Nome, organizacao.Cpfcnpj, organizacao.Pais, organizacao.Cidade,
		organizacao.Estado, organizacao.Telefone, organizacao.Id)

	if err != nil {
		return &apperr.Erro{
			Mensagem: "Falha ao executar atualização da organização.",
			Causa:    err,
		}
	}
	fmt.Println("Organização.go AtualizarOrganizacao recebeu: ", organizacao)
	return nil
}
