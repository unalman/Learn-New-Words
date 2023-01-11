<template>
  <button class="customBtn3">{{ btnName }}</button>
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
  background-color: var(--dkblue);
  color: white;
  cursor: pointer;
  font-size: 1.4rem;
  border: 1px;
  border-radius: 1rem;
  height: 2.5rem;
  box-shadow: 0 2px 0 0 #000;
  padding: 5px;
}
.customBtn3:active {
  box-shadow: 0 2px #000;
  transform: translateY(1px);
}
.customBtn3:hover {
  background-color: rgb(21, 72, 168);
}
</style>
