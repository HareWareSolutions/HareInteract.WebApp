package contas

import (
	"HareInteract.WebApp/db"
)

type Contato struct {
	Id          int    `json:"id" validate:"required"`
	Empresa     int    `json:"empresa"`
	Nome        string `json:"nome" validate:"required"`
	Sobrenome   string `json:"sobrenome" validate:"required"`
	Cargo       string `json:"cargo" validate:"required"`
	Email       string `json:"email"`
	Telefone    string `json:"telefone" validate:"required"`
	Responsavel int    `json:"responsavel" validate:"required"`
	Pausa       bool   `json:"pausa" validate:"required" default:"false"`
	Thread_id   string `json:"thread_id" validate:"required"`
}

func CriarContato(search_path string, empresa int, nome, sobrenome, cargo, email, telefone string, responsavel int, pausa bool, thread_id string) {

	firstChar := search_path[0]
	if firstChar >= '0' && firstChar <= '9' {
		search_path = "C" + search_path
	}

	db := db.ConectaBD(search_path)

	cadastrarContato, err := db.Prepare(`
		INSERT INTO contato 
		(empresa, nome, sobrenome, cargo, email, telefone, responsavel, pausa, thread_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`)
	if err != nil {
		panic(err.Error())
	}

	_, err = cadastrarContato.Exec(empresa, nome, sobrenome, cargo, email, telefone, responsavel, pausa, thread_id)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
}

func DeletaContato(search_path, id string) {

	firstChar := search_path[0]
	if firstChar >= '0' && firstChar <= '9' {
		search_path = "C" + search_path
	}

	db := db.ConectaBD(search_path)

	deletarContato, err := db.Prepare("DELETE FROM contato WHERE id = $1")
	if err != nil {
		panic(err.Error())
	}

	deletarContato.Exec(id)
	defer db.Close()
}

func ObterContato(search_path, id string) Contato {

	firstChar := search_path[0]
	if firstChar >= '0' && firstChar <= '9' {
		search_path = "C" + search_path
	}

	db := db.ConectaBD(search_path)

	contato, err := db.Query("SELECT id, empresa, nome, sobrenome, cargo, email, telefone, responsavel, pausa, thread_id FROM contato WHERE id = $1", id)
	if err != nil {
		panic(err.Error())
	}

	contatoParaEditar := Contato{}

	for contato.Next() {
		err = contato.Scan(
			&contatoParaEditar.Id,
			&contatoParaEditar.Empresa,
			&contatoParaEditar.Nome,
			&contatoParaEditar.Sobrenome,
			&contatoParaEditar.Cargo,
			&contatoParaEditar.Email,
			&contatoParaEditar.Telefone,
			&contatoParaEditar.Responsavel,
			&contatoParaEditar.Pausa,
			&contatoParaEditar.Thread_id,
		)
		if err != nil {
			panic(err.Error())
		}
	}
	defer db.Close()

	return contatoParaEditar
}

func AtualizarContato(search_path string, id, empresa, responsavel int, nome, sobrenome, cargo, email, telefone string, pausa bool, thread_id string) {

	firstChar := search_path[0]
	if firstChar >= '0' && firstChar <= '9' {
		search_path = "C" + search_path
	}

	db := db.ConectaBD(search_path)

	ContatoAtualizado, err := db.Prepare(`
		UPDATE contato 
		SET empresa=$1, nome=$2, sobrenome=$3, cargo=$4, email=$5, telefone=$6, responsavel=$7, pausa=$8, thread_id=$9 
		WHERE id = $10
	`)
	if err != nil {
		panic(err.Error())
	}

	_, err = ContatoAtualizado.Exec(empresa, nome, sobrenome, cargo, email, telefone, responsavel, pausa, thread_id, id)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}
