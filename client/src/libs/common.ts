import type { ILanguageWord } from "typings/interface/ILanguageWord";

export const textValidationAndTrim = (text: string): boolean => {
  const trimmedValue = text?.replace(/\s/g, "");
  if (text === null || trimmedValue == "") {
    return false;
  }
  return true;
};
