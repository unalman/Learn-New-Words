<template>
  <WordTestGame :languageData="langWordStore.list" />
</template>
<script setup lang="ts">
import WordTestGame from "../components/WordTestGame.vue";
import { useLangueageWordStore } from "@/stores/LanguageWord";

const langWordStore = useLangueageWordStore();
</script>
