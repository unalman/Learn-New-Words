<template>
  <div class="wordTestResult-container">
    <div class="wordTestResult-ResultContainer">
      <div class="wordTestResult-ResultText">
        🎉 Tebrikler Testi Başarıyla Tamamladınız 🎉
      </div>
      <div class="wordTestResult-TableContainer">
        <div class="wordTestResult-WordsListText">Kelimeler</div>
        <div class="wordTestResult-WordList">
          <div class="wordTestResult-WordBox">
            <span>Text: </span>
            <span>Yazi</span>
            <span> - 👎</span>
          </div>
          <div class="wordTestResult-WordBox">
            <span>Text2: </span>
            <span>Yazi2</span>
            <span> - 👍</span>
          </div>
          <div class="wordTestResult-WordBox">
            <span>Text10: </span>
            <span>Yazi10</span>
            <span> - 👍</span>
          </div>
        </div>
      </div>
      <div>
        <CustomBtn4 :btnName="btnComplate.name" />
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
      btnComplate: {
        name: "Tamamla",
      },
    };
  },
};
</script>
<style scoped>
.wordTestResult-container {
  display: grid;
  grid-template-columns: 1fr 1fr 8fr 1fr 1fr;
  height: 100%;
  min-height: inherit;
  overflow: overlay;
}
.wordTestResult-ResultContainer {
  grid-column: 3;
  display: flex;
  text-align: center;
  flex-direction: column;
  margin-top: 10vh;
}
.wordTestResult-ResultText {
  padding: 1rem;
  font-size: 2rem;
  padding-bottom: 2rem;
  font-weight: 600;
  color: #4a4a4a;
}
.wordTestResult-TableContainer {
  box-shadow: 0 2px 6px 0 hsla(0, 0%, 0%, 0.2);
  margin-bottom: 2rem;
}
.wordTestResult-WordsListText {
  background-color: var(--lightBlue);
  padding: 1rem;
  font-size: 1.7rem;
  font-weight: 600;
  color: #4a4a4a;
}
.wordTestResult-WordList {
  display: flex;
  flex-direction: column;
  font-size: 1.5rem;
  color: #4a4a4a;
}
.wordTestResult-WordBox {
  border-top: 1px solid var(--lightGreyTableBorder);
  padding: 0.5rem;
}
</style>
