package contas

import (
	"HareInteract.WebApp/db"
)

type Empresa struct {
	Id          int    `json:"id"`
	RazaoSocial string `json:"razaoSocial" validate:"required"`
	Cnpj        string `json:"cnpj"`
	Telefone    string `json:"telefone" validate:"required"`
	Site        string `json:"site"`
	Pais        string `json:"pais"`
	Cidade      string `json:"cidade"`
	Estado      string `json:"estado"`
	Cep         string `json:"cep"`
	Setor       string `json:"setor" validate:"required"`
}

func CriarEmpresa(search_path, razaoSocial, cnpj, telefone, site, pais, cidade, estado, cep, setor string) {

	firstChar := search_path[0]
	if firstChar >= '0' && firstChar <= '9' {
		search_path = "C" + search_path
	}

	db := db.ConectaBD(search_path)

	cadastrarEmpresa, err := db.Prepare("insert into empresa(razaoSocial, cnpj, telefone, site, pais, cidade, estado, cep, setor) values($1, $2, $3, $4, $5, $6, $7, $8, $9)")
	if err != nil {
		panic(err.Error())
	}

	cadastrarEmpresa.Exec(razaoSocial, cnpj, telefone, site, pais, cidade, estado, cep, setor)
	defer db.Close()
}

func DeletaEmpresa(search_path string, id int) {

	firstChar := search_path[0]
	if firstChar >= '0' && firstChar <= '9' {
		search_path = "C" + search_path
	}

	db := db.ConectaBD(search_path)

	deletarEmpresa, err := db.Prepare("delete from empresa where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deletarEmpresa.Exec(id)
	defer db.Close()
}

func ObterEmpresa(search_path string, id int) Empresa {

	firstChar := search_path[0]
	if firstChar >= '0' && firstChar <= '9' {
		search_path = "C" + search_path
	}

	db := db.ConectaBD(search_path)

	empresa, err := db.Query("select * from empresa where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	empresaParaEditar := Empresa{}

	for empresa.Next() {
		var id int
		var razaoSocial, cnpj, telefone, site, pais, cidade, estado, cep, setor string

		err = empresa.Scan(&id, &razaoSocial, &cnpj, &telefone, &site, &pais, &cidade, &estado, &cep, &setor)
		if err != nil {
			panic(err.Error())
		}

		empresaParaEditar.Id = id
		empresaParaEditar.RazaoSocial = razaoSocial
		empresaParaEditar.Cnpj = cnpj
		empresaParaEditar.Telefone = telefone
		empresaParaEditar.Site = site
		empresaParaEditar.Pais = pais
		empresaParaEditar.Cidade = cidade
		empresaParaEditar.Estado = estado
		empresaParaEditar.Cep = cep
		empresaParaEditar.Setor = setor
	}
	defer db.Close()
	return empresaParaEditar
}

func AtualizarEmpresa(search_path string, id int, razaoSocial, cnpj, telefone, site, pais, cidade, estado, cep, setor string) {

	firstChar := search_path[0]
	if firstChar >= '0' && firstChar <= '9' {
		search_path = "C" + search_path
	}

	db := db.ConectaBD(search_path)

	EmpresaAtualizada, err := db.Prepare("update empresa set razaoSocial=$1, cnpj=$2, telefone=$3, site=$4, pais=$5, cidade=$6, estado=$7, cep=$8, setor=$9 where id=$10")
	if err != nil {
		panic(err.Error())
	}

	EmpresaAtualizada.Exec(razaoSocial, cnpj, telefone, site, pais, cidade, estado, cep, setor, id)
	defer db.Close()
}
