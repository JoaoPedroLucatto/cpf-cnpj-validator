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
  const filterDocument = ref("");
  const filterType = ref("");
  const sortBy = ref("created_at");
  const order = ref("asc");

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
      const response = await getDocuments({
        document: filterDocument.value,
        type: filterType.value,
        sortBy: sortBy.value,
        order: order.value,
      });
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
      const status = err.response?.status;
      const api = err.response?.data;

      if (status == 400) {
        showError("Documento Inválido");
        return;
      }

      if (status === 409) {
        showError("Documento já existente!");
      } else {
        showError(api?.details || api?.error || "Erro ao criar documento");
      }
    }
  };

  const edit = async (id: string, newDoc: string) => {
    try {
      await editDocument(id, newDoc);
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
      showSuccess(
        `Documento ${blocked ? "adicionado" : "removido"} da blocklist`
      );
      await load();
    } catch (err: any) {
      showError(err.message || "Erro ao atualizar blocklist");
    }
  };

  onMounted(load);

  return {
    documents,
    newDocument,
    filterDocument,
    filterType,
    sortBy,
    order,
    successMessage,
    errorMessage,
    create,
    remove,
    toggleBlock,
    edit,
    load,
  };
}
