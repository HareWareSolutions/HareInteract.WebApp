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

	conexao := fmt.Sprintf("user=postgres dbname=HareInteractCRM password=HareWare@2024 host=localhost sslmode=disable search_path=%s,public", search_path)

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
	AtualizaEstrutura(db)

	return db
}

func InicializaTabelas(db *sql.DB, search_path string) {
	log.Println("Verificando e inicializando tabelas do banco de dados...")

	schemaPrefix := fmt.Sprintf("%s.", search_path)

	createTableStatements := []string{
		fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS %susuario (
           id SERIAL PRIMARY KEY,
           nome VARCHAR(255) NOT NULL,
           email VARCHAR(255) UNIQUE NOT NULL,
           username VARCHAR(255) UNIQUE NOT NULL,
           senha VARCHAR(255) NOT NULL,
           ativo BOOLEAN NOT NULL
        );`, schemaPrefix),

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

		fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS %susuario_organizacao (
           id SERIAL PRIMARY KEY,
           usuario INTEGER NOT NULL REFERENCES %susuario(id),
           organizacao INTEGER NOT NULL REFERENCES %sorganizacao(id),
           nivel_acesso VARCHAR(50) NOT NULL,
           data_cadastro TIMESTAMP NOT NULL,
           UNIQUE (usuario, organizacao)
        );`, schemaPrefix, schemaPrefix, schemaPrefix),

		fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS %scredencial (
           id SERIAL PRIMARY KEY,
           titulo VARCHAR(255) NOT NULL,
           urlApi VARCHAR(255),
           tokenAPI VARCHAR(255),
           instanceAPI VARCHAR(255),
           assistantId VARCHAR(255)
        );`, schemaPrefix),

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

		fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS %scontato (
           id SERIAL PRIMARY KEY,
           empresa INTEGER REFERENCES %sempresa(id),
           nome VARCHAR(255) NOT NULL,
           sobrenome VARCHAR(255) NOT NULL,
           cargo VARCHAR(255) NOT NULL,
           email VARCHAR(255),
           telefone VARCHAR(20) NOT NULL,
           responsavel INTEGER NOT NULL REFERENCES %susuario(id),
           pausa BOOLEAN NOT NULL,
           thread_id VARCHAR(255) NOT NULL
        );`, schemaPrefix, schemaPrefix, schemaPrefix),

		fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS %slead (
           id SERIAL PRIMARY KEY,
           nome VARCHAR(255) NOT NULL,
           email VARCHAR(255),
           telefone VARCHAR(20) NOT NULL,
           empresa VARCHAR(255),
           origem VARCHAR(255),
           status VARCHAR(50) NOT NULL,
           responsavel INTEGER NOT NULL REFERENCES %susuario(id),
           pausa BOOLEAN NOT NULL,
           thread_id VARCHAR(255) NOT NULL
        );`, schemaPrefix, schemaPrefix),

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

		fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS %scampanha_membros (
           id SERIAL PRIMARY KEY,
           campanha INTEGER NOT NULL REFERENCES %scampanha(id),
           lead INTEGER REFERENCES %slead(id),
           contato INTEGER REFERENCES %scontato(id),
           status_resposta VARCHAR(50) NOT NULL,
           CONSTRAINT chk_lead_or_contato CHECK ((lead IS NOT NULL AND contato IS NULL) OR (lead IS NULL AND contato IS NOT NULL))
        );`, schemaPrefix, schemaPrefix, schemaPrefix, schemaPrefix),

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

		fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %sagendamento (
			id SERIAL PRIMARY KEY,
			data DATE NOT NULL,
			hora TIME NOT NULL,
			contato_id INTEGER NOT NULL REFERENCES %scontato(id) ON DELETE CASCADE,
			usuario_id INTEGER NOT NULL REFERENCES %susuario(id) ON DELETE CASCADE,
			confirmacao BOOLEAN NOT NULL,
			observacao TEXT NOT NULL,
			link TEXT
		);`, schemaPrefix, schemaPrefix, schemaPrefix),

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
            urgencia VARCHAR(10),
            id_organizacao_convite INTEGER DEFAULT 0,
            nivel_acesso VARCHAR(20) DEFAULT 'NONE'
         );`, schemaPrefix, schemaPrefix, schemaPrefix),

		fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %sfuncao (
			id SERIAL PRIMARY KEY,
			titulo TEXT NOT NULL,
			descricao TEXT NOT NULL,
			chamada TEXT NOT NULL
		);`, schemaPrefix),
	}

	for _, stmt := range createTableStatements {
		_, err := db.Exec(stmt)
		if err != nil {
			log.Fatalf("Erro ao criar a tabela com a instrução: %s\nErro: %v", stmt, err)
		}
	}

	log.Println("Todas as tabelas foram verificadas e inicializadas com sucesso.")
}

func AtualizaEstrutura(db *sql.DB) {
	log.Println("Iniciando verificação de atualizações de estrutura em todos os schemas...")

	// 1- Buscar todos os schemas válidos (exceto os padrões do PostgreSQL)
	rows, err := db.Query(`
		SELECT schema_name 
		FROM information_schema.schemata 
		WHERE schema_name NOT IN ('pg_catalog', 'information_schema', 'public')
		ORDER BY schema_name;
	`)
	if err != nil {
		log.Fatalf("Erro ao listar schemas: %v", err)
	}
	defer rows.Close()

	var schemas []string
	for rows.Next() {
		var schema string
		if err := rows.Scan(&schema); err != nil {
			log.Printf("Erro ao ler schema: %v", err)
			continue
		}
		schemas = append(schemas, schema)
	}

	if len(schemas) == 0 {
		log.Println("Nenhum schema customizado encontrado para atualizar.")
		return
	}

	// 2- Executar scripts de atualização em cada schema encontrado
	for _, schema := range schemas {
		schemaPrefix := fmt.Sprintf("%s.", schema)
		log.Printf("Atualizando estrutura no schema: %s", schema)

		scripts := []string{
			// Adicionar campo "pausa" em LEAD
			fmt.Sprintf(`
				DO $$
				BEGIN
					IF NOT EXISTS (
						SELECT 1 FROM information_schema.columns
						WHERE table_schema = '%s' AND table_name = 'lead' AND column_name = 'pausa'
					) THEN
						ALTER TABLE %slead ADD COLUMN pausa BOOLEAN DEFAULT FALSE;
					END IF;
				END$$;`, schema, schemaPrefix),

			// Adicionar campo "thread_id" em LEAD
			fmt.Sprintf(`
				DO $$
				BEGIN
					IF NOT EXISTS (
						SELECT 1 FROM information_schema.columns
						WHERE table_schema = '%s' AND table_name = 'lead' AND column_name = 'thread_id'
					) THEN
						ALTER TABLE %slead ADD COLUMN thread_id VARCHAR(255);
					END IF;
				END$$;`, schema, schemaPrefix),

			// Adicionar campo "pausa" em CONTATO
			fmt.Sprintf(`
				DO $$
				BEGIN
					IF NOT EXISTS (
						SELECT 1 FROM information_schema.columns
						WHERE table_schema = '%s' AND table_name = 'contato' AND column_name = 'pausa'
					) THEN
						ALTER TABLE %scontato ADD COLUMN pausa BOOLEAN DEFAULT FALSE;
					END IF;
				END$$;`, schema, schemaPrefix),

			// Adicionar campo "thread_id" em CONTATO
			fmt.Sprintf(`
				DO $$
				BEGIN
					IF NOT EXISTS (
						SELECT 1 FROM information_schema.columns
						WHERE table_schema = '%s' AND table_name = 'contato' AND column_name = 'thread_id'
					) THEN
						ALTER TABLE %scontato ADD COLUMN thread_id VARCHAR(255);
					END IF;
				END$$;`, schema, schemaPrefix),

			// Adicionar campo "id_organizacao_convite" em MENSAGENS
			fmt.Sprintf(`
				DO $$
				BEGIN
					IF NOT EXISTS (
						SELECT 1 FROM information_schema.columns
						WHERE table_schema = '%s' AND table_name = 'mensagens' AND column_name = 'id_organizacao_convite'
					) THEN
						ALTER TABLE %smensagens ADD COLUMN id_organizacao_convite INTEGER DEFAULT 0;
					END IF;
				END$$;`, schema, schemaPrefix),

			// Adicionar campo "nivel_acesso" em MENSAGENS
			fmt.Sprintf(`
				DO $$
				BEGIN
					IF NOT EXISTS (
						SELECT 1 FROM information_schema.columns
						WHERE table_schema = '%s' AND table_name = 'mensagens' AND column_name = 'nivel_acesso'
					) THEN
						ALTER TABLE %smensagens ADD COLUMN nivel_acesso VARCHAR(20) DEFAULT 'NONE';
					END IF;
				END$$;`, schema, schemaPrefix),
		}

		for i, script := range scripts {
			log.Printf("Executando script %d no schema %s...", i+1, schema)
			_, err := db.Exec(script)
			if err != nil {
				log.Printf("Erro ao aplicar script %d em %s: %v", i+1, schema, err)
			} else {
				log.Printf("Script %d aplicado com sucesso em %s", i+1, schema)
			}
		}

		log.Printf("Estrutura atualizada com sucesso no schema: %s", schema)
	}

	log.Println("Atualizações de estrutura concluídas para todos os schemas.")
}
