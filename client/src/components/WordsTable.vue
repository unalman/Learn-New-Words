<template>
  <div class="word-top">
    <ul class="word-error">
      <li class="word-error-item" v-for="item in errors" :key="item">
        {{ item }}
      </li>
    </ul>
    <div class="word-btn">
      <CustomBtn3 @click="addItem()" :btnName="buttonNames.add" />
    </div>
  </div>
  <div class="word-container">
    <div class="word-header">
      <div class="word-firstLangHeader">Ana Dil</div>
      <div class="word-secondLangHeader">Yabancı Dil</div>
    </div>
    <ul class="word-list" v-auto-animate>
      <li
        class="word-list-item"
        v-for="(item, index) in getLanguageData()"
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
    </ul>
    <div class="last-item-bottom"></div>
  </div>
</template>
<script lang="ts">
import CustomBtn2 from "./buttons/CustomBtn2.vue";
import CustomBtn3 from "./buttons/CustomBtn3.vue";
import type { IWordTable, ILanguageWord } from "typings/interface/IWordTable";
export default {
  props: {
    languageData: {
      type: Array as () => ILanguageWord[],
    },
  },
  data(): IWordTable {
    return {
      wordsList: this.languageData as ILanguageWord[],
      buttonNames: {
        add: "Ekle",
        edit: "Düzenle",
        delete: "Sil",
        ok: "Ok",
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
  components: { CustomBtn2, CustomBtn3 },
  methods: {
    addItem() {
      if (this.isEdit) return;
      this.newItem = true;
      var lang = createNewLangWord(this.wordsList as ILanguageWord[]);
      (this.wordsList as ILanguageWord[]).unshift(lang);
      this.editMode(lang.id);
    },
    editMode(id: string) {
      if (this.isEdit) return;
      this.selectedId = id;
      this.previousValues = Object.assign(
        {},
        this.findLanguagebyId(id) as ILanguageWord
      );
      this.isEdit = true;
    },
    deleteWord(index: number) {
      (this.wordsList as ILanguageWord[]).splice(index, 1);
    },
    updateWord() {
      this.isEdit = !this.isValid();
      this.newItem = false;
    },
    cancelWord(index: number) {
      this.updatePreviousLanguage();
      (this.wordsList as ILanguageWord[]).splice(index, 1);
      this.isEdit = false;
      this.newItem = false;
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
    updatePreviousLanguage() {
      var words = this.wordsList?.map((obj: ILanguageWord) => {
        if (obj.id == this.selectedId) {
          return (obj = this.previousValues);
        }
        return obj;
      });
      this.wordsList = words;
    },
    findLanguagebyId(id: string) {
      return this.wordsList?.find((x: ILanguageWord) => x.id == id);
    },
    getLanguageData(): ILanguageWord[] {
      var words = this.wordsList as ILanguageWord[];
      return words.sort(langDataDescendingSort);
    },
  },
};

function createNewLangWord(languageData: ILanguageWord[]): ILanguageWord {
  let newObj: ILanguageWord;
  newObj = {
    id: getLastId(languageData),
    MainLanguage: "",
    ForeignLanguage: "",
  };
  return newObj;
}
function getLastId(languageData: ILanguageWord[]) {
  return (Math.max(...languageData.map((x) => parseInt(x.id))) + 1).toString();
}
function langDataDescendingSort(a: ILanguageWord, b: ILanguageWord) {
  return parseInt(b.id) - parseInt(a.id);
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
  font-size: 1.4rem;
}
.word-header {
  display: flex;
  align-items: center;
  padding-bottom: 0.8rem;
  justify-content: space-between;
}
.word-firstLangHeader,
.work-secondLangHeader {
  width: 4rem;
}
.word-firstLangHeader {
  padding-left: 10rem;
  padding-right: 14rem;
}
.word-secondLangHeader {
  padding-right: 39rem;
  width: 4rem;
}
.word-list-item {
  display: flex;
  align-items: center;
  padding-top: 0.7rem;
  padding-bottom: 0.7rem;
  border-top: 1px solid #9eb5cc;
  justify-content: space-between;
}
.last-item-bottom {
  border-top: 1px solid #9eb5cc;
}
.word-item-firstLang {
  padding-left: 10rem;
  padding-right: 1rem;
}
.word-item-secondLang {
  padding-right: 10rem;
}
.word-item-btns {
  padding-right: 5rem;
  padding-left: 0.7rem;
}
.textInput {
  border-radius: 0.3rem;
  box-sizing: border-box;
  font-size: 1.3rem;
}
.word-top {
  display: flex;
  font-size: 1rem;
  color: red;
  justify-content: space-between;
}
.word-error {
  width: 80%;
}
.errorBorder {
  border-color: red;
}
</style>
