import { defineStore } from "pinia";
import type { ILanguageWord } from "typings/interface/IWordTable";

type StateShape = {
  list: ILanguageWord[];
};

export const useLangueageWordStore = defineStore({
  id: "LanguageWord",
  //Data
  state: (): StateShape => ({
    list: [
      {
        id: "1",
        MainLanguage: "Yazi1",
        ForeignLanguage: "Text1",
      },
      {
        id: "2",
        MainLanguage: "Yazi2",
        ForeignLanguage: "Text2",
      },
      {
        id: "3",
        MainLanguage: "Yazi3",
        ForeignLanguage: "Text3",
      },
    ],
  }),
  //Computed
  getters: {},
  //Methods
  actions: {},
});
