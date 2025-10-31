package integracao

import (
	"HareInteract.WebApp/db"
)

type Funcao struct {
	Id        int    `json:"id" validate:"required"`
	Titulo    string `json:"titulo" validate:"required"`
	Descricao string `json:"descricao" validate:"required"`
	Chamada   string `json:"chamada" validate:"required"`
}

func CriarFuncao(search_path, titulo, descricao, chamada string) {
	firstChar := search_path[0]
	if firstChar >= '0' && firstChar <= '9' {
		search_path = "C" + search_path
	}

	db := db.ConectaBD(search_path)

	inserirFuncao, err := db.Prepare("INSERT INTO funcao (titulo, descricao, chamada) VALUES ($1, $2, $3)")
	if err != nil {
		panic(err.Error())
	}

	_, err = inserirFuncao.Exec(titulo, descricao, chamada)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}

func ListarFuncoes(search_path string) []Funcao {
	firstChar := search_path[0]
	if firstChar >= '0' && firstChar <= '9' {
		search_path = "C" + search_path
	}

	db := db.ConectaBD(search_path)

	linhas, err := db.Query("SELECT id, titulo, descricao, chamada FROM funcao ORDER BY id ASC")
	if err != nil {
		panic(err.Error())
	}

	var funcoes []Funcao

	for linhas.Next() {
		var f Funcao
		err = linhas.Scan(&f.Id, &f.Titulo, &f.Descricao, &f.Chamada)
		if err != nil {
			panic(err.Error())
		}
		funcoes = append(funcoes, f)
	}

	defer db.Close()
	return funcoes
}

func ObterFuncao(search_path string, id int) Funcao {
	firstChar := search_path[0]
	if firstChar >= '0' && firstChar <= '9' {
		search_path = "C" + search_path
	}

	db := db.ConectaBD(search_path)

	resultado, err := db.Query("SELECT id, titulo, descricao, chamada FROM funcao WHERE id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	funcao := Funcao{}

	for resultado.Next() {
		err = resultado.Scan(&funcao.Id, &funcao.Titulo, &funcao.Descricao, &funcao.Chamada)
		if err != nil {
			panic(err.Error())
		}
	}

	defer db.Close()
	return funcao
}

func AtualizarFuncao(search_path string, id int, titulo, descricao, chamada string) {
	firstChar := search_path[0]
	if firstChar >= '0' && firstChar <= '9' {
		search_path = "C" + search_path
	}

	db := db.ConectaBD(search_path)

	atualizar, err := db.Prepare("UPDATE funcao SET titulo=$1, descricao=$2, chamada=$3 WHERE id=$4")
	if err != nil {
		panic(err.Error())
	}

	_, err = atualizar.Exec(titulo, descricao, chamada, id)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}

func DeletarFuncao(search_path string, id int) {
	firstChar := search_path[0]
	if firstChar >= '0' && firstChar <= '9' {
		search_path = "C" + search_path
	}

	db := db.ConectaBD(search_path)

	deletar, err := db.Prepare("DELETE FROM funcao WHERE id=$1")
	if err != nil {
		panic(err.Error())
	}

	_, err = deletar.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}
