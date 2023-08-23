import type { ILanguageWord } from "typings/interface/ILanguageWord";
interface IWordTest {
  languageWords: Array<ILanguageWord>;
  buttonNames: {
    check: string;
    idk: string;
    continue: string;
    result: string;
  };
  placeholderText:string;
  correct: string;
  answer: string;
}
interface IWordTestReturn extends ILanguageWord {
  isCorrectAnswer: Boolean;
}

export type { IWordTest, IWordTestReturn };
