<template>
  <button data-cy="btn4" class="custombtn4">{{ btnName }}</button>
</template>
<script lang="ts">
export default {
  props: {
    btnName: String,
  },
};
</script>
<style scoped>
.custombtn4 {
  background-color: #5fa7db;
  color: white;
  border: 1px;
  border-radius: 0.28rem;
  cursor: pointer;
  font-size: 1.4rem;
  padding: 7px;
}
.custombtn4:hover {
  background-color: rgb(70 119 211);
}
.custombtn4:active {
  background-color: rgb(70 119 211);
  box-shadow: 0 1px rgb(70 119 211);
  transform: translateY(1px);
}
</style>
