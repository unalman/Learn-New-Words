<template>
  <a class="customBtn1">{{ btnName }}</a>
</template>
<script lang="ts">
export default {
  props: {
    btnName: String,
  },
};
</script>
<style scoped>
.customBtn1 {
  color: white;
  display: inline-block;
  background-color: rgb(106 146 205);
  font-size: 1.5rem;
  border-radius: 0.3rem;
  padding: 3px;
}
div .customBtn1:hover {
  background-color: rgb(38 105 205);
  text-decoration: none;
}
</style>
