# Forms de registro

Este é um projeto de formulário de registro de usuários com backend em Go e integração com um banco de dados PostgreSQL. A aplicação permite que os usuários se registrem e armazenem suas informações de forma segura.

## Funcionalidades

- **Formulário de Registro**: Os usuários podem inserir nome, sobrenome, e-mail e senha.
- **Validação de Senha**: A senha é armazenada de forma segura usando hashing.
- **Banco de Dados**: Os dados do usuário são salvos em um banco de dados PostgreSQL.
- **Página de Sucesso**: Após o registro, os usuários são redirecionados para uma página de sucesso.
- **Estilo Personalizado**: A interface do formulário é estilizada com CSS para uma melhor experiência do usuário.

## Estrutura do Projeto

```
Forms/
├── main.go                 # Arquivo principal para inicializar o servidor
├── config/
│   └── db.go               # Configuração e conexão com o banco de dados
├── handlers/
│   └── user_handlers.go    # Handlers para lidar com as requisições HTTP
├── models/
│   └── user.go             # Estrutura do modelo de usuário
├── template/
│   └── index.html          # Template HTML para o formulário de registro
├── static/
│   ├── css/
│   │   └── style.css       # Arquivo de estilo CSS para o formulário
├── go.mod                  # Gerenciamento de dependências do Go
```

## Configuração

### Pré-requisitos

- **Golang**: [Instalar Go](https://golang.org/dl/)
- **PostgreSQL**: [Instalar PostgreSQL](https://www.postgresql.org/download/)
- **Variáveis de Ambiente**: Crie um arquivo `.env` na raiz do projeto com as seguintes configurações:

```
DB_USER=seu_usuario
DB_PASSWORD=sua_senha
DB_HOST=localhost
DB_PORT=5432
DB_NAME=seu_banco
DB_SSLMODE=disable
```

## **Passos para Executar**

1. **Clone o Repositório**:
   ```bash
   git clone https://github.com/MaiconGavino/post-tiktok
   cd Forms
   ```

2. **Instale as Dependências**:
   ```bash
   go mod tidy
   ```

3. **Configure o Banco de Dados**:
   Certifique-se de que o banco de dados PostgreSQL está rodando e as tabelas necessárias foram criadas:

   ```sql
   CREATE TABLE users (
       id SERIAL PRIMARY KEY,
       firstname VARCHAR(100),
       lastname VARCHAR(100),
       email VARCHAR(100) UNIQUE,
       password TEXT
   );
   ```

4. **Execute o Servidor**:
   ```bash
   go run main.go
   ```

5. **Acesse a Aplicação**:
   Abra o navegador e acesse:
   ```
   http://localhost:8080
   ```

## Tecnologias Utilizadas

- **Golang**: Backend para manipulação de requisições e conexão com o banco de dados.
- **PostgreSQL**: Armazenamento de dados dos usuários.
- **HTML e CSS**: Frontend estilizado para uma melhor experiência do usuário.
- **bcrypt**: Para hashing seguro de senhas.
- **godotenv**: Gerenciamento de variáveis de ambiente.

## Próximos Passos

- Implementar validação de campos no formulário.
- Adicionar autenticação de usuários.
- Melhorar mensagens de erro e feedback ao usuário.
- Implementar testes para handlers e conexão com o banco de dados.

## Créditos

- Frontend Baseado no design do [@micaelgomestavares](https://uiverse.io/micaelgomestavares/purple-cow-92).


