<template>
  <div class="word-top" data-testId="WordsTable-TestId">
    <div class="word-errorDiv">
      <ErrorBlock v-if="showError" :errors="data.errors" />
    </div>
    <div class="word-btn">
      <CustomBtn3 @click="addItem()" data-cy="addbtn" :btnName="t('add')" />
    </div>
  </div>
  <div class="word-container">
    <div class="word-header">
      <div class="word-firstLangHeader">
        {{ t("firstField") }}
      </div>
      <div class="word-secondLangHeader">
        {{ t("secondField") }}
      </div>
      <div style="width: 230.81px"></div>
    </div>
    <ul class="word-list" ref="parent">
      <li
        class="word-list-item"
        data-cy="languageItem"
        v-for="(item, index) in getLanguageData"
        v-bind:key="item.id"
      >
        <div
          class="word-item-firstLang"
          data-cy="mainLang"
          v-if="!data.isEdit || data.selectedId != item.id"
        >
          {{ item.MainLanguage }}
        </div>
        <div
          class="word-item-secondLang"
          data-cy="foreignLang"
          v-if="!data.isEdit || data.selectedId != item.id"
        >
          {{ item.ForeignLanguage }}
        </div>
        <div
          class="word-item-firstLang"
          v-if="data.isEdit && data.selectedId == item.id"
        >
          <input
            class="textInput"
            :class="
              data.validation.mainLanguage ? data.errorClasses.errorBorder : ''
            "
            id="mainLangEdit"
            data-cy="mainLangEdit"
            type="text"
            v-model="item.MainLanguage"
          />
        </div>
        <div
          class="word-item-secondLang"
          v-if="data.isEdit && data.selectedId == item.id"
        >
          <input
            class="textInput"
            :class="
              data.validation.foreignLanguage
                ? data.errorClasses.errorBorder
                : ''
            "
            id="foreignLangEdit"
            data-cy="foreignLangEdit"
            type="text"
            v-model="item.ForeignLanguage"
          />
        </div>
        <div class="word-item-btns">
          <CustomBtn2
            class="edit"
            name="submit"
            data-cy="editbtn"
            v-on:click="editMode(item.id)"
            v-if="!data.isEdit || data.selectedId != item.id"
            :btnName="t('edit')"
          />
          <CustomBtn2
            class="delete"
            data-cy="deletebtn"
            v-on:click="deleteWord(index)"
            v-if="!data.isEdit || data.selectedId != item.id"
            :btnName="t('delete')"
          />
          <CustomBtn2
            class="edit"
            data-cy="okbtn"
            v-on:click="updateWord()"
            v-if="data.isEdit && data.selectedId == item.id"
            :btnName="t('ok')"
          />
          <CustomBtn2
            class="delete"
            data-cy="cancelbtn"
            v-on:click="cancelWord(index)"
            v-if="data.isEdit && data.selectedId == item.id"
            :btnName="t('cancel')"
          />
        </div>
      </li>
      <li
        class="word-list-noresult"
        data-testId="noresult"
        v-if="data.wordsList.length == 0"
      >
        {{ t("noResult") }}
      </li>
    </ul>
  </div>
</template>
<script setup lang="ts">
import { useAutoAnimate } from "@formkit/auto-animate/vue";
import { computed, ref } from "vue";
import CustomBtn2 from "./buttons/CustomBtn2.vue";
import CustomBtn3 from "./buttons/CustomBtn3.vue";
import ErrorBlock from "./ErrorBlock.vue";
import LanguageWord from "../../typings/classes/LanguageWord";
import { textValidationAndTrim } from "../libs/common";
import { getLastId, sortLanguageWordDescending } from "../libs/WordsTable";

import type { IWordTable } from "typings/interface/IWordTable";
import type { ILanguageWord } from "typings/interface/ILanguageWord";
import type { PropType } from "vue";
import { useI18n } from "vue-i18n";
const { t } = useI18n();

const [parent] = useAutoAnimate();
const props = defineProps({
  languageData: {
    type: Array as PropType<ILanguageWord[]>,
  },
});

const data = ref<IWordTable>({
  wordsList: props.languageData as ILanguageWord[],
  validation: {
    mainLanguage: false,
    foreignLanguage: false,
  },
  errorClasses: {
    errorBorder: "errorBorder",
  },
  errors: [],
  isEdit: false,
  selectedId: "",
  newItem: false,
  previousValues: {
    id: "",
    MainLanguage: "",
    ForeignLanguage: "",
  },
});

const addItem = (): void => {
  if (data.value.isEdit) return;
  data.value.newItem = true;
  var lang = createNewLangWord(data.value.wordsList as ILanguageWord[]);
  (data.value.wordsList as ILanguageWord[]).unshift(lang);
  editMode(lang.id);
};
const editMode = (id: string): void => {
  if (data.value.isEdit) return;
  data.value.selectedId = id;
  data.value.previousValues = Object.assign(
    {},
    findLanguagebyId(id) as ILanguageWord
  );
  data.value.isEdit = true;
};
const deleteWord = (index: number): void => {
  (data.value.wordsList as ILanguageWord[]).splice(index, 1);
};
const updateWord = (): void => {
  data.value.isEdit = !isValid();
  data.value.newItem = false;
};
const cancelWord = (index: number): void => {
  updatePreviousLanguage();
  if (data.value.newItem) {
    (data.value.wordsList as ILanguageWord[]).splice(index, 1);
  }
  data.value.isEdit = false;
  data.value.newItem = false;
  data.value.errors = [];
  data.value.validation.mainLanguage = false;
  data.value.validation.foreignLanguage = false;
};
const isValid = (): boolean => {
  var isSuccess = true;
  const mainLanguageInput = document.getElementById(
    "mainLangEdit"
  ) as HTMLInputElement | null;
  const foreignLanguageInput = document.getElementById(
    "foreignLangEdit"
  ) as HTMLInputElement | null;
  data.value.errors = [];
  if (!textValidationAndTrim(mainLanguageInput?.value as string)) {
    data.value.errors.push(
      t("mainLanguageText", { firstField: t("firstField") })
    );
    data.value.validation.mainLanguage = true;
    isSuccess = false;
  }
  if (!textValidationAndTrim(foreignLanguageInput?.value as string)) {
    data.value.errors.push(
      t("foreignLanguageText", {
        secondField: t("secondField"),
      })
    );
    data.value.validation.foreignLanguage = true;
    isSuccess = false;
  }
  return isSuccess;
};
const updatePreviousLanguage = (): void => {
  var words = data.value.wordsList?.map((obj: ILanguageWord) => {
    if (obj.id == data.value.selectedId) {
      return (obj = data.value.previousValues);
    }
    return obj;
  });
  data.value.wordsList = words;
};
const findLanguagebyId = (id: string): ILanguageWord | null => {
  return data.value.wordsList?.find(
    (x: ILanguageWord) => x.id == id
  ) as ILanguageWord | null;
};
const createNewLangWord = (languageData: ILanguageWord[]): ILanguageWord => {
  let newObj: ILanguageWord = LanguageWord.createInstance(
    getLastId(languageData)
  );
  return newObj;
};
const getLanguageData = computed((): ILanguageWord[] => {
  var words = data.value.wordsList as ILanguageWord[];
  return sortLanguageWordDescending(words);
});
const showError = computed(() => {
  return data.value.errors.length > 0;
});
</script>
<style scoped>
@import "../assets/wordsTable.css";
</style>
