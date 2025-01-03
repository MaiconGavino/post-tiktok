# QR Code Generator com Golang e Vue.js 🚀

Este projeto é uma aplicação simples formulário de cadastro de usuário. Utilizando **Golang** no backend e **Vue.js** no frontend, ele demonstra como integrar essas tecnologias de maneira eficiente e funcional.

---

## 🛠️ **Tecnologias Utilizadas**

### **Backend**
- **Linguagem:** Golang
- **Biblioteca:** [`go-qrcode`](https://github.com/skip2/go-qrcode) para geração de QR Codes
- **Servidor HTTP:** Embutido no Go

### **Frontend**
- **Framework:** Vue.js (via CDN)
- **HTML e CSS:** Interface simples e responsiva

---

## 🚀 **Como Rodar o Projeto**

### **Pré-requisitos**
- [Golang instalado](https://golang.org/dl/)
- Um navegador web para acessar a interface

### **Passo a Passo**

1. Clone o repositório:
   ```bash
   git clone https://github.com/MaiconGavino/post-tiktok/tree/main/formulario
   cd qrcode
   ```

2. Instale as dependências:
   ```bash
   go mod tidy
   ```

3. Inicie o servidor backend:
   ```bash
   go run main.go
   ```

---

## 📂 **Estrutura do Projeto**

```plaintext
formulario/

├── main.go          # Código do servidor em Golang
├── go.mod           # Gerenciador de dependências do Go
└── go.sum           # Checksum das dependências do Go
├── template/
│   ├── index.html       # Interface do usuário
│   ├── static/
│       ├── style.css    # Estilo da interface
└── README.md            # Documentação do projeto
```

---

## ✨ **Melhorias Futuras**

- [ ] Adicionar a conexão com o banco de dados.
- [ ] Criar a página de login.

---

## 📥 **Contribua**

Sinta-se à vontade para melhorar este projeto! Envie um pull request ou abra uma issue no repositório.

---

## 💡 **Autor**

Desenvolvido por **Maicon Gavino**  
[LinkedIn](https://www.linkedin.com/in/maicongavino) | [GitHub](https://github.com/MaiconGavino)

---

## 🔗 **Licença**

Este projeto é licenciado sob a [MIT License](LICENSE).
```

Este README fornece informações claras e bem organizadas para facilitar a compreensão e a execução do projeto. Se precisar de ajustes ou personalizações, é só avisar! 😊