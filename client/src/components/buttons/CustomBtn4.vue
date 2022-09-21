<template>
  <button class="custombtn4">{{ btnName }}</button>
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
  border: 1px;
  border-radius: 0.2rem;
  cursor: pointer;
  font-size: 1.4rem;
  height: 3rem;
}
</style>
