<template>
  <div class="container mt-4">
    <h1 class="mb-4">Documentos</h1>

    <BaseAlert :message="successMessage" type="success" />
    <BaseAlert :message="errorMessage" type="error" />

    <form @submit.prevent="create" class="mb-3">
      <div class="input-group">
        <input
          v-model="newDocument"
          type="text"
          class="form-control"
          placeholder="Digite um documento"
        />
        <button class="btn btn-primary" type="submit">Criar</button>
      </div>
    </form>

    <form @submit.prevent="load" class="mb-3 d-flex gap-2">
      <input v-model="filterDocument" type="text" placeholder="Filtrar por documento" class="form-control" />
      <input v-model="filterType" type="text" placeholder="Filtrar por tipo" class="form-control" />
      <select v-model="sortBy" class="form-select">
        <option value="created_at">Criado em</option>
        <option value="number">Número</option>
      </select>
      <select v-model="order" class="form-select">
        <option value="asc">Ascendente</option>
        <option value="desc">Descendente</option>
      </select>
      <button type="submit" class="btn btn-primary">Filtrar</button>
    </form>

    <table class="table table-bordered mt-4">
      <thead>
        <tr>
          <th>ID</th>
          <th>Documento</th>
          <th>Blocklist</th>
          <th>Ações</th>
        </tr>
      </thead>

      <tbody>
        <tr v-for="doc in documents" :key="doc.id">
          <td>{{ doc.id }}</td>
          <td>
            <input
              type="text"
              v-model="doc.number"
              class="form-control form-control-sm"
              @blur="edit(doc.id, doc.number)"
            />
            <small class="text-muted">{{ doc.type }}</small>
          </td>
          <td>
            <button
              class="btn"
              :class="doc.blocked ? 'btn-warning btn-sm' : 'btn-secondary btn-sm'"
              @click="toggleBlock(doc.id, !doc.blocked)"
            >
              {{ doc.blocked ? "Bloqueado" : "Desbloqueado" }}
            </button>
          </td>
          <td>
            <button class="btn btn-danger btn-sm" @click="remove(doc.id)">
              Excluir
            </button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script lang="ts">
import BaseAlert from "../../../shared/components/BaseAlert.vue";
import { useDocumentsViewModel } from "../composables/useDocumentsViewModel";

export default {
  components: { BaseAlert },
  setup() {
    return useDocumentsViewModel();
  },
};
</script>
