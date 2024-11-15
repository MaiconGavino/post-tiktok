# WeatherAPI - Golang

## Descrição

**WeatherAPI** é uma API simples desenvolvida em Go (Golang) que consome a API do OpenWeather para retornar informações climáticas de uma cidade específica. A aplicação retorna a temperatura atual da cidade em formato JSON. 

## Funcionalidades

- Consulta a temperatura atual de uma cidade utilizando a API do OpenWeather.
- Resposta em formato JSON com o nome da cidade e a temperatura em graus Celsius.
- Cidade padrão configurada como "Porto Velho" caso nenhuma cidade seja especificada na requisição.

## Requisitos

- Go 1.16 ou superior.
- Conta no [OpenWeather](https://openweathermap.org/) para obter a API Key.
- Arquivo `.env` para configuração da API Key.

## Configuração

### 1. Clonar o repositório

```bash
git clone git@github.com:MaiconGavino/post-tiktok.git
cd ApiWeather
```

### 2. Configurar o arquivo `.env`

Crie um arquivo `.env` na raiz do projeto e adicione a seguinte linha:

```env
API_KEY=your_openweather_api_key
```

Substitua `your_openweather_api_key` pela sua chave de API do OpenWeather.

### 3. Instalar dependências

Certifique-se de que todas as dependências estão instaladas:

```bash
go mod tidy
```

### 4. Executar o servidor

Para iniciar o servidor:

```bash
go run main.go
```

O servidor estará disponível na porta `8080`.

## Uso

A API fornece um endpoint para consulta climática:

### Endpoint: `/weather`

**Método:** `GET`  
**Parâmetros de consulta:**
- `city` (opcional): Nome da cidade para consulta (exemplo: `city=Cuiaba`).

**Exemplo de Requisição:**

```bash
curl "http://localhost:8080/weather?city=Cuiaba"
```

**Exemplo de Resposta:**

```json
{
  "city": "Cuiaba",
  "temperatura": 15.5
}
```

## Erros Possíveis

- `500 Internal Server Error` se a API Key não estiver configurada ou se houver problemas ao acessar a API do OpenWeather.
- `500 Internal Server Error` se houver erros na leitura ou decodificação dos dados da API do OpenWeather.

## Dependências

- [godotenv](https://github.com/joho/godotenv): Carrega variáveis de ambiente de um arquivo `.env`.

## Autor

Desenvolvido por Maicon Gavino.