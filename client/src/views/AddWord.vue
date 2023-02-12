<template>
  <div class="addword-container">
    <div class="word-item-1">Kelimeler</div>
    <div class="word-item-2">
      <WordsTable :languageData="languageData" />
    </div>
  </div>
</template>
<script lang="ts">
import WordsTable from "../components/WordsTable.vue";
import { defineComponent } from "vue";
import { useLangueageWordStore } from "@/stores/LanguageWord";

export default defineComponent({
  setup() {
    const langWordStore = useLangueageWordStore();
    const list = langWordStore.list;
    return {
      languageData: list,
    };
  },
  components: { WordsTable },
});
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
