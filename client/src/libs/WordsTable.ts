import type { ILanguageWord } from "typings/interface/ILanguageWord";

export const getLastId = (languageData: ILanguageWord[]): string => {
  if ((!languageData && languageData == null) || languageData.length == 0)
    return "1";
  return (Math.max(...languageData.map((x) => parseInt(x.id))) + 1).toString();
};
export const sortLanguageWordDescending = (languageWords: ILanguageWord[]) => {
  return languageWords.sort((a, b) => (parseInt(a.id) > parseInt(b.id) ? -1 : 1));
};