<template>
  <div class="word-top">
    <div class="word-errorDiv">
      <ErrorBlock v-if="showError" :errors="errors" />
    </div>
    <div class="word-btn">
      <CustomBtn3 @click="addItem()" :btnName="buttonNames.add" />
    </div>
  </div>
  <div class="word-container">
    <div class="word-header">
      <div class="word-firstLangHeader">Ana Dil</div>
      <div class="word-secondLangHeader">Yabancı Dil</div>
      <div style="width: 230.81px"></div>
    </div>
    <ul class="word-list" v-auto-animate>
      <li
        class="word-list-item"
        v-for="(item, index) in getLanguageData"
        v-bind:key="item.id"
      >
        <div
          class="word-item-firstLang"
          v-if="!isEdit || selectedId != item.id"
        >
          {{ item.MainLanguage }}
        </div>
        <div
          class="word-item-secondLang"
          v-if="!isEdit || selectedId != item.id"
        >
          {{ item.ForeignLanguage }}
        </div>
        <div class="word-item-firstLang" v-if="isEdit && selectedId == item.id">
          <input
            class="textInput"
            :class="validation.mainLanguage ? errorClasses.errorBorder : ''"
            id="mainLangEdit"
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
            type="text"
            v-model="item.ForeignLanguage"
          />
        </div>
        <div class="word-item-btns">
          <CustomBtn2
            class="edit"
            name="submit"
            v-on:click="editMode(item.id)"
            v-if="!isEdit || selectedId != item.id"
            :btnName="buttonNames.edit"
          />
          <CustomBtn2
            class="delete"
            v-on:click="deleteWord(index)"
            v-if="!isEdit || selectedId != item.id"
            :btnName="buttonNames.delete"
          />
          <CustomBtn2
            class="edit"
            v-on:click="updateWord()"
            v-if="isEdit && selectedId == item.id"
            :btnName="buttonNames.ok"
          />
          <CustomBtn2
            class="delete"
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
      var lang = createNewLangWord(this.wordsList as ILanguageWord[]);
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
      (this.wordsList as ILanguageWord[]).splice(index, 1);
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
  },
  computed: {
    getLanguageData(): ILanguageWord[] {
      var words = this.wordsList as ILanguageWord[];
      return words.sort(langDataDescendingSort);
    },
    showError() {
      return this.errors.length > 0;
    },
  },
});

function createNewLangWord(languageData: ILanguageWord[]): ILanguageWord {
  let newObj: ILanguageWord;
  newObj = {
    id: getLastId(languageData),
    MainLanguage: "",
    ForeignLanguage: "",
  };
  return newObj;
}
function getLastId(languageData: ILanguageWord[]): string {
  return (Math.max(...languageData.map((x) => parseInt(x.id))) + 1).toString();
}
function langDataDescendingSort(
  model1: ILanguageWord,
  model2: ILanguageWord
): number {
  return parseInt(model2.id) - parseInt(model1.id);
}
</script>
<style scoped>
.word-btn {
  padding-top: 1.5rem;
  padding-right: 4rem;
  text-align: end;
}
.word-container {
  margin-top: 1rem;
  box-shadow: 0 2px 6px 0 hsla(0, 0%, 0%, 0.2);
}
.word-header {
  display: flex;
  align-items: center;
  padding-top: 0.8rem;
  padding-bottom: 0.8rem;
  padding-left: 1.5rem;
  justify-content: space-between;
  font-size: 20px;
  font-weight: 600;
  color: #4a4a4a;
  background-color: var(--lightBlue);
}
.word-list-item {
  display: flex;
  font-size: 20px;
  color: #4a4a4a;
  align-items: center;
  padding-left: 1.5rem;
  border-top: 1px solid var(--lightGreyTableBorder);
  justify-content: space-between;
}
.word-list-noresult {
  color: #4a4a4a;
  font-size: 20px;
  padding: 10px;
}
.word-item-btns {
  padding-right: 5rem;
  padding-left: 0.7rem;
}
.textInput {
  border: 1px solid var(--lightGreyTableBorder);
  color: #4a4a4a;
  border-radius: 0.3rem;
  font-size: 20px;
  outline: none;
}
.textInput:focus {
  border: 0.13rem solid var(--InputFocusGrey);
}
.word-top {
  display: flex;
  justify-content: space-between;
}
.errorBorder {
  border-color: #ff6868;
}
.word-errorDiv {
  max-width: 80%;
}
</style>
