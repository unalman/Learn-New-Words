import { createRouter, createWebHistory } from "vue-router";
import HomePage from "../views/HomePage.vue";
import AddWord from "../views/AddWord.vue";
import WordTest from "../views/WordTest.vue";
import WordTestResult from "../views/WordTestResult.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "Home",
      component: HomePage,
    },
    {
      path: "/addword",
      name: "AddWord",
      component: AddWord,
    },
    {
      path: "/wordtest",
      name: "WordTest",
      component: WordTest,
    },
    {
      path: "/wordtestresult",
      name: "WordTestResult",
      component: WordTestResult,
      props: true,
    },
  ],
});

export default router;
