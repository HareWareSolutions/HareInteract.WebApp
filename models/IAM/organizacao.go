package IAM

import (
	"database/sql"
	"net/http"
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

func CriarOrganizacao(nome string, responsavelId int, cpfcnpj string, pais string, cidade string, estado string, telefone string) error {
	db := db.ConectaBD("public")
	defer db.Close()

	dataCadastro := time.Now()

	cadastrarOrganizacao, err := db.Prepare("insert into organizacao(nome, responsavel, cpfcnpj, pais, cidade, estado, telefone, data_cadastro) values($1, $2, $3, $4, $5, $6, $7, $8)")

	if err != nil {
		return &apperr.Erro{
			Mensagem: "Falha ao preparar query de inserção!",
			Causa:    err,
		}
	}

	cadastrarOrganizacao.Close()

	_, err = cadastrarOrganizacao.Exec(nome, responsavelId, cpfcnpj, pais, cidade, estado, telefone, dataCadastro)

	if err != nil {
		return &apperr.Erro{
			Mensagem: "Falha ao executar a inserção da organização!",
			Causa:    err,
		}
	}

	return nil
}

func DeletaOrganizacao(id string) error {
	db := db.ConectaBD("public")
	defer db.Close()

	deletarOrganizacao, err := db.Prepare("delete from organizacao where id = $1")

	if err != nil {
		return &apperr.Erro{
			Mensagem: "Falha ao preparar query de remoção!",
			Causa:    err,
		}
	}

	_, err = deletarOrganizacao.Exec(id)
	if err != nil {
		return &apperr.Erro{
			Mensagem: "Falha ao remover organização!",
			Causa:    err,
		}
	}

	return nil
}

func ObterOrganizacao(id string) (*Organizacao, error) {
	db := db.ConectaBD("public")
	defer db.Close()

	var organizacao Organizacao

	row := db.QueryRow("select * from organizacao where id = $1", id)

	err := row.Scan(&organizacao.Id, &organizacao.Nome, &organizacao.ResponsavelId, &organizacao.Cpfcnpj,
		&organizacao.Pais, &organizacao.Cidade, &organizacao.Estado, &organizacao.Telefone, &organizacao.DataCadastro)

	if err != nil {

		if err == sql.ErrNoRows {
			return nil, &apperr.Erro{
				Mensagem: "Nenhum registro encontrado!",
				Status:   http.StatusNotFound,
			}
		} else {
			return nil, &apperr.Erro{
				Mensagem: "Falha ao consultar organização no banco de dados!",
				Causa:    err,
				Status:   http.StatusInternalServerError,
			}
		}

	}

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

	statement.Close()

	_, err = statement.Exec(organizacao.Nome, organizacao.Cpfcnpj, organizacao.Pais, organizacao.Cidade,
		organizacao.Estado, organizacao.Telefone, organizacao.Id)

	if err != nil {
		return &apperr.Erro{
			Mensagem: "Falha ao executar atualização da organização.",
			Causa:    err,
		}
	}

	return nil
}
