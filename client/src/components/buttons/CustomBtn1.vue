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
  background-color: #5fa7db;
  font-size: 1.5rem;
  border-radius: 0.3rem;
  padding: 3px;
}
.customBtn1:hover {
  background-color: rgb(70 119 211);
}
.customBtn1:active {
  background-color: rgb(70 119 211);
  box-shadow: 0 1px rgb(70 119 211);
  transform: translateY(1px);
}
</style>
