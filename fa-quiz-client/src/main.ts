import { createApp } from "vue";

import App from "./components/App.vue";

import "./assets/style.scss";
import "./assets/color.scss";

import { library } from "@fortawesome/fontawesome-svg-core";
import { fab } from "@fortawesome/free-brands-svg-icons";
import { far } from "@fortawesome/free-regular-svg-icons";
import { fas } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import { createRouter, createWebHistory } from "vue-router";

const router = createRouter({
    history: createWebHistory(),
    routes: App.routes,
})

library.add(fab, far, fas);

createApp(App).component("font-awesome-icon", FontAwesomeIcon).use(router).mount("#app");
