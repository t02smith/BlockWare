import { createRouter, createWebHashHistory } from "vue-router";

const router = new createRouter({
  history: createWebHashHistory(),
  routes: [
    { path: "/", component: () => import("../pages/Login.vue") },
    { path: "/library", component: () => import("../pages/Library.vue") },
    { path: "/upload", component: () => import("../pages/Upload.vue") },
  ],
});

export default router;
