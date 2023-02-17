<template>
  <div class="word-top">
    <div class="word-errorDiv">
      <ErrorBlock v-if="showError" :errors="errors" />
    </div>
    <div class="word-btn">
      <CustomBtn3
        @click="addItem()"
        data-cy="addbtn"
        :btnName="buttonNames.add"
      />
    </div>
  </div>
  <div class="word-container">
    <div class="word-header">
      <div class="word-firstLangHeader">Ana Dil</div>
      <div class="word-secondLangHeader">Yabancı Dil</div>
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
          v-if="!isEdit || selectedId != item.id"
        >
          {{ item.MainLanguage }}
        </div>
        <div
          class="word-item-secondLang"
          data-cy="foreignLang"
          v-if="!isEdit || selectedId != item.id"
        >
          {{ item.ForeignLanguage }}
        </div>
        <div class="word-item-firstLang" v-if="isEdit && selectedId == item.id">
          <input
            class="textInput"
            :class="validation.mainLanguage ? errorClasses.errorBorder : ''"
            id="mainLangEdit"
            data-cy="mainLangEdit"
            type="text"
            v-model="item.MainLanguage"
          />
        </div>
        <div
          class="word-item-secondLang"
          v-if="isEdit && selectedId == item.id"
        >
          <input
            class="textInput"
            :class="validation.foreignLanguage ? errorClasses.errorBorder : ''"
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
            v-if="!isEdit || selectedId != item.id"
            :btnName="buttonNames.edit"
          />
          <CustomBtn2
            class="delete"
            data-cy="deletebtn"
            v-on:click="deleteWord(index)"
            v-if="!isEdit || selectedId != item.id"
            :btnName="buttonNames.delete"
          />
          <CustomBtn2
            class="edit"
            data-cy="okbtn"
            v-on:click="updateWord()"
            v-if="isEdit && selectedId == item.id"
            :btnName="buttonNames.ok"
          />
          <CustomBtn2
            class="delete"
            data-cy="cancelbtn"
            v-on:click="cancelWord(index)"
            v-if="isEdit && selectedId == item.id"
            :btnName="buttonNames.cancel"
          />
        </div>
      </li>
      <li class="word-list-noresult" v-if="wordsList.length == 0">
        {{ errorClasses.noResult }}
      </li>
    </ul>
  </div>
</template>
<script setup lang="ts">
import { useAutoAnimate } from "@formkit/auto-animate/vue";

const [parent] = useAutoAnimate();
</script>
<script lang="ts">
import CustomBtn2 from "./buttons/CustomBtn2.vue";
import CustomBtn3 from "./buttons/CustomBtn3.vue";
import ErrorBlock from "./ErrorBlock.vue";
import type { IWordTable, ILanguageWord } from "typings/interface/IWordTable";
import { defineComponent } from "vue";
import type { PropType } from "vue";

export default defineComponent({
  props: {
    languageData: {
      type: Array as PropType<ILanguageWord[]>,
      required: true,
    },
  },
  data(): IWordTable {
    return {
      wordsList: this.languageData as ILanguageWord[],
      buttonNames: {
        add: "Ekle",
        edit: "Düzenle",
        delete: "Sil",
        ok: "Tamam",
        cancel: "İptal",
      },
      validation: {
        mainLanguage: false,
        mainLanguageText: "Ana Dil'e değer girilmedi",
        foreignLanguage: false,
        foreignLanguageText: "Yabancı Dil'e değer girilmedi",
      },
      errorClasses: {
        errorBorder: "errorBorder",
        noResult: "Herhangi bir sonuç yok",
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
    };
  },
  components: { CustomBtn2, CustomBtn3, ErrorBlock },
  methods: {
    addItem(): void {
      if (this.isEdit) return;
      this.newItem = true;
      var lang = this.createNewLangWord(this.wordsList as ILanguageWord[]);
      (this.wordsList as ILanguageWord[]).unshift(lang);
      this.editMode(lang.id);
    },
    editMode(id: string): void {
      if (this.isEdit) return;
      this.selectedId = id;
      this.previousValues = Object.assign(
        {},
        this.findLanguagebyId(id) as ILanguageWord
      );
      this.isEdit = true;
    },
    deleteWord(index: number): void {
      (this.wordsList as ILanguageWord[]).splice(index, 1);
    },
    updateWord(): void {
      this.isEdit = !this.isValid();
      this.newItem = false;
    },
    cancelWord(index: number): void {
      this.updatePreviousLanguage();
      if (this.newItem) {
        (this.wordsList as ILanguageWord[]).splice(index, 1);
      }
      this.isEdit = false;
      this.newItem = false;
      this.errors = [];
      this.validation.mainLanguage = false;
      this.validation.foreignLanguage = false;
    },
    isValid(): boolean {
      var isSuccess = true;
      const mainLanguageInput = document.getElementById(
        "mainLangEdit"
      ) as HTMLInputElement | null;
      const foreignLanguageInput = document.getElementById(
        "foreignLangEdit"
      ) as HTMLInputElement | null;
      this.errors = [];
      if (!this.textValidation(mainLanguageInput?.value as string)) {
        this.errors.push(this.validation.mainLanguageText);
        this.validation.mainLanguage = true;
        isSuccess = false;
      }
      if (!this.textValidation(foreignLanguageInput?.value as string)) {
        this.errors.push(this.validation.foreignLanguageText);
        this.validation.foreignLanguage = true;
        isSuccess = false;
      }
      return isSuccess;
    },
    textValidation(text: string): boolean {
      var trimmedValue = text?.replace(/\s/g, "");
      if (text === null || trimmedValue == "") {
        return false;
      }
      return true;
    },
    updatePreviousLanguage(): void {
      var words = this.wordsList?.map((obj: ILanguageWord) => {
        if (obj.id == this.selectedId) {
          return (obj = this.previousValues);
        }
        return obj;
      });
      this.wordsList = words;
    },
    findLanguagebyId(id: string): ILanguageWord | null {
      return this.wordsList?.find(
        (x: ILanguageWord) => x.id == id
      ) as ILanguageWord | null;
    },
    getLastId(languageData: ILanguageWord[]): string {
      return (
        Math.max(...languageData.map((x) => parseInt(x.id))) + 1
      ).toString();
    },
    createNewLangWord(languageData: ILanguageWord[]): ILanguageWord {
      let newObj: ILanguageWord;
      newObj = {
        id: this.getLastId(languageData),
        MainLanguage: "",
        ForeignLanguage: "",
      };
      return newObj;
    },
  },
  computed: {
    getLanguageData(): ILanguageWord[] {
      var words = this.wordsList as ILanguageWord[];
      return words.sort((a, b) => (parseInt(a.id) > parseInt(b.id) ? -1 : 1));
    },
    showError() {
      return this.errors.length > 0;
    },
  },
});
</script>
<style scoped>
@import "../assets/wordsTable.css";
</style>
