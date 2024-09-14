# Acervo Comics

Acervo Comics é uma aplicação para gestão e compartilhamento de quadrinhos. O sistema permite que os usuários adicionem quadrinhos ao seu acervo pessoal a partir do ISBN, vejam quadrinhos cadastrados e também realizem login e registro de usuários.

## Funcionalidades

- Registro de usuários
- Login de usuários com autenticação via JWT
- Adição de quadrinhos ao acervo pessoal por meio do ISBN
- Exibição dos quadrinhos vinculados a cada usuário
- Perfil do usuário logado
- [Futuro] Troca de quadrinhos entre usuários
- [Futuro] Avaliações de quadrinhos
- [Futuro] Notificações

## Tecnologias Utilizadas

- **Linguagem**: Go (Golang)
- **Framework Web**: [Fiber](https://gofiber.io/)
- **Banco de Dados**: PostgreSQL
- **ORM**: GORM
- **Autenticação**: JWT (JSON Web Tokens)
- **Documentação**: Swagger
- **API**: BrasilAPI (para buscar informações de quadrinhos pelo ISBN)
- **Log**: Zerolog

## Instalação e Configuração

### Requisitos

- Go 1.22 ou superior
- PostgreSQL
- Git

### Passos para Instalação

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

    ```bash
    go mod tidy
    ```

4. Execute as migrações para criar as tabelas no banco de dados:

    ```bash
    go run main.go
    ```

### Rodando a Aplicação

Após seguir os passos de configuração, você pode iniciar a aplicação com:

```bash
go run main.go
