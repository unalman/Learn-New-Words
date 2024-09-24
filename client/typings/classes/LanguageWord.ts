import type { ILanguageWord } from "typings/interface/ILanguageWord";

export default class LanguageWord implements ILanguageWord {
  id: string;
  MainLanguage: string;
  ForeignLanguage: string;
  created_at: Date;
  public static createInstance(): LanguageWord {
    const instance = new LanguageWord();
    instance.id = "new";
    return instance;
  }
}
