<script setup lang="ts">
import BaseHeaderVue from "./components/BaseHeader.vue";
import BaseMainVue from "./components/BaseMain.vue";
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

#app {
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;
  padding: 0;
  margin: 0;
  font-size: 15px;
  box-sizing: border-box;
}

.grid-container {
  display: grid;
  background-color: #2196f3;
  height: 100%;
  widows: 100%;
}

.grid-container > .item-1,
.item-2,
.item-3 {
  background-color: rgba(255, 255, 255, 0.8);
  text-align: center;
  font-size: 30px;
}

.item-1 {
  height: 6.6vh;
  min-width: 331px;
  min-height: 60.38px;
}
.item-2 {
  height: 91.4vh;
}
.item-3 {
  height: 2vh;
  min-height: 17px;
}
</style>
