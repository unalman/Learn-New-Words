<template>
  <div class="addword-container">
    <div class="word-item-1">Kelimeler</div>
    <div class="word-item-2">
      <WordsTable :languageData="list" />
    </div>
  </div>
</template>
<script setup lang="ts">
import WordsTable from "@/components/WordsTable.vue";
import { useLangueageWordStore } from "@/stores/LanguageWord";

const langWordStore = useLangueageWordStore();
const list = langWordStore.list;
</script>
<style scoped>
.addword-container {
  display: grid;
  min-height: inherit;
  grid-template-columns: 1fr 1fr 8fr 1fr 1fr;
  grid-template-rows: 4rem 1fr;
}
.addword-container > .word-item-1,
.word-item-2 {
  grid-column: 3;
}
.word-item-1 {
  grid-row: 1;
  margin-top: 2.5vh;
  font-weight: 600;
  color: darkslategray;
}
.word-item-2 {
  grid-row: 2;
}
</style>
