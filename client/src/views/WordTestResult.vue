<template>
  <div class="wordTestResult-container">
    <div class="wordTestResult-ResultContainer">
      <div class="wordTestResult-ResultText">
        {{ data.congratulationsText }}
      </div>
      <div class="wordTestResult-TableContainer">
        <div class="wordTestResult-WordsListText">{{ data.words }}</div>
        <div class="wordTestResult-WordList">
          <div
            class="wordTestResult-WordBox"
            v-for="(item, index) in data.wordTestReturns"
            v-bind:key="item.id"
          >
            <span>{{ item.ForeignLanguage }}: </span>
            <span>{{ item.MainLanguage }}</span>
            <span v-if="item.isCorrectAnswer"> - 👍</span>
            <span v-if="!item.isCorrectAnswer"> - 👎</span>
          </div>
        </div>
      </div>
      <div>
        <CustomBtn4
          v-on:click="ReturnAddWordPage()"
          :btnName="data.button.complate"
        />
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import CustomBtn4 from "@/components/buttons/CustomBtn4.vue";
import type { IWordTestResult } from "typings/interface/IWordTestResult";
import type { IWordTestReturn } from "typings/interface/IWordTest";
import { ref } from "vue";
import { useRoute, useRouter } from "vue-router";

const route = useRoute();
const router = useRouter();
var jsonData: Array<IWordTestReturn>;
if (route.query.wordTestReturns != null) {
  jsonData = JSON.parse(route.query.wordTestReturns as string);
  console.log(jsonData);
}
const data = ref<IWordTestResult>({
  wordTestReturns:
    route.query.wordTestReturns != null
      ? JSON.parse(route.query.wordTestReturns as string)
      : Array<IWordTestReturn>,
  congratulationsText: "🎉 Tebrikler Testi Başarıyla Tamamladınız 🎉",
  words: "Kelimeler",
  button: {
    complate: "Tamamla",
  },
});
function ReturnAddWordPage() {
  router.push({
    name: "AddWord",
  });
}
</script>
<style scoped>
.wordTestResult-container {
  display: grid;
  grid-template-columns: 1fr 1fr 8fr 1fr 1fr;
  height: 100%;
  min-height: inherit;
  overflow: overlay;
}
.wordTestResult-ResultContainer {
  grid-column: 3;
  display: flex;
  text-align: center;
  flex-direction: column;
  margin-top: 10vh;
}
.wordTestResult-ResultText {
  padding: 1rem;
  font-size: 2rem;
  padding-bottom: 2rem;
  font-weight: 600;
  color: #4a4a4a;
}
.wordTestResult-TableContainer {
  box-shadow: 0 2px 6px 0 hsla(0, 0%, 0%, 0.2);
  margin-bottom: 2rem;
}
.wordTestResult-WordsListText {
  background-color: var(--lightBlue);
  padding: 1rem;
  font-size: 1.7rem;
  font-weight: 600;
  color: #4a4a4a;
}
.wordTestResult-WordList {
  display: flex;
  flex-direction: column;
  font-size: 1.5rem;
  color: #4a4a4a;
}
.wordTestResult-WordBox {
  border-top: 1px solid var(--lightGreyTableBorder);
  padding: 0.5rem;
}
</style>
