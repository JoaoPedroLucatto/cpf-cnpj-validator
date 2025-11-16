import api from "../../../shared/api";

export async function getDocuments() {
  const res = await api.get("/documents");

  return res.data;
}

export async function editDocument(documentId: string, input: string) {
  const res = await api.patch(`/documents/${documentId}`, {
    document: input
  })

  return res.data
}

export async function createDocument(input: string) {
  const res = await api.post("/documents", {
    document: input
  });

  return res.data;
}

export async function deleteDocument(id: string) {
  const res = await api.delete(`/documents/${id}`);

  return res.data;
}

export async function toggleBlocklist(documentId: string, blocked: boolean) {
  const res = await api.patch(`/documents/${documentId}/blocklist`, { blocked });
  return res.data;
}