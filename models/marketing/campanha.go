package marketing

import (
	"HareInteract.WebApp/db"
	"time"
)

type Campanha struct {
	Id          int       `json:"id" validate:"required"`
	Nome        string    `json:"nome" validate:"required"`
	Tipo        string    `json:"tipo" validate:"required"`
	DataInicio  time.Time `json:"data_inicio" validate:"required"`
	DataFim     time.Time `json:"data_fim" validate:"required"`
	Orcamento   float64   `json:"orcamento" validate:"required"`
	Status      string    `json:"status" validate:"required"`
	Responsavel int       `json:"responsavel" validate:"required"`
}

func CriarCampanha(search_path, nome, tipo, status string, dataInicio, dataFim time.Time, orcamento float64, responsavel int) {

	firstChar := search_path[0]
	if firstChar >= '0' && firstChar <= '9' {
		search_path = "C" + search_path
	}

	db := db.ConectaBD(search_path)

	cadastrarCampanha, err := db.Prepare("insert into campanha(nome, tipo, dataInicio, dataFim, orcamento, status, responsavel) values($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		panic(err.Error())
	}

	cadastrarCampanha.Exec(nome, tipo, dataInicio, dataFim, orcamento, status, responsavel)
	defer db.Close()
}

func DeletaCampanha(search_path, id string) {

	firstChar := search_path[0]
	if firstChar >= '0' && firstChar <= '9' {
		search_path = "C" + search_path
	}

	db := db.ConectaBD(search_path)

	deletarCampanha, err := db.Prepare("delete from campanha where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deletarCampanha.Exec(id)
	defer db.Close()
}

func ObterCampanha(search_path, id string) Campanha {

	firstChar := search_path[0]
	if firstChar >= '0' && firstChar <= '9' {
		search_path = "C" + search_path
	}

	db := db.ConectaBD(search_path)

	campanha, err := db.Query("select * from campanha where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	campanhaParaEditar := Campanha{}

	for campanha.Next() {
		var id, responsavel int
		var nome, tipo, status string
		var dataInicio, dataFim time.Time
		var orcamento float64

		err = campanha.Scan(&id, &nome, &tipo, &dataInicio, &dataFim, &orcamento, &status, &responsavel)
		if err != nil {
			panic(err.Error())
		}

		campanhaParaEditar.Id = id
		campanhaParaEditar.Nome = nome
		campanhaParaEditar.Tipo = tipo
		campanhaParaEditar.DataInicio = dataInicio
		campanhaParaEditar.DataFim = dataFim
		campanhaParaEditar.Orcamento = orcamento
		campanhaParaEditar.Status = status
		campanhaParaEditar.Responsavel = responsavel
	}
	defer db.Close()
	return campanhaParaEditar
}

func AtualizarCampanha(search_path string, id, responsavel int, nome, tipo, status string, dataInicio, dataFim time.Time, orcamento float64) {

	firstChar := search_path[0]
	if firstChar >= '0' && firstChar <= '9' {
		search_path = "C" + search_path
	}

	db := db.ConectaBD(search_path)

	campanhaAtualizada, err := db.Prepare("update campanha set nome=$1, tipo=$2, dataInicio=$3, dataFim=$4, orcamento=$5, status=$6, responsavel=$7 where id=$8")
	if err != nil {
		panic(err.Error())
	}

	campanhaAtualizada.Exec(nome, tipo, dataInicio, dataFim, orcamento, status, responsavel)
	defer db.Close()
}
