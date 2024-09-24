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
            v-on:click="languageWordStore.dispatchDeleteWord(item.id)"
            v-if="!data.isEdit || data.selectedId != item.id"
            :btnName="t('delete')"
          />
          <CustomBtn2
            class="edit"
            data-cy="okbtn"
            v-on:click="updateWord(item.id)"
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
        v-if="languageWordStore.words.length == 0"
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
import { sortLanguageWordDescending } from "../libs/WordsTable";

import type { IWordTable } from "typings/interface/IWordTable";
import type { ILanguageWord } from "typings/interface/ILanguageWord";
import { useLanguageWordStore } from "@/store";
import { useI18n } from "vue-i18n";
const { t } = useI18n();

const [parent] = useAutoAnimate();
const languageWordStore = useLanguageWordStore();

const data = ref<IWordTable>({
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
    created_at: new Date(),
  },
});

const addItem = (): void => {
  if (data.value.isEdit) return;
  data.value.newItem = true;
  var newObj = LanguageWord.createInstance();
  languageWordStore.words.unshift(newObj);
  editMode(newObj.id);
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
const updateWord = async (id: string) => {
  const isValid = modelValidation();
  if (isValid) {
    const newModel = findLanguagebyId(id);
    if (newModel != null) {
      if (data.value.newItem) {
        await languageWordStore.dispatchCreateWord(newModel);
      } else {
        await languageWordStore.dispatchUpdateWord(newModel);
      }
    }
  }
  data.value.isEdit = !isValid;
  data.value.newItem = false;
};
const cancelWord = (index: number): void => {
  updatePreviousLanguage();
  if (data.value.newItem) {
    languageWordStore.words.splice(index, 1);
  }
  data.value.isEdit = false;
  data.value.newItem = false;
  data.value.errors = [];
  data.value.validation.mainLanguage = false;
  data.value.validation.foreignLanguage = false;
};
const modelValidation = (): boolean => {
  var isSuccess = true;
  const mainLanguageInput = document.getElementById(
    "mainLangEdit"
  ) as HTMLInputElement | null;
  const foreignLanguageInput = document.getElementById(
    "foreignLangEdit"
  ) as HTMLInputElement | null;
  data.value.errors = [];
  if (!textValidationAndTrim(mainLanguageInput?.value as string)) {
    data.value.errors.push("mainLanguageText");
    data.value.validation.mainLanguage = true;
    isSuccess = false;
  }
  if (!textValidationAndTrim(foreignLanguageInput?.value as string)) {
    data.value.errors.push("foreignLanguageText");
    data.value.validation.foreignLanguage = true;
    isSuccess = false;
  }
  return isSuccess;
};
const updatePreviousLanguage = (): void => {
  var words = languageWordStore.words?.map((obj: ILanguageWord) => {
    if (obj.id == data.value.selectedId) {
      return (obj = data.value.previousValues);
    }
    return obj;
  });
  languageWordStore.words = words;
};
const findLanguagebyId = (id: string): ILanguageWord | null => {
  return languageWordStore.words?.find(
    (x: ILanguageWord) => x.id == id
  ) as ILanguageWord | null;
};
const getLanguageData = computed((): ILanguageWord[] => {
  return sortLanguageWordDescending(languageWordStore.words);
});
const showError = computed(() => {
  return data.value.errors.length > 0;
});
</script>
<style scoped>
@import "../assets/wordsTable.css";
</style>
