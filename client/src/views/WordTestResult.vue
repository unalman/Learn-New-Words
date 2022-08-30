<template>
  <div class="wordTestResult-container">
    <div class="wordTestResult-ResultContainer">
      <div class="wordTestResult-ResultText">
        🎉 Tebrikler Testi Başarıyla Tamamladınız 🎉
      </div>
      <div class="wordTestResult-WordsListText">Testteki Kelimeler</div>
      <div class="wordTestResult-WordList">
        <div class="wordTestResult-WordBox">
          <span>Text: </span>
          <span>Yazi</span>
        </div>
        <div class="wordTestResult-WordBox">
          <span>Text2: </span>
          <span>Yazi2</span>
        </div>
        <div class="wordTestResult-WordBox">
          <span>Text10: </span>
          <span>Yazi10</span>
        </div>
      </div>
      <div class="wordTestResult-last-item-bottom"></div>
      <div><button class="wordTestResult-ComplatedBtn">Tamamla</button></div>
    </div>
  </div>
</template>
<script lang="ts">
export default {};
</script>
<style scoped>
.wordTestResult-container {
  display: grid;
  grid-template-columns: 1fr 1fr 8fr 1fr 1fr;
  height: inherit;
  min-height: inherit;
  box-sizing: border-box;
  overflow: overlay;
}
.wordTestResult-ResultContainer {
  grid-column: 3;
  display: flex;
  text-align: center;
  flex-direction: column;
  padding-top: 10vh;
}
.wordTestResult-ResultText {
  overflow: hidden;
  padding: 1rem;
  font-size: 2.4rem;
  padding-bottom: 2rem;
}
.wordTestResult-WordsListText{
  padding: 1rem;
}
.wordTestResult-WordList {
  display: flex;
  flex-direction: column;
}
.wordTestResult-WordBox {
  border-top: 1px solid #9eb5cc;
  padding: 0.5rem;
}
.wordTestResult-last-item-bottom {
  border-top: 1px solid #9eb5cc;
  padding: 0.5rem;
}
.wordTestResult-ComplatedBtn{
  border: 1px;
  border-radius: 0.2rem;
  cursor: pointer;
  font-size: 1.4rem;
  height: 3rem;
  margin-top: 1rem;
}
</style>
