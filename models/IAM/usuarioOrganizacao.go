package IAM

import (
	"database/sql"
	"log"
	"time"

	"HareInteract.WebApp/db"
	"HareInteract.WebApp/models/apperr"
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

func CriarUsuarioOrganizacao(ID_Usuario, organizacao int, nivelAcesso string) error {
	db := db.ConectaBD("public")
	defer db.Close()

	dataCadastro := time.Now()

	cadastrarUsuarioOrganizacao, err := db.Prepare("insert into usuario_organizacao(usuario, organizacao, nivel_acesso, dataCadastro) values($1, $2, $3, $4)")
	if err != nil {
		return &apperr.Erro{
			Mensagem: "Falha ao preparar query de inserção!",
			Causa:    err,
		}
	}

	cadastrarUsuarioOrganizacao.Close()

	_, err = cadastrarUsuarioOrganizacao.Exec(ID_Usuario, organizacao, nivelAcesso, dataCadastro)

	if err != nil {
		return &apperr.Erro{
			Mensagem: "Falha ao executar a inserção da organização.",
			Causa:    err,
		}
	}

	return nil
}

func DeletaUsuarioOrganizacao(ID_Usuario int) error {
	db := db.ConectaBD("public")
	defer db.Close()

	deletarUsuarioOrganizacao, err := db.Prepare("delete from usuario_organizacao where usuario = $1")
	if err != nil {
		return &apperr.Erro{
			Mensagem: "Falha ao preparar query de remoção.",
			Causa:    err,
		}
	}

	deletarUsuarioOrganizacao.Close()

	_, err = deletarUsuarioOrganizacao.Exec(ID_Usuario)
	if err != nil {
		return &apperr.Erro{
			Mensagem: "Falha ao remover usuário organização.",
			Causa:    err,
		}
	}

	return nil
}

func ObterUsuariosOrganizacaoPorID_Organizacao(ID_Organizacao int) ([]UsuarioOrganizacao, error) {
	db := db.ConectaBD("public")
	defer db.Close()

	statement := "SELECT * FROM usuario_organizacao WHERE organizacao = $1 ORDER BY id"

	rows, err := db.Query(statement, ID_Organizacao)
	if err != nil {
		return nil, &apperr.Erro{
			Mensagem: "Falha ao executar query de busca de usuários da organização.",
			Causa:    err,
		}
	}

	defer rows.Close()

	usuariosOrganizacao := []UsuarioOrganizacao{}

	for rows.Next() {
		var usuarioOrganizacao UsuarioOrganizacao

		err := rows.Scan(&usuarioOrganizacao.Id, &usuarioOrganizacao.Usuario, &usuarioOrganizacao.Organizacao, &usuarioOrganizacao.NivelAcesso, &usuarioOrganizacao.DataCadastro)

		if err != nil {
			log.Printf("Erro escanear linha: ", err)
			continue
		}

		usuariosOrganizacao = append(usuariosOrganizacao, usuarioOrganizacao)
	}

	if err = rows.Err(); err != nil {
		return nil, &apperr.Erro{
			Mensagem: "Erro na leitura dos resultados da busca por usuários.",
			Causa:    err,
		}
	}

	return usuariosOrganizacao, nil
}

func ObterUsuariosOrgPublicoPorIdOrg(id int) ([]UsuarioOrganizacaoPublico, error) {

	Usuarios_Organizacao, err := ObterUsuariosOrganizacaoPorID_Organizacao(id)
	if err != nil {
		return nil, &apperr.Erro{
			Mensagem: "Falha ao obter usuários organização público.",
			Causa:    err,
		}
	}

	usuariosOrganizacaoPublico := []UsuarioOrganizacaoPublico{}

	for _, usuario := range Usuarios_Organizacao {
		UsuarioOrganizacaoPublico, _ := ConverterUsuarioOrgPublico(&usuario)
		usuariosOrganizacaoPublico = append(usuariosOrganizacaoPublico, *UsuarioOrganizacaoPublico)
	}

	return usuariosOrganizacaoPublico, nil
}

func ObterUsuarioOrgPublicoPorUsuario(ID_usuario int) (*UsuarioOrganizacaoPublico, error) {
	UsuarioOrganizacao, err := ObterUsuarioOrganizacaoPorUsuario(ID_usuario)
	if err != nil {
		return nil, &apperr.Erro{
			Mensagem: "Falha ao obter usuário organização público.",
			Causa:    err,
		}
	}

	UsuarioOrganizacaoPublico, err := ConverterUsuarioOrgPublico(UsuarioOrganizacao)
	if err != nil {
		return nil, &apperr.Erro{
			Mensagem: "Falha ao obter usuário organização público.",
			Causa:    err,
		}
	}

	return UsuarioOrganizacaoPublico, nil
}

func ObterUsuarioOrganizacao(ID_UsuarioOrganizacao int) (*UsuarioOrganizacao, error) {
	db := db.ConectaBD("public")
	defer db.Close()

	var usuarioOrganizacao UsuarioOrganizacao

	row := db.QueryRow("select id, usuario, organizacao, nivel_acesso from usuario_organizacao where id = $1", ID_UsuarioOrganizacao)

	err := row.Scan(
		&usuarioOrganizacao.Id,
		&usuarioOrganizacao.Usuario,
		&usuarioOrganizacao.Organizacao,
		&usuarioOrganizacao.NivelAcesso)

	if err != nil {
		if err == sql.ErrNoRows {
			return &UsuarioOrganizacao{}, &apperr.Erro{
				Mensagem: "Nenhum registro encontrado!",
			}
		} else {
			return &UsuarioOrganizacao{}, &apperr.Erro{
				Mensagem: "Falha ao consultar usuário organização no banco de dados",
				Causa:    err,
			}
		}
	}

	return &usuarioOrganizacao, nil

}

func ObterUsuarioOrganizacaoPorUsuario(ID_usuario int) (*UsuarioOrganizacao, error) {
	db := db.ConectaBD("public")
	defer db.Close()

	var usuarioOrganizacao UsuarioOrganizacao

	row := db.QueryRow("select id, usuario, organizacao, nivel_acesso from usuario_organizacao where usuario = $1", ID_usuario)

	err := row.Scan(
		&usuarioOrganizacao.Id,
		&usuarioOrganizacao.Usuario,
		&usuarioOrganizacao.Organizacao,
		&usuarioOrganizacao.NivelAcesso,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return &UsuarioOrganizacao{}, &apperr.Erro{
				Mensagem: "Nenhum registro encontrado!",
				Causa:    err,
			}
		} else {
			return &UsuarioOrganizacao{}, &apperr.Erro{
				Mensagem: "Falha ao consultar usuário organização no banco de dados!",
				Causa:    err,
			}
		}
	}

	return &usuarioOrganizacao, nil
}

func AtualizarUsuarioOrganizacao(UsuarioOrganizacao *UsuarioOrganizacao) error {
	db := db.ConectaBD("public")
	defer db.Close()

	statement, err := db.Prepare("update usuario_organizacao set usuario=$1, organizacao=$2, nivel_acesso=$3 where id = $4")
	if err != nil {
		return &apperr.Erro{
			Mensagem: "Falha ao preparar query de atualização.",
			Causa:    err,
		}
	}

	_, err = statement.Exec(UsuarioOrganizacao.Usuario, UsuarioOrganizacao.Organizacao, UsuarioOrganizacao.NivelAcesso, UsuarioOrganizacao.Id)

	if err != nil {
		return &apperr.Erro{
			Mensagem: "Falha ao executar atualização do usuário organização.",
			Causa:    err,
		}
	}

	return nil
}

func ConverterUsuarioOrgPublico(usuarioOrganizacao *UsuarioOrganizacao) (*UsuarioOrganizacaoPublico, error) {

	usuario, err := ObterUsuario(usuarioOrganizacao.Usuario)

	if err != nil {
		return &UsuarioOrganizacaoPublico{}, &apperr.Erro{
			Mensagem: "Erro ao converter usuário organização para público.",
			Causa:    err,
		}
	}

	UsuarioOrganizacaoPublico := UsuarioOrganizacaoPublico{}

	UsuarioOrganizacaoPublico.Id = usuarioOrganizacao.Id
	UsuarioOrganizacaoPublico.Usuario = usuarioOrganizacao.Usuario
	UsuarioOrganizacaoPublico.Organizacao = usuarioOrganizacao.Organizacao
	UsuarioOrganizacaoPublico.NivelAcesso = usuarioOrganizacao.NivelAcesso
	UsuarioOrganizacaoPublico.DataCadastro = usuarioOrganizacao.DataCadastro
	UsuarioOrganizacaoPublico.Nome = usuario.Nome
	UsuarioOrganizacaoPublico.Email = usuario.Email

	return &UsuarioOrganizacaoPublico, nil
}

func ValidarNivelAcesso(nivelUsuario string, nivelRequerido string) error {

	if nivelUsuario == nivelRequerido {
		return nil
	} else {
		return &apperr.Erro{
			Mensagem: "Você não possuí o nível de acesso necessário.",
		}
	}
}
