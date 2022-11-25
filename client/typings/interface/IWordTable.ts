interface IWordTable {
  isEdit: boolean;
  selectedId: string;
  buttonNames: {
    add: string;
    edit: string;
    delete: string;
    ok: string;
    cancel: string;
  };
}

interface ILanguageWord {
  id: string;
  MainLanguage: string;
  ForeignLanguage: string;
}

export type { IWordTable, ILanguageWord };
