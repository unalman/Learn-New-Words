import type { ILanguageWord } from "typings/interface/ILanguageWord";

interface IWordTable {
  wordsList: ILanguageWord[];
  isEdit: boolean;
  selectedId: string;
  newItem: boolean;
  buttonNames: {
    add: string;
    edit: string;
    delete: string;
    ok: string;
    cancel: string;
  };
  gridFieldNames: {
    firstField: string;
    secondField: string;
  };
  errors: string[];
  validation: {
    mainLanguage: boolean;
    mainLanguageText: string;
    foreignLanguage: boolean;
    foreignLanguageText: string;
    noResult: string;
  };
  errorClasses: {
    errorBorder: string;
  };
  previousValues: ILanguageWord;
}

export type { IWordTable };
