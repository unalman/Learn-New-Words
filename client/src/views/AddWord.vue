<template>
  <div class="addword-container">
    <div class="title item-1">H1</div>
    <div class="title item-2">content</div>
  </div>
</template>
<script lang="ts">
export default {};
</script>
<style scoped>
.addword-container {
  display: grid;
  height: inherit;
  min-height: inherit;
  align-items: center;
  background-color: blueviolet;
  grid-template-columns: 1fr 1fr 8fr 1fr 1fr;
  grid-template-rows: 1fr 1fr;
  box-sizing: border-box;
  overflow: hidden;
}

.addword-container > .item-1,
.item-2 {
  grid-column: 3;
  box-sizing: border-box;
}

.item-1 {
}
.item-2 {
  overflow: hidden;
}
</style>
