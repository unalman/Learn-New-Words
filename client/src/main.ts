import { createApp } from "vue";
import { createPinia } from "pinia";

import App from "./App.vue";
import router from "./router";

import WordsTable from "./components/WordsTable.vue";
import CustomBtn1 from "./components/buttons/CustomBtn1.vue";
import CustomBtn2 from "./components/buttons/CustomBtn2.vue";
import CustomBtn3 from "./components/buttons/CustomBtn3.vue";

const app = createApp(App);

app.use(createPinia());
app.use(router);
app.component("WordsTable", WordsTable);
app.component("CustomBtn1", CustomBtn1);
app.component("CustomBtn2", CustomBtn2);
app.component("CustomBtn3", CustomBtn3);

app.mount("#app");
