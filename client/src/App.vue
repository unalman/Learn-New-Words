<script setup lang="ts">
import BaseHeaderVue from "./components/BaseHeader.vue";
import BaseFooterVue from "./components/BaseFooter.vue";
</script>

<template>
  <div class="grid-container">
    <header class="item item-1">
      <BaseHeaderVue></BaseHeaderVue>
    </header>
    <main class="item item-2">
      <router-view />
    </main>
    <footer class="item item-3">
      <BaseFooterVue></BaseFooterVue>
    </footer>
  </div>
</template>

<style>
@import "./assets/base.css";
@import "https://meyerweb.com/eric/tools/css/reset/reset.css";

*{
  box-sizing: border-box;
  padding: 0;
  margin: 0;
}

.grid-container {
  display: grid;
  background-color: black;
  grid-template-rows: 6.6vh calc(100% - 6.6vh - 2.8vh) 2.8vh;
  width: 100vw;
  height: 100vh;
}

.grid-container > .item-1,
.item-2,
.item-3 {
  grid-column: 1;
  background-color: var(--mainColor);
  text-align: center;
  font-size: 30px;
  box-sizing: border-box;
}

.item-1 {
  grid-column: 1;
  grid-row: 1;
}
.item-2 {
  grid-column: 1;
  grid-row: 2;
}
.item-3 {
  grid-column: 1;
  grid-row: 3;
}
</style>
