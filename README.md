Acervo Comics

Acervo Comics é uma aplicação para gestão e compartilhamento de quadrinhos. O sistema permite que os usuários adicionem quadrinhos ao seu acervo pessoal a partir do ISBN, vejam quadrinhos cadastrados e também realizem login e registro de usuários.
Funcionalidades

- Registro de usuários
- Login de usuários com autenticação via JWT
- Adição de quadrinhos ao acervo pessoal por meio do ISBN
- Exibição dos quadrinhos vinculados a cada usuário
- Perfil do usuário logado
- [Futuro] Troca de quadrinhos entre usuários
- [Futuro] Avaliações de quadrinhos
- [Futuro] Notificações

Tecnologias Utilizadas

- **Linguagem**: Go (Golang)
- **Framework Web**: [Fiber](https://gofiber.io/)
- **Banco de Dados**: PostgreSQL
- **ORM**: GORM
- **Autenticação**: JWT (JSON Web Tokens)
- **Documentação**: Swagger
- **API**: BrasilAPI (para buscar informações de quadrinhos pelo ISBN)
- **Log**: Zerolog

Instalação e Configuração

Requisitos

- Go 1.22 ou superior
- PostgreSQL
- Git

Passos para Instalação

1. Clone o repositório:

    ```bash
    git clone https://github.com/seu-usuario/acervocomics.git
    cd acervocomics
    ```

2. Crie o arquivo `.env` baseado no modelo `.env.example` e configure as variáveis de ambiente, como a URL do banco de dados:

    ```env
    DATABASE_URL=postgres://usuario:senha@localhost:5432/acervocomics
    ```

3. Instale as dependências:

    go mod tidy
    ```

4. Execute as migrações para criar as tabelas no banco de dados:
   
    go run main.go
    ```

Rodando a Aplicação

Após seguir os passos de configuração, você pode iniciar a aplicação com:

go run main.go

A aplicação será iniciada em http://localhost:3000.
Rotas da API
Usuário

    Registrar um usuário
    POST /user/register
    Request body:

    json

{
  "email": "usuario@exemplo.com",
  "password": "senha",
  "name": "Nome do Usuário"
}

Login de usuário
POST /user/login
Request body:

json

    {
      "email": "usuario@exemplo.com",
      "password": "senha"
    }

    Exibir perfil do usuário logado
    GET /user/me
    Header: Authorization: Bearer <token>

Quadrinhos

    Adicionar um quadrinho ao acervo pelo ISBN
    PUT /comic
    Request body:

    json

    {
      "isbn": "978-3-16-148410-0"
    }

    Header: Authorization: Bearer <token>

    Listar quadrinhos do usuário logado
    GET /comic
    Header: Authorization: Bearer <token>

Documentação da API

A documentação da API pode ser acessada através do Swagger. Com a aplicação rodando, acesse:

http://localhost:3000/swagger/index.html

Estrutura do Projeto

acervoback/
│
├── db/                  # Inicialização e configuração do banco de dados
│   └── db.go
│
├── docs/                # Documentação gerada pelo Swagger
│   ├── docs.go
│   └── swagger.json
│
├── handlers/            # Manipuladores das requisições HTTP
│   ├── comic.handler.go
│   └── user.handler.go
│
├── models/              # Definição dos modelos da aplicação (User, Comic)
│   ├── requests/        # Estruturas de requisição
│   └── responses/       # Estruturas de resposta
│
├── repository/          # Implementação dos repositórios para interação com o banco
│   ├── comic.repo.go
│   ├── jwt.repo.go
│   └── user.repo.go
│
├── services/            # Lógica de negócios (serviços)
│   ├── comic.svc.go
│   └── user.svc.go
│
├── .env.example         # Exemplo de arquivo de variáveis de ambiente
├── air.toml             # Configuração para o live reload com o Air
├── go.mod               # Dependências do projeto
├── go.sum               # Checksums das dependências
└── main.go              # Ponto de entrada da aplicação
