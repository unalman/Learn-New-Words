import http from "./api";
import type { APIResponse } from "typings/types/ApiResponse";
import type { ILanguageWord } from "typings/interface/ILanguageWord";

async function getWords() {
  return await http.get<APIResponse<ILanguageWord[]>>("words");
}
async function getRandomWords() {
  return await http.get<APIResponse<ILanguageWord[]>>("random-words");
}
async function getWord(id: string) {
  return await http.get<APIResponse<ILanguageWord>>(`words/${id}`);
}
async function deleteWord(id: string) {
  return await http.delete<APIResponse<null>>(`words/${id}`);
}
async function createWord(word: ILanguageWord) {
  return await http.post<APIResponse<ILanguageWord>>("words", {
    MainLanguage: word.MainLanguage,
    ForeignLanguage: word.ForeignLanguage,
  });
}
async function updateWord(word: ILanguageWord) {
  return await http.put<APIResponse<null>>("words", word);
}

export default {
  getWords,
  getRandomWords,
  getWord,
  deleteWord,
  createWord,
  updateWord,
};
