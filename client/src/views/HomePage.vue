<template>
  <div class="content-container">
    <div class="content-item1">
      <div>{{ t("homeTitleText") }}</div>
      <div class="content-btnContainer">
        <CustomBtn1
          data-cy="addWord"
          :btnName="t('addWord')"
          :href="data.btnAddWord.href"
        />
      </div>
    </div>
    <div class="content-item2">
      <div class="item2-testItem">{{ t("testYourself") }}</div>
      <div class="item2-testItem content-btnContainer">
        <CustomBtn1
          data-cy="wordTest"
          :btnName="t('wordTest')"
          @click="CheckAndNavigateToWordTest()"
        />
      </div>
      <div class="item2-testItem">
        <ErrorBlock v-if="showError" :errors="data.error" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { IHomePageVM } from "typings/interface/IHomePageWord";
import CustomBtn1 from "../components/buttons/CustomBtn1.vue";
import ErrorBlock from "../components/ErrorBlock.vue";
import { computed, onMounted, ref } from "vue";
import { useI18n } from "vue-i18n";
import { useRouter } from "vue-router";
import { useLanguageWordStore } from "../store/index";
const { t } = useI18n();
const router = useRouter();
const languageWordStore = useLanguageWordStore();
onMounted(async () => {
  const { success, error } =
    await languageWordStore.dispatchGetRandomLanguageWords();
  if (!success) {
    console.log(error);
  }
});

var data = ref<IHomePageVM>({
  btnAddWord: {
    href: "/addword",
  },
  btnTestWord: {
    name: "WordTest",
    isDisabled: false,
  },
  error: ["minimumFiveWord"],
});
const showError = computed(() => {
  return languageWordStore.randomWords.length < 5;
});
function CheckAndNavigateToWordTest() {
  if (languageWordStore.randomWords.length >= 5) {
    router.push({
      name: data.value.btnTestWord.name,
    });
  } else {
    data.value.btnTestWord.isDisabled = true;
  }
}
</script>

<style scoped>
.content-container {
  display: grid;
  padding: 2vh;
  grid-gap: 10px;
  height: 100%;
  min-height: inherit;
  align-items: center;
  justify-content: space-around;
  color: darkslategray;
}

.content-container > .content-item1,
.content-item2 {
  text-align: center;
  font-size: 1.7rem;
  box-sizing: border-box;
  width: 40vh;
  border: solid 0.4px;
  border-color: var(--lightGrey);
  min-width: 180px;
  height: 14vh;
  border-radius: 1rem;
  padding-top: 1rem;
  background-color: var(--lightBlue);
}
.content-item1 {
  grid-template-columns: 100px auto;
  min-height: 144px;
}
.content-item2 {
  grid-row-start: 2;
  grid-column-start: 2;
  min-height: 101px;
}

@media screen and (max-width: 800px) {
  .content-item2 {
    grid-row-start: 2;
    grid-column-start: 1;
    min-height: 101px;
  }
}
.content-btnContainer {
  /* padding-top: 1rem; */
  text-align: center;
}
.item2-testItem {
  display: flex;
  justify-content: center;
}
</style>
