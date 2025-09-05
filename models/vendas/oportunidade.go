package vendas

import (
	"HareInteract.WebApp/db"
	"time"
)

type Oportunidade struct {
	Id            int       `json:"id" validate:"required"`
	Titulo        string    `json:"titulo" validate:"required"`
	Empresa       int       `json:"empresa" validate:"required"`
	Contato       int       `json:"contato" validate:"required"`
	ValorEstimado float64   `json:"valor_estimado" validate:"required"`
	EtapaFunil    string    `json:"etapa_funil" validate:"required"`
	Probabilidade int       `json:"probabilidade"`
	Status        string    `json:"status" validate:"required"`
	Responsavel   int       `json:"responsavel" validate:"required"`
	DataCriacao   time.Time `json:"dataCriacao" validate:"required"`
}

func CriarOportunidade(search_path string, titulo, etapaFunil, status string, empresa, contato, probabilidade, responsavel int, valorEstimado float64) {
	db := db.ConectaBD(search_path)

	dataCriacao := time.Now()

	inserirOportunidade, err := db.Prepare("insert into oportunidade(titulo, empresa, contato, valorEstimado, etapaFunil, probabilidade, status, responsavel, dataCriacao) values($1, $2, $3, $4, $5, $6, $7, $8, $9)")
	if err != nil {
		panic(err.Error())
	}

	inserirOportunidade.Exec(titulo, empresa, contato, valorEstimado, etapaFunil, probabilidade, status, responsavel, dataCriacao)
	defer db.Close()
}

func DeletaOportunidade(search_path, id string) {
	db := db.ConectaBD(search_path)

	deletarOportunidade, err := db.Prepare("delete * from oportunidade where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deletarOportunidade.Exec(id)
	defer db.Close()
}

func ObterOportunidade(search_path, id string) Oportunidade {
	db := db.ConectaBD(search_path)

	oportunidade, err := db.Query("select * from oportunidade where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	oportunidadeParaEditar := Oportunidade{}

	for oportunidade.Next() {
		var id, empresa, contato, probabilidade, responsavel int
		var titulo, etapaFunil, status string
		var valorEstimado float64

		err = oportunidade.Scan(&id, &titulo, &empresa, &contato, &valorEstimado, &etapaFunil, &probabilidade, &status, &responsavel)
		if err != nil {
			panic(err.Error())
		}

		oportunidadeParaEditar.Id = id
		oportunidadeParaEditar.Titulo = titulo
		oportunidadeParaEditar.Empresa = empresa
		oportunidadeParaEditar.Contato = contato
		oportunidadeParaEditar.ValorEstimado = valorEstimado
		oportunidadeParaEditar.EtapaFunil = etapaFunil
		oportunidadeParaEditar.Probabilidade = probabilidade
		oportunidadeParaEditar.Status = status
		oportunidadeParaEditar.Responsavel = responsavel
	}
	defer db.Close()
	return oportunidadeParaEditar
}

func AtualizarOportunidade(search_path string, id, empresa, contato, probabilidade, responsavel int, titulo, etapaFunil, status string, valorEstimado float64) {
	db := db.ConectaBD(search_path)

	oportunidadeAtualizada, err := db.Prepare("update oportunidade set titulo=$1, empresa=$2, contato=$3, valorEstimado=$4, etapaFunil=$5, probabilidade=$6, status=$7, responsavel=$8 where id=$9")
	if err != nil {
		panic(err.Error())
	}

	oportunidadeAtualizada.Exec(titulo, empresa, contato, valorEstimado, etapaFunil, probabilidade, status, responsavel)
	defer db.Close()
}
