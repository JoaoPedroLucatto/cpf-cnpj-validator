import { ref, onMounted } from "vue";
import {
  createDocument,
  deleteDocument,
  editDocument,
  getDocuments,
  toggleBlocklist,
} from "../services/documents.service";

export function useDocumentsViewModel() {
  const documents = ref<any[]>([]);
  const newDocument = ref("");

  const successMessage = ref("");
  const errorMessage = ref("");

  const showSuccess = (msg: string) => {
    successMessage.value = msg;
    setTimeout(() => (successMessage.value = ""), 3000);
  };

  const showError = (msg: string) => {
    errorMessage.value = msg;
    setTimeout(() => (errorMessage.value = ""), 5000);
  };

  const load = async () => {
    try {
      const response = await getDocuments();
      documents.value = response.documents;
    } catch {
      showError("Erro ao carregar documentos");
    }
  };

  const create = async () => {
    if (!newDocument.value.trim()) return;

    try {
      await createDocument(newDocument.value);
      showSuccess("Documento criado com sucesso!");
      newDocument.value = "";
      await load();
    } catch (err: any) {
      const api = err.response?.data;
      showError(api?.details || api?.error || "Erro ao criar documento");
    }
  };

  const edit = async (id: string, newDocument: string) => {
    try {
      await editDocument(id, newDocument);
      showSuccess("Documento atualizado com sucesso!");
      await load();
    } catch (err: any) {
      const api = err.response?.data;
      showError(api?.details || api?.error || "Erro ao atualizar documento");
    }
  };

  const remove = async (id: string) => {
    try {
      await deleteDocument(id);
      showSuccess("Documento removido com sucesso!");
      await load();
    } catch (err: any) {
      const api = err.response?.data;
      showError(api?.details || api?.error || "Erro ao remover documento");
    }
  };

  const toggleBlock = async (id: string, blocked: boolean) => {
    try {
      await toggleBlocklist(id, blocked);
      successMessage.value = `Documento ${
        blocked ? "adicionado" : "removido"
      } da blocklist`;
      await load();
    } catch (err: any) {
      errorMessage.value = err.message || "Erro ao atualizar blocklist";
    }
  };

  onMounted(load);

  return {
    documents,
    newDocument,
    successMessage,
    errorMessage,
    create,
    remove,
    toggleBlock,
    edit,
  };
}
