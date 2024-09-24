import { defineStore } from "pinia";
import { ref } from "vue";
import type { ILanguageWord } from "typings/interface/ILanguageWord";
import type { APIResponse } from "typings/types/ApiResponse";
import { AxiosError } from "axios";
import { API } from "@/services";

export const useLanguageWordStore = defineStore("languageWordStore", () => {
  const words = ref<ILanguageWord[]>([]);
  const randomWords = ref<ILanguageWord[]>([]);
  function initLanguageWord(data: ILanguageWord[]) {
    words.value = data;
  }
  function initRandomLanguageWord(data: ILanguageWord[]) {
    randomWords.value = data;
  }

  function removeLanguageWord(id: string) {
    const idx = words.value.findIndex((x) => x.id == id);
    if (idx == -1) return;
    words.value.splice(idx, 1);
  }
  function addNewLanguageWord(oldId: string, id: string) {
    const index = words.value.findIndex((x) => x.id == oldId);
    if (index != -1) {
      words.value[index].id = id;
    }
  }

  async function dispatchGetLanguageWords(): Promise<APIResponse<null>> {
    try {
      const { status, data } = await API.languageWord.getWords();
      if (status == 200) {
        initLanguageWord(data.data);
        return {
          success: true,
          data: null,
          error: "",
        };
      }
    } catch (error) {
      const _error = error as AxiosError<string>;
      return {
        data: null,
        error: _error.cause?.message,
        success: false,
      } as APIResponse<null>;
    }
    return {
      data: null,
      error: "Access problem",
      success: false,
    };
  }
  async function dispatchCreateWord(
    model: ILanguageWord
  ): Promise<APIResponse<null>> {
    try {
      const { status, data } = await API.languageWord.createWord(model);
      if (status == 200) {
        addNewLanguageWord(model.id, data.data.id);
        return {
          success: true,
          data: null,
          error: "",
        };
      }
    } catch (error) {
      const _error = error as AxiosError<string>;
      return {
        data: null,
        error: _error.cause?.message,
        success: false,
      } as APIResponse<null>;
    }
    return {
      data: null,
      error: "Access problem",
      success: false,
    };
  }
  async function dispatchDeleteWord(id: string): Promise<APIResponse<null>> {
    try {
      const { status, data } = await API.languageWord.deleteWord(id);
      if (status == 200) {
        removeLanguageWord(id);
        return {
          success: true,
          data: null,
          error: "",
        };
      }
    } catch (error) {
      const _error = error as AxiosError<string>;
      return {
        data: null,
        error: _error.cause?.message,
        success: false,
      } as APIResponse<null>;
    }
    return {
      data: null,
      error: "Access problem",
      success: false,
    };
  }
  async function dispatchUpdateWord(
    model: ILanguageWord
  ): Promise<APIResponse<null>> {
    try {
      const { status, data } = await API.languageWord.updateWord(model);
      if (status == 200) {
        return {
          success: true,
          data: null,
          error: "",
        };
      }
    } catch (error) {
      const _error = error as AxiosError<string>;
      return {
        data: null,
        error: _error.cause?.message,
        success: false,
      } as APIResponse<null>;
    }
    return {
      data: null,
      error: "Access problem",
      success: false,
    };
  }
  async function dispatchGetRandomLanguageWords(): Promise<APIResponse<null>> {
    try {
      const { status, data } = await API.languageWord.getRandomWords();
      if (status == 200) {
        initRandomLanguageWord(data.data);
        return {
          success: true,
          data: null,
          error: "",
        };
      }
    } catch (error) {
      const _error = error as AxiosError<string>;
      return {
        data: null,
        error: _error.cause?.message,
        success: false,
      } as APIResponse<null>;
    }
    return {
      data: null,
      error: "Access problem",
      success: false,
    };
  }

  return {
    words,
    randomWords,
    initLanguageWord,
    removeLanguageWord,
    addNewLanguageWord,
    dispatchGetLanguageWords,
    dispatchCreateWord,
    dispatchDeleteWord,
    dispatchUpdateWord,
    dispatchGetRandomLanguageWords,
  };
});

// export const useWordStore = defineStore("wordStore2", {
//   state: () => ({
//     languageWords2: [] as ILanguageWord[],
//   }),

//   actions: {
//     async dispatchGetLanguageWords(): Promise<APIResponse<null>> {
//       try {
//         const { status, data } = await API.languageWord.getWords();
//         if (status == 200) {
//           this.languageWords2 = data.data
//           return {
//             success: true,
//             data: null,
//             error: "",
//           };
//         }
//       } catch (error) {
//         const _error = error as AxiosError<string>;
//         return {
//           data: null,
//           error: _error.cause?.message,
//           success: false,
//         } as APIResponse<null>;
//       }
//       return {
//         data: null,
//         error: "Access problem",
//         success: false,
//       };
//     }
//   },
// });
