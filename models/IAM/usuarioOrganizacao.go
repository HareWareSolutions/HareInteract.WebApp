package IAM

import (
	"HareInteract.WebApp/db"
	"time"
)

type UsuarioOrganizacao struct {
	Id           int       `json:"id" validate:"required"`
	Usuario      int       `json:"usuario" validate:"required"`
	Organizacao  int       `json:"organizacao" validate:"required"`
	NivelAcesso  string    `json:"nivelAcesso" validate:"required"`
	DataCadastro time.Time `json:"dataCadastro"`
}

func CriarUsuarioOrganizacao(usuario, organizacao int, nivelAcesso string) {
	db := db.ConectaBD("public")

	dataCadastro := time.Now()

	inserirUsuarioOrganizacao, err := db.Prepare("insert into usuario_organizacao(usuario, organizacao, nivel_acesso, dataCadastro) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	inserirUsuarioOrganizacao.Exec(usuario, organizacao, nivelAcesso, dataCadastro)
	defer db.Close()
}

func DeletaUsuarioOrganizacao(usuario int) {
	db := db.ConectaBD("public")

	deletarUsuarioOrganizacao, err := db.Prepare("delete from usuario_organizacao where usuario = $1")
	if err != nil {
		panic(err.Error())
	}

	deletarUsuarioOrganizacao.Exec(usuario)
	defer db.Close()
}

func ObterUsuarioOrganizacao(id int) UsuarioOrganizacao {
	db := db.ConectaBD("public")

	usuarioOrganizacaoParaEditar := UsuarioOrganizacao{}

	row := db.QueryRow("select id, usuario, organizacao, nivel_acesso from usuario_organizacao where id = $1", id)

	var id_db, usuario_db, organizacao_db int
	var nivelAcesso_db string

	err := row.Scan(&id_db, &usuario_db, &organizacao_db, &nivelAcesso_db)
	if err != nil {
		panic(err.Error())
	}

	usuarioOrganizacaoParaEditar.Id = id_db
	usuarioOrganizacaoParaEditar.Usuario = usuario_db
	usuarioOrganizacaoParaEditar.Organizacao = organizacao_db
	usuarioOrganizacaoParaEditar.NivelAcesso = nivelAcesso_db

	defer db.Close()
	return usuarioOrganizacaoParaEditar
}

func ObterUsuarioOrganizacaoPorUsuario(usuarioId int) UsuarioOrganizacao {
	db := db.ConectaBD("public")

	usuarioOrganizacaoRecuperado := UsuarioOrganizacao{}

	row := db.QueryRow("select id, usuario, organizacao, nivel_acesso from usuario_organizacao where usuario = $1", usuarioId)

	var id_db, usuario_db, organizacao_db int
	var nivelAcesso_db string

	err := row.Scan(&id_db, &usuario_db, &organizacao_db, &nivelAcesso_db)
	if err != nil {
		panic(err.Error())
	}

	usuarioOrganizacaoRecuperado.Id = id_db
	usuarioOrganizacaoRecuperado.Usuario = usuario_db
	usuarioOrganizacaoRecuperado.Organizacao = organizacao_db
	usuarioOrganizacaoRecuperado.NivelAcesso = nivelAcesso_db

	defer db.Close()
	return usuarioOrganizacaoRecuperado
}

func AtualizarUsuarioOrganizacao(id, usuario, organizacao int, nivelAcesso string) {
	db := db.ConectaBD("public")

	UsuarioOrganizacaoAtualizada, err := db.Prepare("update usuario_organizacao set usuario=$1, organizacao=$2, nivel_acesso=$3 where id = $4")
	if err != nil {
		panic(err.Error())
	}

	UsuarioOrganizacaoAtualizada.Exec(usuario, organizacao, nivelAcesso, id)
	defer db.Close()
}
