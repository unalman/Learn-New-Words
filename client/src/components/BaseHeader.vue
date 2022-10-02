<template>
  <nav>
    <ul>
      <li><CustomBtn5 :btnName="homePage.name" :href="homePage.href" /></li>
      <li>
        <CustomBtn5 :btnName="contactPage.name" :href="contactPage.href" />
      </li>
    </ul>
  </nav>
</template>
<script lang="ts">
import CustomBtn5 from "./buttons/CustomBtn5.vue";
export default {
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
};
</script>
<style scoped>
nav {
  height: 100%;
  font-size: 1.2rem;
  background-color: rgb(230, 182, 124);
  display: flex;
}
nav ul {
  width: 100%;
  display: flex;
  flex-flow: row wrap;
  box-sizing: border-box;
}
ul li {
  list-style-type: none;
  align-self: center;
  margin-right: auto;
}
nav ul li:first-child {
  margin-left: 1rem;
}
ul li:nth-child(2) {
  margin-right: 1rem;
}
</style>
