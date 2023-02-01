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
.errorBlock-container {
  background-color: #fff4f4;
  border: 1px solid #ffb9b9;
  border-radius: 4px;
  text-align: center;
  width: 100%;
  height: 100%;
  display: flex;
  align-items: baseline;
  text-align: initial;
}
.errorBlock-ul {
  font-size: 14px;
  color: #cf0000;
  margin: 8px;
}
.errorBlock-ul li {
  margin-top: 7px;
}
.errorBlock-iconTab-errorIcon {
  color: #b52121;
  margin: 5px;
}
</style>
