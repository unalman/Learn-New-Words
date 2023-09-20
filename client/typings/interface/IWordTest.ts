import type { ILanguageWord } from "typings/interface/ILanguageWord";
interface IWordTest {
  languageWords: Array<ILanguageWord>;
}
interface IWordTestReturn extends ILanguageWord {
  isCorrectAnswer: Boolean;
}

export type { IWordTest, IWordTestReturn };
