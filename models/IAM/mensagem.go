package IAM

import (
	"log"
	"time"

	"HareInteract.WebApp/db"
)

type Mensagem struct {
	Id                int
	Id_remetente      int
	Id_destinatario   int
	Mensagem_conteudo string
	Data_envio        time.Time
	Status            bool
	Urgencia          string
}

func CriarMensagem(id_remetente int, id_destinatario int, mensagem_conteudo string, status bool, urgencia string) {
	db := db.ConectaBD("public")

	data_envio := time.Now()

	inserirMensagem, err := db.Prepare("insert into mensagens(id_remetente, id_destinatario, conteudo_mensagem, data_envio, status, urgencia) values($1, $2, $3, $4, $5, $6)")
	if err != nil {
		panic(err.Error())
	}

	inserirMensagem.Exec(id_remetente, id_destinatario, mensagem_conteudo, data_envio, status, urgencia)
	defer db.Close()
}

func DeletaMensagem(id int) {
	db := db.ConectaBD("public")

	statement, err := db.Prepare("delete from mensagems where id = $1")
	if err != nil {
		panic(err.Error())
	}

	statement.Exec(id)
	defer db.Close()
}

func ObterMensagens() []Mensagem {
	db := db.ConectaBD("public")

	statement := "SELECT * FROM mensagens"

	rows, err := db.Query(statement)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	mensagens := []Mensagem{}

	for rows.Next() {
		var m Mensagem

		err := rows.Scan(&m.Id, &m.Id_remetente, &m.Id_destinatario, &m.Mensagem_conteudo, &m.Data_envio, &m.Status, &m.Urgencia)

		if err != nil {
			log.Printf("Erro escanear linha: ", err)
			continue
		}

		mensagens = append(mensagens, m)

		if err = rows.Err(); err != nil {
			log.Fatal("Erro iteração das linhas: %v", err)
		}
	}

	return mensagens
}

func ObterMensagem(id int) Mensagem {
	db := db.ConectaBD("public")

	mensagem := Mensagem{}

	row := db.QueryRow("select * from mensagems where id = $1", id)

	err := row.Scan(&mensagem.Id, &mensagem.Id_remetente, &mensagem.Id_destinatario, &mensagem.Mensagem_conteudo, &mensagem.Data_envio, &mensagem.Status, &mensagem.Urgencia)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	return mensagem
}
