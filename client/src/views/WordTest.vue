<template>
  <div class="wordTest-container">
    <div class="wordTest-mainItem">
      <div class="wordTest-Question">Yazi</div>
      <div class="wordTest-Answer">
        <input type="text" class="answerText" placeholder="Cevap" />
      </div>
      <div class="wordTest-btnList">
        <CustomBtn4 :btnName="btnidk.name" />
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
  overflow: overlay;
}
.wordTest-mainItem {
  grid-column: 3;
  margin-top: 14vh;
}
.wordTest-Question {
  font-size: 3rem;
  margin-bottom: 4rem;
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
}
.answerText:focus {
  border: 1px solid var(--InputFocusGrey);
}
.wordTest-btnList {
  margin-top: 1rem;
  display: flex;
  justify-content: space-evenly;
}
</style>
