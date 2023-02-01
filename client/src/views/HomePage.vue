<template>
  <div class="content-container">
    <div class="content-item1">
      <div>Ezberlemek istediğin kelimeleri ekle!</div>
      <div class="content-btnContainer">
        <CustomBtn1 :btnName="btnAddWord.name" :href="btnAddWord.href" />
      </div>
    </div>
    <div class="content-item2">
      <div>Kendini test et!</div>
      <div class="content-btnContainer">
        <CustomBtn1 :btnName="btnTestWord.name" :href="btnTestWord.href" />
      </div>
    </div>
    <div class="content-item3">
      <div class="content-btnContainer">
        <CustomBtn1 :btnName="btnResultPage.name" :href="btnResultPage.href" />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import CustomBtn1 from "../components/buttons/CustomBtn1.vue";
export default {
  components: { CustomBtn1 },
  data() {
    return {
      btnAddWord: {
        name: "Kelime Ekle",
        href: "/addword",
      },
      btnTestWord: {
        name: "Kelime Testi",
        href: "/wordtest",
      },
      btnResultPage: {
        name: "Result Sayfası",
        href: "/wordtestresult",
      },
    };
  },
};
</script>

<style scoped>
.content-container {
  display: grid;
  padding: 2vh;
  grid-gap: 10px;
  height: 100%;
  min-height: inherit;
  align-items: center;
  justify-content: space-around;
  color: darkslategray;
}

.content-container > .content-item1,
.content-item2 {
  text-align: center;
  font-size: 1.7rem;
  box-sizing: border-box;
  width: 40vh;
  border: solid 0.4px;
  border-color: var(--lightGrey);
  min-width: 180px;
  height: 14vh;
  border-radius: 1rem;
  padding-top: 1rem;
  background-color: var(--lightBlue);
}
.content-item1 {
  grid-template-columns: 100px auto;
  min-height: 144px;
}
.content-item2 {
  grid-row-start: 2;
  grid-column-start: 2;
  min-height: 101px;
}

@media screen and (max-width: 800px) {
  .content-item2 {
    grid-row-start: 2;
    grid-column-start: 1;
    min-height: 101px;
  }
}
.content-btnContainer {
  padding-top: 1rem;
  text-align: center;
}
</style>
