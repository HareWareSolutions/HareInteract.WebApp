package atendimento

import (
	"HareInteract.WebApp/db"
	"time"
)

type Ticket struct {
	Id             int        `json:"id" validate:"required"`
	Assunto        string     `json:"assunto" validate:"required"`
	Descricao      string     `json:"descricao" validate:"required"`
	Contato        string     `json:"contato" validate:"required"`
	Status         string     `json:"status" validate:"required"`
	Prioridade     string     `json:"prioridade" validate:"required"`
	Responsavel    int        `json:"responsavel" validate:"required"`
	DataAbertura   time.Time  `json:"dataAbertura" validate:"required"`
	DataFechamento *time.Time `json:"dataFechamento"`
}

func CriarTicket(search_path, assunto, descricao, contato, status, prioridade string, responsavel int) {

	firstChar := search_path[0]
	if firstChar >= '0' && firstChar <= '9' {
		search_path = "C" + search_path
	}

	db := db.ConectaBD(search_path)

	dataAbertura := time.Now()

	cadastrarTicket, err := db.Prepare("insert into ticket (assunto, descricao, contato, status, prioridade, responsavel, dataAbertura, dataFechamento) values ($1, $2, $3, $4, $5, $6, $7, null)")
	if err != nil {
		panic(err.Error())
	}

	cadastrarTicket.Exec(assunto, descricao, contato, status, prioridade, responsavel, dataAbertura)
	defer db.Close()
}

func DeletaTicket(search_path, id string) {

	firstChar := search_path[0]
	if firstChar >= '0' && firstChar <= '9' {
		search_path = "C" + search_path
	}

	db := db.ConectaBD(search_path)

	deletarTicket, err := db.Prepare("delete from ticket where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deletarTicket.Exec(id)
	defer db.Close()
}

func ObterTicket(search_path, id string) Ticket {

	firstChar := search_path[0]
	if firstChar >= '0' && firstChar <= '9' {
		search_path = "C" + search_path
	}

	db := db.ConectaBD(search_path)

	ticket, err := db.Query("select id, assunto, descricao, contato, status, prioridade, responsavel, dataAbertura, dataFechamento from ticket where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	ticketParaEditar := Ticket{}

	for ticket.Next() {
		var id, responsavel int
		var assunto, descricao, contato, status, prioridade string
		var dataAbertura time.Time
		var dataFechamento *time.Time

		err = ticket.Scan(&id, &assunto, &descricao, &contato, &status, &prioridade, &responsavel, &dataAbertura, &dataFechamento)
		if err != nil {
			panic(err.Error())
		}

		ticketParaEditar.Id = id
		ticketParaEditar.Assunto = assunto
		ticketParaEditar.Descricao = descricao
		ticketParaEditar.Contato = contato
		ticketParaEditar.Status = status
		ticketParaEditar.Prioridade = prioridade
		ticketParaEditar.Responsavel = responsavel
		ticketParaEditar.DataAbertura = dataAbertura
		ticketParaEditar.DataFechamento = dataFechamento
	}
	defer db.Close()
	return ticketParaEditar
}

func AtualizarTicket(search_path, assunto, descricao, contato, status, prioridade string, responsavel int, dataFechamento *time.Time, id int) {

	firstChar := search_path[0]
	if firstChar >= '0' && firstChar <= '9' {
		search_path = "C" + search_path
	}

	db := db.ConectaBD(search_path)

	ticketAtualizado, err := db.Prepare("update ticket set assunto=$1, descricao=$2, contato=$3, status=$4, prioridade=$5, responsavel=$6, dataFechamento=$7 where id=$8")
	if err != nil {
		panic(err.Error())
	}

	ticketAtualizado.Exec(assunto, descricao, contato, status, prioridade, responsavel, dataFechamento, id)
	defer db.Close()
}
