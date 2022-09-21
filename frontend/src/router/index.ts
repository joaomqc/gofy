import Vue from "vue";
import VueRouter, { RouteConfig } from "vue-router";
import MonitorView from "@/views/monitor/MonitorView.vue";

Vue.use(VueRouter);

const routes: Array<RouteConfig> = [
  {
    path: "/monitor",
    name: "monitor",
    component: MonitorView,
  },
  {
    path: "/search",
    name: "search",
    component: () => import("@/views/search/SearchView.vue"),
  },
  {
    path: "/settings",
    name: "settings",
    component: () => import("@/views/settings/SettingsView.vue"),
  },
  {
    path: "*",
    redirect: "/monitor",
  },
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
});

export default router;
