import type { ILanguageWord } from "typings/interface/ILanguageWord";

interface IWordTable {
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
