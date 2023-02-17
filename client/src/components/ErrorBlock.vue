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
        <li v-for="item in errors" :key="item">
          {{ item }}
        </li>
      </ul>
    </div>
  </div>
</template>
<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  props: { errors: Array as () => string[] },
});
</script>
<style>
@import "../assets/errorBlock.css";
</style>
