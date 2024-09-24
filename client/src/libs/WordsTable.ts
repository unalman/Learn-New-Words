import type { ILanguageWord } from "typings/interface/ILanguageWord";

export const sortLanguageWordDescending = (languageWords: ILanguageWord[]) => {
  if (!languageWords && languageWords == null) return [];
  return languageWords.sort((a, b) =>
    parseInt(a.id) > parseInt(b.id) ? -1 : 1
  );
};
