interface IWordTable {
  wordsList: ILanguageWord[];
  isEdit: boolean;
  selectedId: string;
  buttonNames: {
    add: string;
    edit: string;
    delete: string;
    ok: string;
    cancel: string;
  };
  errors: string[];
  validation: {
    mainLanguage: boolean;
    mainLanguageText: string;
    foreignLanguage: boolean;
    foreignLanguageText: string;
  };
  errorClasses: {
    errorBorder: "errorBorder";
  };
  previousValues: ILanguageWord;
}

interface ILanguageWord {
  id: string;
  MainLanguage: string;
  ForeignLanguage: string;
}

export type { IWordTable, ILanguageWord };
