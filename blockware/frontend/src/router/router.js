import { createRouter, createWebHashHistory } from "vue-router";

const router = new createRouter({
  history: createWebHashHistory(),
  routes: [
    { path: "/", component: () => import("../pages/Login.vue") },
    { path: "/home", component: () => import("../pages/Home.vue") },
    { path: "/upload", component: () => import("../pages/Upload.vue") },
    { path: "/library", component: () => import("../pages/Library.vue") },
    { path: "/store", component: () => import("../pages/Store.vue") },
    { path: "/downloads", component: () => import("../pages/Downloads.vue") },
  ],
});

export default router;
