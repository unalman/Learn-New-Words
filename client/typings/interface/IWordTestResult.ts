import type { IWordTestReturn } from "typings/interface/IWordTest";
interface IWordTestResult {
  wordTestReturns: Array<IWordTestReturn>;
  congratulationsText: string;
  words: string;
  button: {
    complate: string;
  };
}
export type { IWordTestResult };
