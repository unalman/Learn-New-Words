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

#app {
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;
  padding: 0;
  margin: 0;
  font-size: 15px;
  box-sizing: border-box;
  min-width: 560px;
}

.grid-container {
  display: grid;
  background-color: #2196f3;
  box-sizing: border-box;
}

.grid-container > .item-1,
.item-2,
.item-3 {
  background-color: rgba(255, 255, 255, 0.8);
  text-align: center;
  font-size: 30px;
  box-sizing: border-box;
}

.item-1 {
  height: 6.6vh;
  min-height: 25px;
}
.item-2 {
  height: 91.4vh;
  min-height: 185px;
  width: 100%;
}
.item-3 {
  height: 2vh;
  min-height: 17px;
}
</style>
