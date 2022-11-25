<template>
  <div class="word-btn">
    <CustomBtn3 @click="addItem()" :btnName="buttonNames.add" />
  </div>
  <div class="word-container">
    <div class="word-header">
      <div class="word-firstLangHeader">Ana Dil</div>
      <div class="word-secondLangHeader">Yabancı Dil</div>
    </div>
    <ul class="word-list" v-auto-animate>
      <li
        class="word-list-item"
        v-for="item in getLanguageData()"
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
          <input class="textInput" type="text" v-model="item.MainLanguage" />
        </div>
        <div
          class="word-item-secondLang"
          v-if="isEdit && selectedId == item.id"
        >
          <input class="textInput" type="text" v-model="item.ForeignLanguage" />
        </div>
        <div class="word-item-btns">
          <CustomBtn2
            class="edit"
            v-on:click="editMode(item.id)"
            v-if="!isEdit || selectedId != item.id"
            :btnName="buttonNames.edit"
          />
          <CustomBtn2
            class="delete"
            v-on:click="editMode(item.id)"
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
            v-on:click="cancelWord()"
            v-if="isEdit && selectedId == item.id"
            :btnName="buttonNames.cancel"
          />
        </div>
      </li>
    </ul>
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
      buttonNames: {
        add: "Ekle",
        edit: "Düzenle",
        delete: "Sil",
        ok: "Ok",
        cancel: "İptal",
      },
      isEdit: false,
      selectedId: "",
    };
  },
  components: { CustomBtn2, CustomBtn3 },
  methods: {
    addItem(this: any) {
      var lang = createNewLangWord(this.languageData);
      (this.languageData as ILanguageWord[]).unshift(lang);
      this.selectedId = lang.id;
      this.isEdit = true;
    },
    editWord(this: any) {
      this.isEdit = true;
    },
    deleteWord(this: any) {
      //remove event
    },
    updateWord(this: any) {
      this.isEdit = false;
    },
    cancelWord(this: any) {
      this.isEdit = false;
    },
    editMode(id: string) {
      console.log(id);
      this.selectedId = id;
      this.isEdit = true;
    },
    getLanguageData(this: any): ILanguageWord[] {
      var data = this.languageData as ILanguageWord[];
      return data.sort(langDataDescendingSort);
    },
  },
};

function createNewLangWord(languageData: ILanguageWord[]): ILanguageWord {
  let obj: ILanguageWord;
  obj = {
    id: getLastId(languageData),
    MainLanguage: "",
    ForeignLanguage: "",
  };
  return obj;
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
</style>
