import { createI18n } from "vue-i18n";
import en from "./en.json";
import tr from "./tr.json";

const defaultLocale = "tr";
const currentLocale = localStorage.getItem("culture") ?? defaultLocale;
const i18n = createI18n({
  locale: currentLocale,
  fallbackLocale: "en",
  messages: {
    en: en,
    tr: tr,
  },
  legacy: false,
  sync: true,
  allowComposition: true,
});

export { i18n, defaultLocale,currentLocale };
