import type { ILanguageWord } from "typings/interface/ILanguageWord";

interface IWordTable {
  wordsList: ILanguageWord[];
  isEdit: boolean;
  selectedId: string;
  newItem: boolean;
  errors: string[];
  validation: {
    mainLanguage: boolean;
    foreignLanguage: boolean;
  };
  errorClasses: {
    errorBorder: string;
  };
  previousValues: ILanguageWord;
}

export type { IWordTable };
