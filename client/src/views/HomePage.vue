<template>
  <div class="content-container">
    <div class="content-item1">
      <div>{{ t("homeTitleText") }}</div>
      <div class="content-btnContainer">
        <CustomBtn1 :btnName="t('addWord')" :href="data.btnAddWord.href" />
      </div>
    </div>
    <div class="content-item2">
      <div>{{ t("testYourself") }}</div>
      <div class="content-btnContainer">
        <CustomBtn1 :btnName="t('wordTest')" :href="data.btnTestWord.href" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { IHomePageWord } from "typings/interface/IHomePageWord";
import CustomBtn1 from "../components/buttons/CustomBtn1.vue";
import { ref } from "vue";
import { useI18n } from "vue-i18n";
const { t } = useI18n();

var data = ref<IHomePageWord>({
  btnAddWord: {
    href: "/addword",
  },
  btnTestWord: {
    href: "/wordtest",
  },
});
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
  padding-top: 1rem;
  text-align: center;
}
</style>
