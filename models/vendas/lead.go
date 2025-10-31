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
	Pausa       bool   `json:"pausa" validate:"required"`
	Thread_id   string `json:"thread_id" validate:"required"`
}

func CriarLead(search_path string, id, responsavel int, nome, email, telefone, empresa, origem, status string, pausa bool, thread_id string) {
	firstChar := search_path[0]
	if firstChar >= '0' && firstChar <= '9' {
		search_path = "C" + search_path
	}

	db := db.ConectaBD(search_path)

	cadastrarLead, err := db.Prepare(`
		INSERT INTO lead (nome, email, telefone, empresa, origem, status, responsavel, pausa, thread_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`)
	if err != nil {
		panic(err.Error())
	}

	_, err = cadastrarLead.Exec(nome, email, telefone, empresa, origem, status, responsavel, pausa, thread_id)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
}

func DeletaLead(search_path string, id int) {
	firstChar := search_path[0]
	if firstChar >= '0' && firstChar <= '9' {
		search_path = "C" + search_path
	}

	db := db.ConectaBD(search_path)

	deletarLead, err := db.Prepare("DELETE FROM lead WHERE id = $1")
	if err != nil {
		panic(err.Error())
	}

	deletarLead.Exec(id)
	defer db.Close()
}

func ObterLead(search_path string, id int) Lead {
	firstChar := search_path[0]
	if firstChar >= '0' && firstChar <= '9' {
		search_path = "C" + search_path
	}

	db := db.ConectaBD(search_path)

	lead, err := db.Query(`
		SELECT id, nome, email, telefone, empresa, origem, status, responsavel, pausa, thread_id 
		FROM lead WHERE id = $1
	`, id)
	if err != nil {
		panic(err.Error())
	}

	leadParaEditar := Lead{}

	for lead.Next() {
		err = lead.Scan(
			&leadParaEditar.Id,
			&leadParaEditar.Nome,
			&leadParaEditar.Email,
			&leadParaEditar.Telefone,
			&leadParaEditar.Empresa,
			&leadParaEditar.Origem,
			&leadParaEditar.Status,
			&leadParaEditar.Responsavel,
			&leadParaEditar.Pausa,
			&leadParaEditar.Thread_id,
		)
		if err != nil {
			panic(err.Error())
		}
	}
	defer db.Close()
	return leadParaEditar
}

func AtualizarLead(search_path string, id, responsavel int, nome, email, telefone, empresa, origem, status string, pausa bool, thread_id string) {
	firstChar := search_path[0]
	if firstChar >= '0' && firstChar <= '9' {
		search_path = "C" + search_path
	}

	db := db.ConectaBD(search_path)

	leadAtualizado, err := db.Prepare(`
		UPDATE lead 
		SET nome=$1, email=$2, telefone=$3, empresa=$4, origem=$5, status=$6, responsavel=$7, pausa=$8, thread_id=$9 
		WHERE id=$10
	`)
	if err != nil {
		panic(err.Error())
	}

	_, err = leadAtualizado.Exec(nome, email, telefone, empresa, origem, status, responsavel, pausa, thread_id, id)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}
