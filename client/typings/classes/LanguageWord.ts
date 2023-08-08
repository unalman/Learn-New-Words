import type { ILanguageWord } from "typings/interface/ILanguageWord";

export default class LanguageWord implements ILanguageWord {
  id: string;
  MainLanguage: string;
  ForeignLanguage: string;
  public static createInstance(id: string): LanguageWord {
    const instance = new LanguageWord();
    instance.id = id;
    return instance;
  }
}
