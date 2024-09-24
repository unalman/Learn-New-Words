<template>
  <div class="errorBlock-container">
    <div class="errorBlock-iconTab">
      <font-awesome-icon
        icon="fa-solid fa-triangle-exclamation"
        class="errorBlock-iconTab-errorIcon"
      />
    </div>
    <div>
      <ul class="errorBlock-ul">
        <li v-for="item in props.errors" :key="item">
          {{ t(item) }}
        </li>
      </ul>
    </div>
  </div>
</template>
<script setup lang="ts">
import { useI18n } from "vue-i18n";
const { t } = useI18n();

const props = defineProps({
  errors: Array as () => string[],
});

</script>
<style>
@import "../assets/errorBlock.css";
</style>
