<template>
  <div class="wordTest-container">
    <div class="wordTest-mainItem">
      <div class="wordTest-Question">
        {{ data.languageWords[answerCount].ForeignLanguage }}
      </div>
      <div class="wordTest-Answer">
        <input
          type="text"
          v-model="text"
          class="answerText"
          placeholder="Cevap"
        />
      </div>
      <div class="wordTest-btnList">
        <CustomBtn4 :btnName="data.buttonNames.idk" />
        <CustomBtn4
          v-if="answerCount != 4 && !isAnswered"
          v-on:click="updateResult()"
          :btnName="data.buttonNames.check"
        />
        <CustomBtn4
          v-if="answerCount != 4 && isAnswered"
          v-on:click="nextQuestion()"
          :btnName="data.buttonNames.continue"
        />
        <CustomBtn4
          v-if="answerCount == 4"
          v-on:click="routeResultPage()"
          :btnName="data.buttonNames.result"
        />
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import CustomBtn4 from "../components/buttons/CustomBtn4.vue";
import { ref } from "vue";
import type { IWordTest, IWordTestReturn } from "typings/interface/IWordTest";
import type { ILanguageWord } from "typings/interface/ILanguageWord";
import { useLangueageWordStore } from "@/stores/LanguageWord";
import { useRoute, useRouter } from "vue-router";

const langWordStore = useLangueageWordStore();
const data = ref<IWordTest>({
  languageWords: langWordStore.list,
  buttonNames: {
    check: "Kontrol et",
    idk: "Bilmiyorum",
    continue: "Sonraki",
    result: "Sonuç",
  },
});
var correctAnswer = 0,
  wrongAnswer = 0,
  answerCount = 0;
var isAnswered = ref<Boolean>(false);
var wordTestReturns = Array<IWordTestReturn>();
const text = ref<string>("");
const updateResult = () => {
  if (answerCount != 4 && text.value.trim() != "") {
    var isCorrect = false;
    if (
      data.value.languageWords[answerCount].MainLanguage.toLowerCase() ==
      text.value.toLowerCase()
    ) {
      isCorrect = true;
    }
    var wordTestReturn: IWordTestReturn = {
      id: data.value.languageWords[answerCount].id,
      MainLanguage: data.value.languageWords[answerCount].MainLanguage,
      ForeignLanguage: data.value.languageWords[answerCount].ForeignLanguage,
      isCorrectAnswer: isCorrect,
    };
    wordTestReturns.push(wordTestReturn);
    isAnswered.value = true;
    answerCount++;
  }
};
const nextQuestion = () => {
  isAnswered.value = false;
};
const router = useRouter();
function routeResultPage() {
  var jsonArray = JSON.stringify(wordTestReturns);
  router.push({
    name: "WordTestResult",
    query: { wordTestReturns: jsonArray }
  });
}
</script>
<style scoped>
.wordTest-container {
  display: grid;
  height: 100%;
  min-height: inherit;
  grid-template-columns: 1fr 1fr 2fr 1fr 1fr;
  overflow: overlay;
}
.wordTest-mainItem {
  grid-column: 3;
  margin-top: 14vh;
}
.wordTest-Question {
  font-size: 3rem;
  margin-bottom: 4rem;
  color: darkslategray;
}
.wordTest-Answer {
  margin-bottom: 1rem;
  height: 50px;
}
.answerText {
  border-radius: 0.2rem;
  border: 1px solid var(--lightGreyTableBorder);
  font-size: inherit;
  outline: none;
}
.answerText:focus {
  border: 1px solid var(--InputFocusGrey);
}
.wordTest-btnList {
  margin-top: 1rem;
  display: flex;
  justify-content: space-evenly;
}
</style>
