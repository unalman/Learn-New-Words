<script lang="ts">
import WordsTable from "../components/WordsTable.vue";
export default {
  component: {
    WordsTable,
  },
};
</script>
<template>
  <div class="addword-container">
    <div class="word-item-1">Kelimeler</div>
    <div class="word-item-2">
      <WordsTable />
    </div>
  </div>
</template>

<style scoped>
.addword-container {
  display: grid;
  height: inherit;
  min-height: inherit;
  align-items: center;
  grid-template-columns: 1fr 1fr 8fr 1fr 1fr;
  grid-template-rows: 4rem 1fr;
  box-sizing: border-box;
  overflow: overlay;
}
.addword-container > .word-item-1,
.word-item-2 {
  grid-column: 3;
  box-sizing: border-box;
}
.word-item-1 {
  grid-row: 1;
  padding-top: 1.5vh;
  min-height: 38px;
}
.word-item-2 {
  grid-row: 2;
  box-sizing: border-box;
  height: inherit;
}
</style>
