<template>
  <div
    class="wordTest-container"
    :class="{ wordTestContainerFirstSize: !isAnswered }"
  >
    <div class="wordTest-mainItem">
      <div data-cy="wordTestQuestion" class="wordTest-Question">
        {{ languageWordStore.randomWords[answerCount]?.ForeignLanguage }}
      </div>
      <div class="wordTest-Answer">
        <input
          type="text"
          data-cy="answerText"
          v-model="text"
          class="answerText"
          :placeholder="t('placeholderText')"
        />
      </div>
      <div class="wordTest-btnList">
        <CustomBtn4
          :class="{
            disabledBtn: isAnswered,
          }"
          :btnName="t('idk')"
          data-cy="dontKnowBtn"
          v-on:click="dontKnow()"
          :disabled="isAnswered"
        />
        <CustomBtn4
          v-if="answerCount != 4 && !isAnswered"
          v-on:click="updateResult()"
          data-cy="checkBtn"
          :btnName="t('check')"
        />
        <CustomBtn4
          v-if="answerCount != 4 && isAnswered"
          data-cy="nextBtn"
          v-on:click="nextQuestion()"
          :btnName="t('continue')"
        />
        <CustomBtn4
          v-if="answerCount == 4"
          v-on:click="routeResultPage()"
          :btnName="t('result')"
        />
      </div>
    </div>
  </div>
  <div
    class="wordTest-Result"
    data-cy="wordTestResult"
    :class="{
      wordTestResultPassive: !isAnswered,
      wordTestResultCorrect: isCorrect,
      wordTestResultWrong: !isCorrect,
    }"
  >
    {{
      isCorrect
        ? t("correct")
        : `${t("correct")}: ${
            languageWordStore.randomWords[answerCount]?.MainLanguage
          }`
    }}
  </div>
</template>
<script setup lang="ts">
import CustomBtn4 from "../components/buttons/CustomBtn4.vue";
import { ref } from "vue";
import type { IWordTestReturn } from "typings/interface/IWordTest";
import { useRouter } from "vue-router";
import { useI18n } from "vue-i18n";
import { useLanguageWordStore } from "@/store";
const { t } = useI18n();

const languageWordStore = useLanguageWordStore();

var answerCount = 0;
var isAnswered = ref<Boolean>(false);
var isCorrect = ref<Boolean>(false);

var wordTestReturns = Array<IWordTestReturn>();
const text = ref<string>("");

const updateResult = () => {
  if (answerCount != parseInt(import.meta.env.VITE_Word_Test_Count) && text.value.trim() != "") {
    isCorrect.value = false;
    if (
      languageWordStore.randomWords[answerCount].MainLanguage.toLowerCase() ==
      text.value.toLowerCase()
    ) {
      isCorrect.value = true;
    }
    var wordTestReturn: IWordTestReturn = {
      id: languageWordStore.randomWords[answerCount].id,
      MainLanguage: languageWordStore.randomWords[answerCount].MainLanguage,
      ForeignLanguage:
        languageWordStore.randomWords[answerCount].ForeignLanguage,
      isCorrectAnswer: isCorrect.value,
      created_at: new Date(),
    };
    wordTestReturns.push(wordTestReturn);
    isAnswered.value = true;
  }
};
const dontKnow = () => {
  isCorrect.value = false;
  isAnswered.value = true;
  var wordTestReturn: IWordTestReturn = {
    id: languageWordStore.randomWords[answerCount].id,
    MainLanguage: languageWordStore.randomWords[answerCount].MainLanguage,
    ForeignLanguage: languageWordStore.randomWords[answerCount].ForeignLanguage,
    isCorrectAnswer: isCorrect.value,
    created_at: new Date(),
  };
  wordTestReturns.push(wordTestReturn);
};
const nextQuestion = () => {
  isAnswered.value = false;
  answerCount++;
  text.value = "";
};
const router = useRouter();
function routeResultPage() {
  var jsonArray = JSON.stringify(wordTestReturns);
  router.push({
    name: "WordTestResult",
    query: { wordTestReturns: jsonArray },
  });
}
</script>
<style scoped>
.wordTest-container {
  display: grid;
  height: 93%;
  min-height: inherit;
  grid-template-columns: 1fr 1fr 2fr 1fr 1fr;
  overflow: overlay;
}
.wordTestContainerFirstSize {
  height: 100%;
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
  text-align: center;
}
.answerText:focus {
  border: 1px solid var(--InputFocusGrey);
}
.wordTest-btnList {
  margin-top: 1rem;
  display: flex;
  justify-content: space-evenly;
}
.wordTest-Result {
  height: 7%;
  padding-top: 1.3vh;
}
.wordTestResultPassive {
  display: none;
  height: 0%;
}
.wordTestResultCorrect {
  background: radial-gradient(
    circle,
    rgba(187, 226, 255, 1) 0%,
    rgba(123, 212, 119, 1) 20%,
    rgba(187, 226, 255, 1) 47%,
    rgba(187, 226, 255, 1) 100%
  );
}
.wordTestResultWrong {
  background: radial-gradient(
    circle,
    rgba(187, 226, 255, 1) 0%,
    rgba(255, 146, 146, 1) 20%,
    rgba(187, 226, 255, 1) 47%,
    rgba(187, 226, 255, 1) 100%
  );
}
</style>
