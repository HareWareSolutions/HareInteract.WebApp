package IAM

import (
	"log"
	"time"

	"HareInteract.WebApp/db"
	//"HareInteract.WebApp/models/IAM"
)

type UsuarioOrganizacao struct {
	Id           int       `json:"id" validate:"required"`
	Usuario      int       `json:"usuario" validate:"required"`
	Organizacao  int       `json:"organizacao" validate:"required"`
	NivelAcesso  string    `json:"nivelAcesso" validate:"required"`
	DataCadastro time.Time `json:"dataCadastro"`
}

type UsuarioOrganizacaoPublico struct {
	UsuarioOrganizacao
	Nome  string
	Email string
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

func ObterUsuariosOrganizacaoPorIdOrg(id int) []UsuarioOrganizacao {
	db := db.ConectaBD("public")

	rows, err := db.Query("SELECT * FROM usuario_organizacao WHERE organizacao = $1", id)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	usuarios := []UsuarioOrganizacao{}

	for rows.Next() {
		var u UsuarioOrganizacao

		err := rows.Scan(&u.Id, &u.Usuario, &u.Organizacao, &u.NivelAcesso, &u.DataCadastro)

		if err != nil {
			log.Printf("Erro escanear linha: ", err)
			continue
		}
		usuarios = append(usuarios, u)

		if err = rows.Err(); err != nil {
			log.Fatal("Erro iteração das linhas: %v", err)
		}
	}

	return usuarios
}

func ObterUsuariosOrgPublicoPorIdOrg(id int) []UsuarioOrganizacaoPublico {

	usuariosOrg := ObterUsuariosOrganizacaoPorIdOrg(id)
	usuariosOrgPub := []UsuarioOrganizacaoPublico{}

	for _, usuario := range usuariosOrg {
		usuarioOrgPub := ConverterUsuarioOrgPublico(usuario)
		usuariosOrgPub = append(usuariosOrgPub, usuarioOrgPub)
	}

	return usuariosOrgPub
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

	err := row.Scan(&usuarioOrganizacaoRecuperado.Id,
		&usuarioOrganizacaoRecuperado.Usuario,
		&usuarioOrganizacaoRecuperado.Organizacao,
		&usuarioOrganizacaoRecuperado.NivelAcesso)
	if err != nil {
		panic(err.Error())
	}

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

func ConverterUsuarioOrgPublico(u UsuarioOrganizacao) UsuarioOrganizacaoPublico {

	usuario := ObterUsuario(u.Usuario)

	usuarioOrgPub := UsuarioOrganizacaoPublico{}

	usuarioOrgPub.Id = u.Id
	usuarioOrgPub.Usuario = u.Usuario
	usuarioOrgPub.Organizacao = u.Organizacao
	usuarioOrgPub.NivelAcesso = u.NivelAcesso
	usuarioOrgPub.DataCadastro = u.DataCadastro
	usuarioOrgPub.Nome = usuario.Nome
	usuarioOrgPub.Email = usuario.Email

	return usuarioOrgPub
}
