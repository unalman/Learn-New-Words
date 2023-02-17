<template>
  <a data-cy="btn5" class="custombtn5">{{ btnName }}</a>
</template>
<script lang="ts">
export default {
  props: {
    btnName: String,
  },
};
</script>
<style scoped>
.custombtn5 {
  text-decoration: none;
  color: rgb(38 105 205);
}
</style>
