package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func ConectaBD(search_path string) *sql.DB {
	if search_path == "" {
		log.Fatal("search_path não pode ser vazio.")
	}

	firstChar := search_path[0]
	if firstChar >= '0' && firstChar <= '9' {
		search_path = "C" + search_path
	}

	conexao := fmt.Sprintf("user=postgres dbname=HareInteractCRM password=12345 host=localhost sslmode=disable search_path=%s,public", search_path)

	db, err := sql.Open("postgres", conexao)
	if err != nil {
		log.Fatalf("Erro ao abrir a conexão com o banco de dados: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Erro ao pingar o banco de dados: %v", err)
	}

	createSchemaSQL := fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s;", search_path)
	_, err = db.Exec(createSchemaSQL)
	if err != nil {
		log.Fatalf("Erro ao criar o schema %s: %v", search_path, err)
	}

	InicializaTabelas(db, search_path)

	return db
}

func InicializaTabelas(db *sql.DB, search_path string) {
	log.Println("Verificando e inicializando tabelas do banco de dados...")

	// Prefixo do schema para cada nome de tabela
	schemaPrefix := fmt.Sprintf("%s.", search_path)

	// Lista de instruções SQL para criar cada tabela individualmente
	createTableStatements := []string{
		// Tabela de Usuários
		fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS %susuario (
           id SERIAL PRIMARY KEY,
           nome VARCHAR(255) NOT NULL,
           email VARCHAR(255) UNIQUE NOT NULL,
           username VARCHAR(255) UNIQUE NOT NULL,
           senha VARCHAR(255) NOT NULL,
           ativo BOOLEAN NOT NULL
        );`, schemaPrefix),

		// Tabela de Organizações
		fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS %sorganizacao (
           id SERIAL PRIMARY KEY,
           nome VARCHAR(255) NOT NULL,
           responsavel INTEGER NOT NULL REFERENCES %susuario(id),
           cpfcnpj VARCHAR(255) NOT NULL,
           pais VARCHAR(255),
           cidade VARCHAR(255),
           estado VARCHAR(255),
           telefone VARCHAR(20),
           data_cadastro TIMESTAMP NOT NULL
        );`, schemaPrefix, schemaPrefix),

		// Tabela de Relação Usuário-Organização
		fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS %susuario_organizacao (
           id SERIAL PRIMARY KEY,
           usuario INTEGER NOT NULL REFERENCES %susuario(id),
           organizacao INTEGER NOT NULL REFERENCES %sorganizacao(id),
           nivel_acesso VARCHAR(50) NOT NULL,
           data_cadastro TIMESTAMP NOT NULL,
           UNIQUE (usuario, organizacao)
        );`, schemaPrefix, schemaPrefix, schemaPrefix),

		// Tabela de Credenciais
		fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS %scredencial (
           id SERIAL PRIMARY KEY,
           titulo VARCHAR(255) NOT NULL,
           urlApi VARCHAR(255),
           tokenAPI VARCHAR(255),
           instanceAPI VARCHAR(255),
           assistantId VARCHAR(255)
        );`, schemaPrefix),

		// Tabela de Empresas
		fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS %sempresa (
           id SERIAL PRIMARY KEY,
           razao_social VARCHAR(255) NOT NULL,
           cnpj VARCHAR(20) UNIQUE,
           telefone VARCHAR(20),
           site VARCHAR(255),
           pais VARCHAR(255),
           cidade VARCHAR(255),
           estado VARCHAR(255),
           cep VARCHAR(10),
           setor VARCHAR(255) NOT NULL
        );`, schemaPrefix),

		// Tabela de Contatos
		fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS %scontato (
           id SERIAL PRIMARY KEY,
           empresa INTEGER REFERENCES %sempresa(id),
           nome VARCHAR(255) NOT NULL,
           sobrenome VARCHAR(255) NOT NULL,
           cargo VARCHAR(255) NOT NULL,
           email VARCHAR(255),
           telefone VARCHAR(20) NOT NULL,
           responsavel INTEGER NOT NULL REFERENCES %susuario(id)
        );`, schemaPrefix, schemaPrefix, schemaPrefix),

		// Tabela de Leads
		fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS %slead (
           id SERIAL PRIMARY KEY,
           nome VARCHAR(255) NOT NULL,
           email VARCHAR(255),
           telefone VARCHAR(20) NOT NULL,
           empresa VARCHAR(255),
           origem VARCHAR(255),
           status VARCHAR(50) NOT NULL,
           responsavel INTEGER NOT NULL REFERENCES %susuario(id)
        );`, schemaPrefix, schemaPrefix),

		// Tabela de Oportunidades
		fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS %soportunidade (
           id SERIAL PRIMARY KEY,
           titulo VARCHAR(255) NOT NULL,
           empresa INTEGER NOT NULL REFERENCES %sempresa(id),
           contato INTEGER NOT NULL REFERENCES %scontato(id),
           valor_estimado DECIMAL(10, 2) NOT NULL,
           etapa_funil VARCHAR(50) NOT NULL,
           probabilidade INTEGER,
           status VARCHAR(50) NOT NULL,
           responsavel INTEGER NOT NULL REFERENCES %susuario(id),
           data_criacao TIMESTAMP NOT NULL
        );`, schemaPrefix, schemaPrefix, schemaPrefix, schemaPrefix),

		// Tabela de Campanhas
		fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS %scampanha (
           id SERIAL PRIMARY KEY,
           nome VARCHAR(255) NOT NULL,
           tipo VARCHAR(50) NOT NULL,
           data_inicio TIMESTAMP NOT NULL,
           data_fim TIMESTAMP NOT NULL,
           orcamento DECIMAL(10, 2) NOT NULL,
           status VARCHAR(50) NOT NULL,
           responsavel INTEGER NOT NULL REFERENCES %susuario(id)
        );`, schemaPrefix, schemaPrefix),

		// Tabela de Membros de Campanha
		fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS %scampanha_membros (
           id SERIAL PRIMARY KEY,
           campanha INTEGER NOT NULL REFERENCES %scampanha(id),
           lead INTEGER REFERENCES %slead(id),
           contato INTEGER REFERENCES %scontato(id),
           status_resposta VARCHAR(50) NOT NULL,
           CONSTRAINT chk_lead_or_contato CHECK ((lead IS NOT NULL AND contato IS NULL) OR (lead IS NULL AND contato IS NOT NULL))
        );`, schemaPrefix, schemaPrefix, schemaPrefix, schemaPrefix),

		// Tabela de Atividades
		fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS %satividade (
           id SERIAL PRIMARY KEY,
           assunto VARCHAR(255) NOT NULL,
           tipo VARCHAR(50) NOT NULL,
           data_vencimento TIMESTAMP NOT NULL,
           status VARCHAR(50),
           descricao TEXT NOT NULL,
           data_criacao TIMESTAMP NOT NULL
        );`, schemaPrefix),

		// Tabela de Tickets
		fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS %sticket (
           id SERIAL PRIMARY KEY,
           assunto VARCHAR(255) NOT NULL,
           descricao TEXT NOT NULL,
           contato INTEGER NOT NULL REFERENCES %scontato(id),
           status VARCHAR(50) NOT NULL,
           prioridade VARCHAR(50) NOT NULL,
           responsavel INTEGER NOT NULL REFERENCES %susuario(id),
           data_abertura TIMESTAMP NOT NULL,
           data_fechamento TIMESTAMP
        );`, schemaPrefix, schemaPrefix, schemaPrefix),

		// Tabela de Mensagens
		fmt.Sprintf(`
         CREATE TABLE IF NOT EXISTS %smensagens(
            id SERIAL PRIMARY KEY,
            id_remetente INTEGER NOT NULL REFERENCES %susuario(id),
            id_destinatario INTEGER NOT NULL REFERENCES %susuario(id),
            conteudo_mensagem VARCHAR(1000) NOT NULL,
            data_envio TIMESTAMP NOT NULL,
            status BOOLEAN NOT NULL,
            urgencia VARCHAR(10)
         );`, schemaPrefix, schemaPrefix, schemaPrefix),
	}

	// Executa cada instrução de criação de tabela
	for _, stmt := range createTableStatements {
		_, err := db.Exec(stmt)
		if err != nil {
			log.Fatalf("Erro ao criar a tabela com a instrução: %s\nErro: %v", stmt, err)
		}
	}

	log.Println("Todas as tabelas foram verificadas e inicializadas com sucesso.")
}
