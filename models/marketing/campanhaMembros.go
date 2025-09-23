package marketing

import (
	"HareInteract.WebApp/db"
	"log"
)

type CampanhaMembros struct {
	Id             int    `json:"id" validate:"required"`
	Campanha       int    `json:"campanha" validate:"required"`
	Lead           int    `json:"lead"`
	Contato        int    `json:"contato"`
	StatusResposta string `json:"status_resposta" validate:"required"`
}

func CriarCampanhaMembros(search_path, statusResposta string, campanha, lead, contato int) {

	firstChar := search_path[0]
	if firstChar >= '0' && firstChar <= '9' {
		search_path = "C" + search_path
	}

	db := db.ConectaBD(search_path)

	incluirCampanhaMembros, err := db.Prepare("insert into campanhaMembros")
	if err != nil {
		log.Fatal(err.Error())
	}

	incluirCampanhaMembros.Exec(campanha, lead, contato, statusResposta)
	defer db.Close()
}

func DeletaCampanhaMembros(search_path, id string) {

	firstChar := search_path[0]
	if firstChar >= '0' && firstChar <= '9' {
		search_path = "C" + search_path
	}

	db := db.ConectaBD(search_path)

	deletarCampanhaMembros, err := db.Prepare("delete * from campanhaMembros where id=$1")
	if err != nil {
		log.Fatal(err.Error())
	}

	deletarCampanhaMembros.Exec(id)
	defer db.Close()
}

func ObterCampanhaMembros(search_path, id string) CampanhaMembros {

	firstChar := search_path[0]
	if firstChar >= '0' && firstChar <= '9' {
		search_path = "C" + search_path
	}

	db := db.ConectaBD(search_path)

	campanhaMembros, err := db.Query("select * from campanhaMembros where id=$1", id)
	if err != nil {
		log.Fatal(err.Error())
	}

	campanhaMembrosParaEditar := CampanhaMembros{}

	for campanhaMembros.Next() {
		var id, campanha, lead, contato int
		var statusResposta string

		err = campanhaMembros.Scan(&id, &campanha, &lead, &contato, &statusResposta)
		if err != nil {
			log.Fatal(err.Error())
		}

		campanhaMembrosParaEditar.Id = id
		campanhaMembrosParaEditar.Campanha = campanha
		campanhaMembrosParaEditar.Lead = lead
		campanhaMembrosParaEditar.Contato = contato
		campanhaMembrosParaEditar.StatusResposta = statusResposta
	}
	defer db.Close()
	return campanhaMembrosParaEditar
}

func AtualizarCampanhaMembros(search_path, statusResposta string, id, campanha, lead, contato int) {

	firstChar := search_path[0]
	if firstChar >= '0' && firstChar <= '9' {
		search_path = "C" + search_path
	}

	db := db.ConectaBD(search_path)

	campanhaMembrosAtualizada, err := db.Prepare("update campanhaMembros set campanha=$1, lead=$2, contato=$3, statusResposta=$4 where id=$5")
	if err != nil {
		log.Fatal(err.Error())
	}

	campanhaMembrosAtualizada.Exec(campanha, lead, contato, statusResposta, id)
	defer db.Close()
}
