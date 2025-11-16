# 📘 CPF e CNPJ — Validação

## 📌 Visão Geral

Este projeto é composto por uma API em Go (Golang) e uma SPA em Vue + TypeScript, responsáveis por validar CPF e CNPJ e expor funcionalidades via endpoints REST.

A arquitetura foi projetada com foco em flexibilidade, escalabilidade e manutenibilidade.

---

## 🧱 Microsserviços

### `cpf-cnpj-api`
Responsável por:
- Disponibilizar endpoints GET, POST, PATCH e DELETE
- Processar regras de validação
- Persistir e consultar dados

### `cpf-cnpj-ui`
Responsável por:
- Consumir os endpoints da API
- Entregar uma interface SPA construída com Vue + TypeScript

---

## 🧠 Abordagem Técnica

### Clean Architecture
O projeto foi estruturado seguindo os princípios da **Clean Architecture**, visando:
- Separação clara entre lógica de negócio, controle de fluxo e detalhes de infraestrutura

### Clean Code
A base de código segue práticas de legibilidade e padronização, como:
- Métodos com responsabilidades únicas
- Nomenclatura clara e objetiva
- Redução de dependências acopladas diretamente

---

## 🚀 Subindo o Projeto

Precisa conter docker instalado na sua máquina.

```bash
make build up logs
```

Caso precise matar os containers e limpar o volume criado.

```bash
make clean
```

## 🛠️ Comandos Extras

Rodar todos os testes automatizados

```bash
make clean
```

Executar análise estática e linters

```bash
make lint
```

