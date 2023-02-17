<template>
  <button data-cy="btn3" class="customBtn3">{{ btnName }}</button>
</template>
<script lang="ts">
export default {
  props: {
    btnName: String,
  },
};
</script>
<style scoped>
.customBtn3 {
  background-color: #5fa7db;
  color: white;
  cursor: pointer;
  font-size: 1.4rem;
  border: 1px;
  border-radius: 0.28rem;
  height: 2.5rem;
  box-shadow: 0 2px 0 0 #b1d7f3;
  padding: 5px;
}
.customBtn3:active {
  box-shadow: 0 2px #b1d7f3;
  transform: translateY(1px);
}
.customBtn3:hover {
  background-color: rgb(70 119 211);
}
</style>
