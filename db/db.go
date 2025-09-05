package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func ConectaBD(search_path string) *sql.DB {
	conexao := "user=postgres dbname=HareInteractCRM password=HareWare@2024 host=localhost sslmode=disable search_path=" + search_path

	db, err := sql.Open("postgres", conexao)
	if err != nil {
		log.Fatalf("Erro ao abrir a conexão com o banco de dados: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Erro ao pingar o banco de dados: %v", err)
	}

	InicializaTabelas(db)

	return db
}

func InicializaTabelas(db *sql.DB) {
	log.Println("Verificando e inicializando tabelas do banco de dados...")

	createAllTablesSQL := `
	-- Tabela de Usuários
	CREATE TABLE IF NOT EXISTS usuario (
		id SERIAL PRIMARY KEY,
		nome VARCHAR(255) NOT NULL,
		email VARCHAR(255) UNIQUE NOT NULL,
		username VARCHAR(255) UNIQUE NOT NULL,
		senha VARCHAR(255) NOT NULL,
		ativo BOOLEAN NOT NULL
	);

	-- Tabela de Organizações
	CREATE TABLE IF NOT EXISTS organizacao (
		id SERIAL PRIMARY KEY,
		nome VARCHAR(255) NOT NULL,
		responsavel INTEGER NOT NULL REFERENCES usuario(id),
		pais VARCHAR(255),
		cidade VARCHAR(255),
		estado VARCHAR(255),
		telefone VARCHAR(20),
		data_cadastro TIMESTAMP NOT NULL
	);

	-- Tabela de Relação Usuário-Organização
	CREATE TABLE IF NOT EXISTS usuario_organizacao (
		id SERIAL PRIMARY KEY,
		usuario INTEGER NOT NULL REFERENCES usuario(id),
		organizacao INTEGER NOT NULL REFERENCES organizacao(id),
		nivel_acesso VARCHAR(50) NOT NULL,
		data_cadastro TIMESTAMP NOT NULL,
		UNIQUE (usuario, organizacao)
	);

	-- Tabela de Empresas
	CREATE TABLE IF NOT EXISTS empresa (
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
	);

	-- Tabela de Contatos
	CREATE TABLE IF NOT EXISTS contato (
		id SERIAL PRIMARY KEY,
		empresa INTEGER REFERENCES empresa(id),
		nome VARCHAR(255) NOT NULL,
		sobrenome VARCHAR(255) NOT NULL,
		cargo VARCHAR(255) NOT NULL,
		email VARCHAR(255),
		telefone VARCHAR(20) NOT NULL,
		responsavel INTEGER NOT NULL REFERENCES usuario(id)
	);

	-- Tabela de Leads
	CREATE TABLE IF NOT EXISTS lead (
		id SERIAL PRIMARY KEY,
		nome VARCHAR(255) NOT NULL,
		email VARCHAR(255),
		telefone VARCHAR(20) NOT NULL,
		empresa VARCHAR(255),
		origem VARCHAR(255),
		status VARCHAR(50) NOT NULL,
		responsavel INTEGER NOT NULL REFERENCES usuario(id)
	);

	-- Tabela de Oportunidades
	CREATE TABLE IF NOT EXISTS oportunidade (
		id SERIAL PRIMARY KEY,
		titulo VARCHAR(255) NOT NULL,
		empresa INTEGER NOT NULL REFERENCES empresa(id),
		contato INTEGER NOT NULL REFERENCES contato(id),
		valor_estimado DECIMAL(10, 2) NOT NULL,
		etapa_funil VARCHAR(50) NOT NULL,
		probabilidade INTEGER,
		status VARCHAR(50) NOT NULL,
		responsavel INTEGER NOT NULL REFERENCES usuario(id),
		data_criacao TIMESTAMP NOT NULL
	);

	-- Tabela de Campanhas
	CREATE TABLE IF NOT EXISTS campanha (
		id SERIAL PRIMARY KEY,
		nome VARCHAR(255) NOT NULL,
		tipo VARCHAR(50) NOT NULL,
		data_inicio TIMESTAMP NOT NULL,
		data_fim TIMESTAMP NOT NULL,
		orcamento DECIMAL(10, 2) NOT NULL,
		status VARCHAR(50) NOT NULL,
		responsavel INTEGER NOT NULL REFERENCES usuario(id)
	);

	-- Tabela de Membros de Campanha
	CREATE TABLE IF NOT EXISTS campanha_membros (
		id SERIAL PRIMARY KEY,
		campanha INTEGER NOT NULL REFERENCES campanha(id),
		lead INTEGER REFERENCES lead(id),
		contato INTEGER REFERENCES contato(id),
		status_resposta VARCHAR(50) NOT NULL,
		CONSTRAINT chk_lead_or_contato CHECK ((lead IS NOT NULL AND contato IS NULL) OR (lead IS NULL AND contato IS NOT NULL))
	);

	-- Tabela de Atividades
	CREATE TABLE IF NOT EXISTS atividade (
		id SERIAL PRIMARY KEY,
		assunto VARCHAR(255) NOT NULL,
		tipo VARCHAR(50) NOT NULL,
		data_vencimento TIMESTAMP NOT NULL,
		status VARCHAR(50),
		descricao TEXT NOT NULL,
		data_criacao TIMESTAMP NOT NULL
	);

	-- Tabela de Tickets
	CREATE TABLE IF NOT EXISTS ticket (
		id SERIAL PRIMARY KEY,
		assunto VARCHAR(255) NOT NULL,
		descricao TEXT NOT NULL,
		contato INTEGER NOT NULL REFERENCES contato(id),
		status VARCHAR(50) NOT NULL,
		prioridade VARCHAR(50) NOT NULL,
		responsavel INTEGER NOT NULL REFERENCES usuario(id),
		data_abertura TIMESTAMP NOT NULL,
		data_fechamento TIMESTAMP
	);
	`

	_, err := db.Exec(createAllTablesSQL)
	if err != nil {
		log.Fatalf("Erro ao criar as tabelas: %v", err)
	}

	log.Println("Todas as tabelas foram verificadas e inicializadas com sucesso.")
}
