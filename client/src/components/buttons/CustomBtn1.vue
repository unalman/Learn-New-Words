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
  background-color: #e67e22;
  font-size: 1.6rem;
  border-radius: 0.6rem;
  width: 10rem;
}
div .customBtn1:hover {
  color: var(--hotmag);
  text-decoration: none;
}
</style>
