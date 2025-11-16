import { createRouter, createWebHistory } from "vue-router";
import documentRoute from "../modules/documents/routes/document.route";

const router = createRouter({
  history: createWebHistory(),
  routes: [{ path: "/", redirect: "/documents" }, ...documentRoute],
});

export default router;
