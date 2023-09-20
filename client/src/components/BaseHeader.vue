<template>
  <nav>
    <ul>
      <li>
        <CustomBtn5
          data-cy="homeBtn"
          :btnName="t('homePageName')"
          :href="data.homePage.href"
        />
      </li>
      <li>
        <ul>
          <li class="li-locales">
            <div class="locale-changer">
              <select
                class="locale"
                v-model="$i18n.locale"
                @change="onChange($event)"
              >
                <option
                  v-for="locale in $i18n.availableLocales"
                  :key="`locale-${locale}`"
                  :value="locale"
                >
                  {{ locale.toUpperCase() }}
                </option>
              </select>
            </div>
          </li>
          <li>
            <CustomBtn5
              :btnName="t('contactPageName')"
              :href="data.contactPage.href"
            />
          </li>
        </ul>
      </li>
    </ul>
  </nav>
</template>
<script setup lang="ts">
import { ref } from "vue";
import CustomBtn5 from "./buttons/CustomBtn5.vue";
import type { IHeaderPageWord } from "typings/interface/IHeaderPageWord";
import { useI18n } from "vue-i18n";
const { t } = useI18n({ useScope: "global" });
const data = ref<IHeaderPageWord>({
  homePage: {
    href: "/",
  },
  contactPage: {
    href: "/contact",
  },
});
function onChange(event: any) {
  localStorage.setItem("culture", event.target.value.toLowerCase());
}
</script>
<style scoped>
@import "../assets/baseHeader.css";
</style>
