# Servidor Web em Go

Este é um exemplo de servidor web simples em Go utilizando o pacote `net/http`. Ele serve arquivos estáticos, trata rotas de formulários e responde com uma mensagem personalizada.

## Funcionalidades

- **Rota `/hello`**: Responde com a mensagem "Hello!" quando acessada via método GET.
- **Rota `/form`**: Recebe dados via formulário POST (nome e email) e exibe essas informações.
- **Arquivos Estáticos**: Serve arquivos da pasta `./static`.

## Fluxo do Sistema

Abaixo está o fluxo básico do servidor:

![Fluxo Web Serve com GO](WebServeGo/fluxo.png)

## Requisitos

- Go (versão 1.18 ou superior)

## Como Rodar o Código

### 1. Instalar o Go

Caso ainda não tenha o Go instalado, siga os passos abaixo para instalá-lo:

- Baixe e instale a versão mais recente do Go [aqui](https://golang.org/dl/).
- Após a instalação, verifique se o Go foi instalado corretamente executando o comando abaixo no terminal:

```bash
go version
```

### 2. Clonar o Repositório

Clone o repositório para o seu ambiente local:

```bash
git clone <git@github.com:MaiconGavino/post-tiktok.git>
```

### 3. Executar o Servidor

1. Navegue até o diretório do projeto:

```bash
cd WebServeGo
```

2. Execute o servidor Go:

```bash
go run main.go
```

O servidor será iniciado na porta `8080` por padrão.

### 4. Acessar as Rotas

- Abra o navegador e acesse a seguinte URL para testar a rota `/hello`:

```
http://localhost:8080/hello
```

Você verá a mensagem "Hello!".

- Para testar a rota `/form`e abra no navegador em:

```
http://localhost:8080/form
```

Após preencher e enviar o formulário, os dados de `nome` e `email` serão exibidos na tela.

### 5. Servir Arquivos Estáticos

Os arquivos estáticos (HTML, CSS, imagens, etc.) são servidos a partir da pasta `./static`. Coloque seus arquivos nesta pasta e acesse-os diretamente no navegador, por exemplo:

```
http://localhost:8080/form.html
```

## Estrutura do Projeto

```
/nome-do-repositorio
  ├── main.go           # Código do servidor Go
  └── /static           # Pasta contendo arquivos estáticos
      └── form.html     # Exemplo de arquivo HTML para o formulário
      └── hello.html	# Exemplo de arquivo HTML para o hello
```



Se precisar de mais alguma coisa, estou à disposição!
