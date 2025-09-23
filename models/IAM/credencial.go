package IAM

import (
	"HareInteract.WebApp/db"
	"fmt"
)

type Credencial struct {
	Id          int    `json:"id" validate:"required"`
	Titulo      string `json:"titulo" validate:"required"`
	UrlAPI      string `json:"url_api"`
	TokenApi    string `json:"token_api"`
	InstanceApi string `json:"instance_api"`
	AssistantId string `json:"assistant_id"`
}

func CriarCredencial(search_path, titulo, urlApi, tokenApi, InstanceApi, AssistantId string) {

	firstChar := search_path[0]
	if firstChar >= '0' && firstChar <= '9' {
		search_path = "C" + search_path
	}

	db := db.ConectaBD(search_path)

	cadastrarCredencial, err := db.Prepare("insert into credencial (titulo, urlApi, tokenApi, InstanceApi, AssistantId) values($1, $2, $3, $4, $5)")
	if err != nil {
		panic(err.Error())
	}

	cadastrarCredencial.Exec(titulo, urlApi, tokenApi, InstanceApi, AssistantId)
	defer db.Close()
}

func DeletarCredencial(search_path, id string) {

	firstChar := search_path[0]
	if firstChar >= '0' && firstChar <= '9' {
		search_path = "C" + search_path
	}

	db := db.ConectaBD(search_path)

	deletarCredencial, err := db.Prepare("delete from credencial where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deletarCredencial.Exec(id)
	defer db.Close()
}

func ObterCredencial(search_path, id string) Credencial {

	firstChar := search_path[0]
	if firstChar >= '0' && firstChar <= '9' {
		search_path = "C" + search_path
	}

	db := db.ConectaBD(search_path)

	credencial, err := db.Query("select * from credencial where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	credencialParaEditar := Credencial{}

	for credencial.Next() {
		var id int
		var titulo, urlApi, tokenApi, InstanceApi, AssistantId string

		err = credencial.Scan(&id, &titulo, &urlApi, &tokenApi, &InstanceApi, &AssistantId)
		if err != nil {
			panic(err.Error())
		}

		credencialParaEditar.Id = id
		credencialParaEditar.Titulo = titulo
		credencialParaEditar.UrlAPI = urlApi
		credencialParaEditar.TokenApi = tokenApi
		credencialParaEditar.InstanceApi = InstanceApi
		credencialParaEditar.AssistantId = AssistantId
	}
	defer db.Close()
	return credencialParaEditar
}

func ObterCredencialPorTitulo(search_path, titulo string) Credencial {

	firstChar := search_path[0]
	if firstChar >= '0' && firstChar <= '9' {
		search_path = "C" + search_path
		fmt.Println("search_path ajustado: ", search_path)
	} else {
		fmt.Println("search_path usado: ", search_path)
	}

	db := db.ConectaBD(search_path)
	defer db.Close()

	credencial, err := db.Query("select * from credencial where titulo=$1", titulo)
	if err != nil {
		fmt.Println("Erro ao executar query:", err)
		return Credencial{}
	}

	fmt.Println("Query executada, resultado:", credencial)

	credencialEncontrada := Credencial{}

	if credencial.Next() {
		var id int
		var titulo, urlApi, tokenApi, InstanceApi, AssistantId string

		err = credencial.Scan(&id, &titulo, &urlApi, &tokenApi, &InstanceApi, &AssistantId)
		if err != nil {
			fmt.Println("Erro ao fazer scan:", err)
			return Credencial{}
		}

		credencialEncontrada.Id = id
		credencialEncontrada.Titulo = titulo
		credencialEncontrada.UrlAPI = urlApi
		credencialEncontrada.TokenApi = tokenApi
		credencialEncontrada.InstanceApi = InstanceApi
		credencialEncontrada.AssistantId = AssistantId
	} else {
		fmt.Println("Nenhuma credencial encontrada para o titulo:", titulo)
	}

	return credencialEncontrada
}

func AtualizarCredencial(search_path string, id int, titulo, urlApi, tokenApi, instanceApi, assistantId string) {

	firstChar := search_path[0]
	if firstChar >= '0' && firstChar <= '9' {
		search_path = "C" + search_path
	}

	db := db.ConectaBD(search_path)

	CredencialAtualizada, err := db.Prepare("update credencial set titulo=$1, url_api=$2, token_api=$3, instance_api=$4, assistant_id=$5 where id=$6")
	if err != nil {
		panic(err.Error())
	}

	CredencialAtualizada.Exec(titulo, urlApi, tokenApi, instanceApi, assistantId, id)
	defer db.Close()
}
