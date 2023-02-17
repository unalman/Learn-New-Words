<template>
  <nav>
    <ul>
      <li>
        <CustomBtn5
          data-cy="homeBtn"
          :btnName="homePage.name"
          :href="homePage.href"
        />
      </li>
      <li>
        <CustomBtn5 :btnName="contactPage.name" :href="contactPage.href" />
      </li>
    </ul>
  </nav>
</template>
<script lang="ts">
import CustomBtn5 from "./buttons/CustomBtn5.vue";
import { defineComponent } from "vue";

export default defineComponent({
  component: { CustomBtn5 },
  data() {
    return {
      homePage: {
        name: "Anasayfa",
        href: "/",
      },
      contactPage: {
        name: "İletişim",
        href: "/contact",
      },
    };
  },
  components: { CustomBtn5 },
});
</script>
<style scoped>
@import "../assets/baseHeader.css";
</style>
