package timeline

import (
	"HareInteract.WebApp/db"
	"time"
)

type Atividade struct {
	Id             int       `json:"id" validate:"required"`
	Assunto        string    `json:"assunto" validate:"required"`
	Tipo           string    `json:"tipo" validate:"required"`
	DataVencimento time.Time `json:"data_vencimento" validate:"required"`
	Status         string    `json:"status"`
	Descricao      string    `json:"descricao" validate:"required"`
	DataCriacao    time.Time `json:"data_criacao" validate:"required"`
}

func CriarAtividade(search_path, assunto, tipo, status, descricao string, dataVencimento time.Time) {

	firstChar := search_path[0]
	if firstChar >= '0' && firstChar <= '9' {
		search_path = "C" + search_path
	}

	db := db.ConectaBD(search_path)

	dataCriacao := time.Now()

	cadastrarAtividade, err := db.Prepare("insert into atividade (assunto, tipo, data_vencimento, status, descricao, data_criacao) values ($1, $2, $3, $4, $5, $6)")
	if err != nil {
		panic(err.Error())
	}

	cadastrarAtividade.Exec(assunto, tipo, dataVencimento, status, descricao, dataCriacao)
	defer db.Close()
}

func DeletaAtividade(search_path, id string) {

	firstChar := search_path[0]
	if firstChar >= '0' && firstChar <= '9' {
		search_path = "C" + search_path
	}

	db := db.ConectaBD(search_path)

	deletarAtividade, err := db.Prepare("delete from atividade where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deletarAtividade.Exec(id)
	defer db.Close()
}

func ObterAtividade(search_path, id string) Atividade {

	firstChar := search_path[0]
	if firstChar >= '0' && firstChar <= '9' {
		search_path = "C" + search_path
	}

	db := db.ConectaBD(search_path)

	atividade, err := db.Query("select id, assunto, tipo, data_vencimento, status, descricao, data_criacao from atividade where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	atividadeParaEditar := Atividade{}

	for atividade.Next() {
		var id int
		var assunto, tipo, status, descricao string
		var dataVencimento, dataCriacao time.Time

		err = atividade.Scan(&id, &assunto, &tipo, &dataVencimento, &status, &descricao, &dataCriacao)
		if err != nil {
			panic(err.Error())
		}

		atividadeParaEditar.Id = id
		atividadeParaEditar.Assunto = assunto
		atividadeParaEditar.Tipo = tipo
		atividadeParaEditar.DataVencimento = dataVencimento
		atividadeParaEditar.Status = status
		atividadeParaEditar.Descricao = descricao
		atividadeParaEditar.DataCriacao = dataCriacao
	}
	defer db.Close()
	return atividadeParaEditar
}

func AtualizarAtividade(search_path, assunto, tipo, status, descricao string, dataVencimento time.Time, id int) {

	firstChar := search_path[0]
	if firstChar >= '0' && firstChar <= '9' {
		search_path = "C" + search_path
	}

	db := db.ConectaBD(search_path)

	atividadeAtualizada, err := db.Prepare("update atividade set assunto=$1, tipo=$2, data_vencimento=$3, status=$4, descricao=$5 where id=$6")
	if err != nil {
		panic(err.Error())
	}

	atividadeAtualizada.Exec(assunto, tipo, dataVencimento, status, descricao, id)
	defer db.Close()
}
