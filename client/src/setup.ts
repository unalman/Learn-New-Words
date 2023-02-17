import "./main.ts";
import "./assets/base.css";
import { createPinia } from "pinia";
import router from "./router";
import { createApp } from "vue";
import App from "./App.vue";
import { autoAnimatePlugin } from "@formkit/auto-animate/vue";

/* import the fontawesome core */
import { library } from "@fortawesome/fontawesome-svg-core";
/* import font awesome icon component */
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
/* import specific icons */
import { faTriangleExclamation } from "@fortawesome/free-solid-svg-icons";
/* add icons to the library */
library.add(faTriangleExclamation);

export const createStore = () => {
  return createPinia();
};

export const createRouter = () => {
  return router;
};

export const createAppProcess = () => {
  const app = createApp(App);
  app.use(createStore());
  app.use(createRouter());
  app.component("font-awesome-icon", FontAwesomeIcon);
  app.use(autoAnimatePlugin);
  //app.mount("#app");
  return app;
};
