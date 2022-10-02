<template>
  <a>{{ btnName }}</a>
</template>
<script lang="ts">
export default {
  props: {
    btnName: String,
  },
};
</script>
<style scoped>
a {
  text-decoration: none;
  color: var(--white);
}
a:hover {
  color: var(--hotmag);
  text-decoration: none;
}
</style>
