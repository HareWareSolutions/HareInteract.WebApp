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

type FullCalendarEvent struct {
	Title           string `json:"title"`
	Start           string `json:"start"`
	End             string `json:"end"`
	AllDay          bool   `json:"allDay"`
	BackgroundColor string `json:"backgroundColor"`
	ExtendedProps   struct {
		Contato_id  int    `json:"contato_id"`
		Usuario_id  int    `json:"usuario_id"`
		Confirmacao bool   `json:"confirmacao"`
		Observacao  string `json:"observacao"`
		Link        string `json:"link"`
	} `json:"extendedProps"`
}

func CriarAgendamento(agendamento *Agendamento, search_path string) error {
	firstChar := search_path[0]
	if firstChar >= '0' && firstChar <= '9' {
		search_path = "C" + search_path
	}

	db := db.ConectaBD(search_path)
	defer db.Close()

	statement, err := db.Prepare(`
		INSERT INTO agendamento (data, hora, contato_id, usuario_id, confirmacao, observacao, link)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`)

	if err != nil {
		return err
	}

	_, err = statement.Exec(agendamento.Data, agendamento.Hora, agendamento.Contato_id, agendamento.Usuario_id, agendamento.Confirmacao, agendamento.Observacao, agendamento.Link)
	if err != nil {
		return err
	}

	return nil
}

func DeletaAgendamento(search_path string, id int) error {
	firstChar := search_path[0]
	if firstChar >= '0' && firstChar <= '9' {
		search_path = "C" + search_path
	}

	db := db.ConectaBD(search_path)
	defer db.Close()

	statement, err := db.Prepare("DELETE FROM agendamento WHERE id = $1")
	if err != nil {
		return err
	}

	_, err = statement.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func ObterAgendamento(search_path string, id int) (*Agendamento, error) {
	firstChar := search_path[0]
	if firstChar >= '0' && firstChar <= '9' {
		search_path = "C" + search_path
	}

	db := db.ConectaBD(search_path)
	defer db.Close()

	result, err := db.Query(`
		SELECT id, data, hora, contato_id, usuario_id, confirmacao, observacao, link
		FROM agendamento WHERE id = $1
	`, id)

	if err != nil {
		return nil, err
	}

	agendamentoObtido := Agendamento{}

	for result.Next() {
		err = result.Scan(
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
			return nil, err
		}
	}

	return &agendamentoObtido, nil
}

func AtualizarAgendamento(agendamento *Agendamento, search_path string) error {
	firstChar := search_path[0]
	if firstChar >= '0' && firstChar <= '9' {
		search_path = "C" + search_path
	}

	db := db.ConectaBD(search_path)
	defer db.Close()

	statement, err := db.Prepare(`
		UPDATE agendamento
		SET data=$1, hora=$2, contato_id=$3, usuario_id=$4, confirmacao=$5, observacao=$6, link=$7
		WHERE id=$8
	`)
	if err != nil {
		return err
	}

	_, err = statement.Exec(agendamento.Data, agendamento.Hora, agendamento.Contato_id, agendamento.Usuario_id, agendamento.Confirmacao, agendamento.Observacao, agendamento.Link, agendamento.Id)
	if err != nil {
		return err
	}
	return nil

}

func ListarAgendamentos(search_path, mes string) ([]Agendamento, error) {
	firstChar := search_path[0]
	if firstChar >= '0' && firstChar <= '9' {
		search_path = "C" + search_path
	}

	db := db.ConectaBD(search_path)
	defer db.Close()

	result, err := db.Query(`
		SELECT id, data, hora, contato_id, usuario_id, confirmacao, observacao, link
		FROM agendamento
		WHERE TO_CHAR(data, 'YYYY-MM') = $1
		ORDER BY data, hora
	`, mes)
	if err != nil {
		return nil, err
	}

	var agendamentos []Agendamento

	for result.Next() {
		agendamento := Agendamento{}
		err = result.Scan(&agendamento.Id, &agendamento.Data, &agendamento.Hora, &agendamento.Contato_id, &agendamento.Usuario_id, &agendamento.Confirmacao, &agendamento.Observacao, &agendamento.Link)
		if err != nil {
			return nil, err
		}
		agendamentos = append(agendamentos, agendamento)
	}

	defer db.Close()
	return agendamentos, nil
}
