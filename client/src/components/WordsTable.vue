<template>
  <div class="word-btn">
    <CustomBtn3 :btnName="buttonNames.add" />
  </div>
  <div class="word-container">
    <div class="word-header">
      <div class="word-firstLangHeader">Ana Dil</div>
      <div class="word-secondLangHeader">Yabancı Dil</div>
    </div>
    <div class="word-list">
      <div
        class="word-list-item"
        v-for="item in languageData"
        v-bind:key="item.id"
      >
        <div class="word-item-firstLang">{{ item.MainLanguage }}</div>
        <div class="word-item-secondLang">{{ item.ForeignLanguage }}</div>
        <div class="word-item-btns">
          <CustomBtn2 :btnName="buttonNames.edit" />
          <CustomBtn2 :btnName="buttonNames.delete" />
        </div>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
import CustomBtn2 from "./buttons/CustomBtn2.vue";
import CustomBtn3 from "./buttons/CustomBtn3.vue";
export default {
  component: {
    CustomBtn3,
    CustomBtn2,
  },
  data() {
    return {
      buttonNames: {
        add: "Ekle",
        edit: "Düzenle",
        delete: "Sil",
      },
      languageData: [
        {
          id: "1",
          MainLanguage: "Yazi1",
          ForeignLanguage: "Text1",
        },
        {
          id: "2",
          MainLanguage: "Yazi2",
          ForeignLanguage: "Text2",
        },
        {
          id: "3",
          MainLanguage: "Yazi3",
          ForeignLanguage: "Text3",
        },
      ],
    };
  },
  components: { CustomBtn2, CustomBtn3 },
};
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
</style>
