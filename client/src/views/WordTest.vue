<template>
  <div class="wordTest-container">
    <div class="wordTest-mainItem">
      <div class="wordTest-Question">Yazi</div>
      <div class="wordTest-Answer">
        <input type="text" class="answerText" placeholder="Cevabı Buraya Yaz" />
      </div>
      <div class="wordTest-btnList">
        <CustomBtn4 class="wordTest-btnDontKnow" :btnName="btnidk.name" />
        <CustomBtn4 :btnName="btnCheck.name" />
      </div>
    </div>
  </div>
</template>
<script lang="ts">
import CustomBtn4 from "../components/buttons/CustomBtn4.vue";
export default {
  components: { CustomBtn4 },
  data() {
    return {
      btnCheck: {
        name: "Kontrol et",
      },
      btnidk: {
        name: "Bilmiyorum",
      },
    };
  },
};
</script>
<style scoped>
.wordTest-container {
  display: grid;
  height: 100%;
  min-height: inherit;
  grid-template-columns: 1fr 1fr 8fr 1fr 1fr;
  box-sizing: border-box;
  overflow: overlay;
}
.wordTest-mainItem {
  grid-column: 3;
  background-color: rgb(43, 86, 226);
  border-radius: 0.3rem;
  padding: 5rem;
  padding-top: 18vh;
}
.wordTest-Question {
  font-size: 3rem;
  padding-bottom: 4rem;
}
.answerText {
  border-radius: 0.2rem;
  border: 1px;
  height: 2rem;
  width: 25rem;
  font-size: inherit;
}
.wordTest-btnList {
  padding-top: 1rem;
  justify-content: space-between;
}
.wordTest-btnDontKnow {
  margin-right: 12rem;
}
</style>
