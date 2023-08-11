import { defineStore } from "pinia";
import type { ILanguageWord } from "typings/interface/ILanguageWord";

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
      {
        id: "4",
        MainLanguage: "Yazi4",
        ForeignLanguage: "Text4",
      },
      {
        id: "5",
        MainLanguage: "Yazi5",
        ForeignLanguage: "Text5",
      }
    ],
  }),
  //Computed
  getters: {},
  //Methods
  actions: {},
});
