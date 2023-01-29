!
<template>
  <button type="submit" class="customBtn2">{{ btnName }}</button>
</template>
<script lang="ts">
export default {
  props: {
    btnName: String,
  },
};
</script>
<style scoped>
.customBtn2 {
  background-color: #5fa7db;
  color: white;
  cursor: pointer;
  font-size: 20px;
  border: 1px;
  border-radius: 0.28rem;
  margin: 6px;
  box-sizing: border-box;
  padding: 5px;
}
.customBtn2:hover {
  background-color: rgb(70 119 211);
}
.customBtn2:active {
  background-color: rgb(70 119 211);
  box-shadow: 0 1px rgb(70 119 211);
  transform: translateY(1px);
}
</style>
