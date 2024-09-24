interface IHomePageVM {
  btnAddWord: {
    href: string;
  };
  btnTestWord: {
    name: string;
    isDisabled: boolean;
  };
  error: string[];
}
export type { IHomePageVM };
