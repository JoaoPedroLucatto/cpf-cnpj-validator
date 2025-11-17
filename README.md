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

## 📚 Documentação da API

### Cadastra um documento

**POST /documents**

| Parâmetro  | Tipo     | Descrição                                     |
| :--------- | :------- | :-------------------------------------------- |
| `document` | `string` | **Obrigatório**. Número ou valor do documento |

**Resposta de sucesso (201 Created):**

```json
{
  "document": {
    "id": "b34e6d4d-3c1c-44ff-ac73-04b5b43913b1",
    "number": "72664691000139",
    "type": "CNPJ",
    "blocked": false,
    "createdAt": "2025-11-17T00:39:24.664908626Z",
    "updatedAt": "2025-11-17T00:39:24.664908626Z",
    "deletedAt": null
  }
}
```

**Resposta se já existe (409 Conflict):**

```json
{
  "message": "documento existed"
}
```

---

### Retorna um documento específico

**GET /documents/:document**

| Parâmetro  | Tipo     | Descrição                                |
| :--------- | :------- | :--------------------------------------- |
| `document` | `string` | **Obrigatório**. Documento a ser buscado |

**Resposta de sucesso (200 OK):**

```json
{
  "document": {
    "id": "b34e6d4d-3c1c-44ff-ac73-04b5b43913b1",
    "number": "72664691000139",
    "type": "CNPJ",
    "blocked": false,
    "createdAt": "2025-11-17T00:39:24.664908626Z",
    "updatedAt": "2025-11-17T00:39:24.664908626Z",
    "deletedAt": null
  }
}
```

**Resposta se não encontrado (404 Not Found):**

```json
{
  "error": "document not found"
}
```

---

### Atualiza um documento

**PATCH /documents/:id**

| Parâmetro  | Tipo     | Descrição                                         |
| :--------- | :------- | :------------------------------------------------ |
| `id`       | `string` | **Obrigatório**. ID do documento a ser atualizado |
| `document` | `string` | **Obrigatório**. Novo valor do documento          |

**Resposta de sucesso (202 Accepted):**

```json
{
  "document": {
    "id": "b34e6d4d-3c1c-44ff-ac73-04b5b43913b1",
    "number": "98765432100",
    "type": "CNPJ",
    "blocked": false,
    "createdAt": "2025-11-17T00:39:24.664908626Z",
    "updatedAt": "2025-11-17T00:50:00.123456789Z",
    "deletedAt": null
  }
}
```

---

### Deleta um documento

**DELETE /documents/:id**

| Parâmetro | Tipo     | Descrição                                       |
| :-------- | :------- | :---------------------------------------------- |
| `id`      | `string` | **Obrigatório**. ID do documento a ser deletado |

**Resposta de sucesso (202 Accepted):**

```json
{
  "document": {
    "id": "b34e6d4d-3c1c-44ff-ac73-04b5b43913b1",
    "number": "72664691000139",
    "type": "CNPJ",
    "blocked": false,
    "createdAt": "2025-11-17T00:39:24.664908626Z",
    "updatedAt": "2025-11-17T00:39:24.664908626Z",
    "deletedAt": "2025-11-17T01:00:00.000000000Z"
  }
}
```

---

### Lista documentos

**GET /documents**

| Parâmetro  | Tipo     | Descrição                                         |
| :--------- | :------- | :------------------------------------------------ |
| `document` | `string` | Opcional. Filtra por número do documento          |
| `type`     | `string` | Opcional. Tipo do documento                       |
| `sortBy`   | `string` | Opcional. Campo para ordenar (padrão: created_at) |
| `order`    | `string` | Opcional. Ordem da ordenação (asc ou desc)        |

**Resposta de sucesso (200 OK):**

```json
{
  "documents": [
    {
      "id": "b34e6d4d-3c1c-44ff-ac73-04b5b43913b1",
      "number": "72664691000139",
      "type": "CNPJ",
      "blocked": false,
      "createdAt": "2025-11-17T00:39:24.664908626Z",
      "updatedAt": "2025-11-17T00:39:24.664908626Z",
      "deletedAt": null
    },
    {
      "id": "c56a6d4d-3c1c-44ff-ac73-04b5b4390000",
      "number": "12345678000190",
      "type": "CNPJ",
      "blocked": false,
      "createdAt": "2025-11-17T01:10:00.123456789Z",
      "updatedAt": "2025-11-17T01:10:00.123456789Z",
      "deletedAt": null
    }
  ]
}
```

---

### Bloqueia ou desbloqueia um documento

**PATCH /documents/:id/blocklist**

| Parâmetro | Tipo     | Descrição                                                      |
| :-------- | :------- | :------------------------------------------------------------- |
| `id`      | `string` | **Obrigatório**. ID do documento                               |
| `blocked` | `bool`   | **Obrigatório**. Define se bloqueia ou desbloqueia o documento |

**Resposta de sucesso (200 OK):**

```json
{
  "id": "b34e6d4d-3c1c-44ff-ac73-04b5b43913b1",
  "blocked": true
}
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

