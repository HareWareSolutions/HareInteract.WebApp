package IAM

import (
	"log"
	"time"

	"HareInteract.WebApp/db"
)

type Mensagem struct {
	Id                          int
	Id_remetente                int
	Id_destinatario             int
	Mensagem_conteudo           string
	Data_envio                  time.Time
	Status                      bool
	Urgencia                    string
	Tipo                        string
	IdOrganizacaoConvite        int
	NivelAcessoUsuarioConvidado string
}

func CriarMensagem(id_remetente int, id_destinatario int, mensagem_conteudo string, urgencia string, tipo string) {
	db := db.ConectaBD("public")
	data_envio := time.Now()
	status := false

	inserirMensagem, err := db.Prepare("insert into mensagens(id_remetente, id_destinatario, conteudo_mensagem, data_envio, status, urgencia,  tipo) values($1, $2, $3, $4, $5)")
	if err != nil {
		panic(err.Error())
	}

	_, err = inserirMensagem.Exec(id_remetente, id_destinatario, mensagem_conteudo, data_envio, status, urgencia, tipo)

	if err != nil {
		log.Printf("Erro ao criar mensagem: ", err)
		return
	}

	defer db.Close()
}

func CriarConvite(id_remetente int, id_destinatario int, mensagem_conteudo string, urgencia string, tipo string, id_organizacao_convite int, nivel_acesso string) {
	db := db.ConectaBD("public")
	data_envio := time.Now()
	status := false

	inserirMensagem, err := db.Prepare("insert into mensagens(id_remetente, id_destinatario, conteudo_mensagem, data_envio, status, urgencia,  tipo, id_organizacao_convite, nivel_acesso) values($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		log.Println("Erro execução função CriarConvite :", err)
	}

	_, err = inserirMensagem.Exec(id_remetente, id_destinatario, mensagem_conteudo, data_envio, status, urgencia, tipo, id_organizacao_convite, nivel_acesso)

	if err != nil {
		log.Println("Erro ao inserir mensagem: ", err)
		return
	}

	defer db.Close()
}

func DeletarMensagem(id int) {
	db := db.ConectaBD("public")

	statement, err := db.Prepare("delete from mensagens where id = $1")
	if err != nil {
		panic(err.Error())
	}

	statement.Exec(id)
	defer db.Close()
}

func ObterMensagens(id_destinatario int) []Mensagem {
	db := db.ConectaBD("public")

	//idUser := r.Context().Value(userIdKey)

	statement, err := db.Prepare("SELECT * FROM mensagens where id_destinatario = $1")

	if err != nil {
		log.Println("Erro ao preparar statement (mensagem.ObterMensagens):", err)
		return nil
	}

	defer statement.Close()

	rows, err := statement.Query(id_destinatario)
	if err != nil {
		log.Println("Erro ao executar query (mensagem.ObterMensagens): ", err)
		return nil
	}

	defer rows.Close()

	mensagens := []Mensagem{}

	for rows.Next() {
		var m Mensagem

		err := rows.Scan(&m.Id, &m.Id_remetente, &m.Id_destinatario, &m.Mensagem_conteudo, &m.Data_envio, &m.Status, &m.Urgencia, &m.Tipo, &m.IdOrganizacaoConvite, &m.NivelAcessoUsuarioConvidado)

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

	err := row.Scan(&mensagem.Id, &mensagem.Id_remetente, &mensagem.Id_destinatario, &mensagem.Mensagem_conteudo, &mensagem.Data_envio, &mensagem.Status, &mensagem.Urgencia, &mensagem.Tipo)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	return mensagem
}
