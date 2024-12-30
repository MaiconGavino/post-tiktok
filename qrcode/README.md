# QR Code Generator com Golang e Vue.js 🚀

Este projeto é uma aplicação simples para gerar QR Codes a partir de links ou textos fornecidos pelo usuário. Utilizando **Golang** no backend e **Vue.js** no frontend, ele demonstra como integrar essas tecnologias de maneira eficiente e funcional.

---

## 🎯 **Objetivo**

O objetivo do projeto é proporcionar uma solução prática para criar QR Codes enquanto ensina conceitos importantes de integração entre backend e frontend. Este exemplo é ideal para quem está aprendendo ou deseja explorar as capacidades de **Golang** e **Vue.js**.

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

## 📌 **Como Funciona?**

1. O usuário insere um texto ou link na interface web.
2. O frontend envia o texto para o backend por meio de uma requisição HTTP POST.
3. O backend utiliza a biblioteca `go-qrcode` para gerar o QR Code e retorna uma imagem PNG.
4. O frontend exibe o QR Code gerado para o usuário.

---

## 🚀 **Como Rodar o Projeto**

### **Pré-requisitos**
- [Golang instalado](https://golang.org/dl/)
- Um navegador web para acessar a interface

### **Passo a Passo**

1. Clone o repositório:
   ```bash
   git clone https://github.com/MaiconGavino/post-tiktok/tree/main/qrcode
   cd qrcode
   ```

2. Instale as dependências do backend:
   ```bash
   cd backend
   go mod tidy
   ```

3. Inicie o servidor backend:
   ```bash
   go run main.go
   ```

4. Abra o arquivo `index.html` no navegador:
   ```bash
   cd ../frontend
   open index.html
   ```

---

## 📂 **Estrutura do Projeto**

```plaintext
qr-code-generator/
├── backend/
│   ├── main.go          # Código do servidor em Golang
│   ├── go.mod           # Gerenciador de dependências do Go
│   └── go.sum           # Checksum das dependências do Go
├── frontend/
│   ├── index.html       # Interface do usuário
│   ├── style.css        # Estilo da interface
└── README.md            # Documentação do projeto
```

---

## ✨ **Melhorias Futuras**

- [ ] Adicionar funcionalidade para download do QR Code gerado.
- [ ] Permitir personalização do QR Code (cores, tamanhos, etc.).
- [ ] Implementar histórico de QR Codes gerados.

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