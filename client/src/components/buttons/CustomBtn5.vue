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
  color: rgb(38 105 205);
}
/* a:hover {
  color: var(--hotmag);
  text-decoration: none;
} */
</style>
