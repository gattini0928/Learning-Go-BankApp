# BankApp CLI

Projeto de aprendizado — uma CLI que simula operações bancárias: criação e login de contas, transferências, compras com débito e cartão de crédito (parceladas em até 12x), checagem de saldo, visualização de faturas e extrato de transferências.

Projeto que me ensinou modelagem de banco de dados em SQLite, queries SQL e conexão com models, além de trabalhar bastante com structs em Go.

**Funcionalidades:** Autenticação, geração automática de cartão de crédito, testes, buscas e queries no banco de dados.

---

## Como Rodar

```bash
go mod init bankapp
go get modernc.org/sqlite
```

```bash
go run main.go
```

---

## Como Usar

Crie uma conta — exemplo de dados válidos:

```
Nome: João Oliveira
Email: joaooliveira@icloud.com
Senha: Password123#
Senha Bancária: 098932
```

Após a criação você é logado automaticamente. Para logins futuros, use email e senha:

```
Email: joaooliveira@icloud.com
Senha: Password123#
```

> ⚠️ Para simular transferências entre usuários, você precisará criar mais de uma conta.

---

## Pontos de Melhoria Futura

Projeto encerrado intencionalmente para avançar no aprendizado — aberto a contribuições e melhorias.

Exemplo de função candidata a refatoração: `MakePurchase` em `bankActions`.