import { createRouter, createWebHashHistory } from "vue-router";

const router = new createRouter({
  history: createWebHashHistory(),
  routes: [{ path: "/", component: () => import("../pages/Welcome.vue") }],
});

export default router;
