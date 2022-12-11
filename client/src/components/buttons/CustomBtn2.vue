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
  background-color: var(--dkblue);
  color: white;
  cursor: pointer;
  font-size: 1.4rem;
  border: 1px;
  border-radius: 1rem;
  margin: 6px;
  box-sizing: border-box;
}
.customBtn2:hover {
  background-color: rgb(21, 72, 168);
}
.customBtn2:active {
  background-color: var(--dkblue);
  box-shadow: 0 1px #000;
  transform: translateY(1px);
}
</style>
