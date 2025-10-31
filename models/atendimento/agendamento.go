package atendimento

import (
	"HareInteract.WebApp/db"
)

type Agendamento struct {
	Id          int    `json:"id" validate:"required"`
	Data        string `json:"data" validate:"required"`
	Hora        string `json:"hora" validate:"required"`
	Contato_id  int    `json:"contato_id" validate:"required"`
	Usuario_id  int    `json:"usuario_id" validate:"required"`
	Confirmacao bool   `json:"confirmacao" validate:"required"`
	Observacao  string `json:"observacao" validate:"required"`
	Link        string `json:"link"`
}

func CriarAgendamento(search_path string, data, hora string, contato_id, usuario_id int, confirmacao bool, observacao, link string) {
	firstChar := search_path[0]
	if firstChar >= '0' && firstChar <= '9' {
		search_path = "C" + search_path
	}

	db := db.ConectaBD(search_path)

	cadastrarAgendamento, err := db.Prepare(`
		INSERT INTO agendamento (data, hora, contato_id, usuario_id, confirmacao, observacao, link)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`)
	if err != nil {
		panic(err.Error())
	}

	_, err = cadastrarAgendamento.Exec(data, hora, contato_id, usuario_id, confirmacao, observacao, link)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
}

func DeletaAgendamento(search_path string, id int) {
	firstChar := search_path[0]
	if firstChar >= '0' && firstChar <= '9' {
		search_path = "C" + search_path
	}

	db := db.ConectaBD(search_path)

	deletarAgendamento, err := db.Prepare("DELETE FROM agendamento WHERE id = $1")
	if err != nil {
		panic(err.Error())
	}

	deletarAgendamento.Exec(id)
	defer db.Close()
}

func ObterAgendamento(search_path string, id int) Agendamento {
	firstChar := search_path[0]
	if firstChar >= '0' && firstChar <= '9' {
		search_path = "C" + search_path
	}

	db := db.ConectaBD(search_path)

	agendamento, err := db.Query(`
		SELECT id, data, hora, contato_id, usuario_id, confirmacao, observacao, link
		FROM agendamento WHERE id = $1
	`, id)
	if err != nil {
		panic(err.Error())
	}

	agendamentoObtido := Agendamento{}

	for agendamento.Next() {
		err = agendamento.Scan(
			&agendamentoObtido.Id,
			&agendamentoObtido.Data,
			&agendamentoObtido.Hora,
			&agendamentoObtido.Contato_id,
			&agendamentoObtido.Usuario_id,
			&agendamentoObtido.Confirmacao,
			&agendamentoObtido.Observacao,
			&agendamentoObtido.Link,
		)
		if err != nil {
			panic(err.Error())
		}
	}
	defer db.Close()

	return agendamentoObtido
}

func AtualizarAgendamento(search_path string, id, contato_id, usuario_id int, data, hora, observacao, link string, confirmacao bool) {
	firstChar := search_path[0]
	if firstChar >= '0' && firstChar <= '9' {
		search_path = "C" + search_path
	}

	db := db.ConectaBD(search_path)

	agendamentoAtualizado, err := db.Prepare(`
		UPDATE agendamento
		SET data=$1, hora=$2, contato_id=$3, usuario_id=$4, confirmacao=$5, observacao=$6, link=$7
		WHERE id=$8
	`)
	if err != nil {
		panic(err.Error())
	}

	_, err = agendamentoAtualizado.Exec(data, hora, contato_id, usuario_id, confirmacao, observacao, link, id)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}

func ListarAgendamentos(search_path, mes string) []Agendamento {
	firstChar := search_path[0]
	if firstChar >= '0' && firstChar <= '9' {
		search_path = "C" + search_path
	}

	db := db.ConectaBD(search_path)

	rows, err := db.Query(`
		SELECT id, data, hora, contato_id, usuario_id, confirmacao, observacao, link
		FROM agendamento
		WHERE TO_CHAR(data, 'YYYY-MM') = $1
		ORDER BY data, hora
	`, mes)
	if err != nil {
		panic(err.Error())
	}

	var agendamentos []Agendamento

	for rows.Next() {
		a := Agendamento{}
		err = rows.Scan(&a.Id, &a.Data, &a.Hora, &a.Contato_id, &a.Usuario_id, &a.Confirmacao, &a.Observacao, &a.Link)
		if err != nil {
			panic(err.Error())
		}
		agendamentos = append(agendamentos, a)
	}

	defer db.Close()
	return agendamentos
}
