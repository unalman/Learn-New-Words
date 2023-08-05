/**
 * @vitest-environment happy-dom
 */
import { describe, expect, it } from "vitest";
import { mount } from "@vue/test-utils";
import { getLastId, sortLanguageWordDescending } from "../../libs/WordsTable";
import WordsTable from "../../components/WordsTable.vue";
import type { ILanguageWord } from "typings/interface/ILanguageWord";

describe("WordsTable rendering", () => {
  it("render grid correctly", async () => {
    const wrapper = mount(WordsTable, {
      props: { languageData: models },
    });
    expect(
      wrapper.find("[data-testId='WordsTable-TestId']").exists()
    ).toBeTruthy();
  });
  it("render grid correctly", async () => {
    const wrapper = mount(WordsTable, {
      props: { languageData: models },
    });
    expect(wrapper.text()).toMatch(model.ForeignLanguage);
  });
  it("when send empty array, you should see no result message", async () => {
    const wrapper = mount(WordsTable, {
      props: { languageData: [] },
    });
    expect(wrapper.find("[data-testId='noresult']").exists()).toBeTruthy();
  });
});
describe("WordTable props", () => {
  it("when ILanguageWord array length 0, function should be return '1'", () => {
    expect(getLastId([])).toBe("1");
  });
  it("when ILanguageWord array has one ILanguageWord object, function should be return '2'", () => {
    expect(getLastId(models)).toBe("2");
  });
});
describe("sortLanguageWordDescending tests", () => {
  it("ILanguageWord array sorted descending", () => {
    var descendingIds = ["4", "3", "2", "1"];
    expect(sortLanguageWordDescending(sortList).map((x) => x.id)).toEqual(
      descendingIds
    );
  });
});

const models: Array<ILanguageWord> = [];
const model = {
  id: "1",
  ForeignLanguage: "Test Data",
  MainLanguage: "Test Data 2",
} as ILanguageWord;
models.push(model);

var sortList: Array<ILanguageWord> = [
  {
    id: "4",
    MainLanguage: "Yazi4",
    ForeignLanguage: "Text4",
  },
  {
    id: "2",
    MainLanguage: "Yazi2",
    ForeignLanguage: "Text2",
  },
  {
    id: "3",
    MainLanguage: "Yazi3",
    ForeignLanguage: "Text3",
  },
  {
    id: "1",
    MainLanguage: "Yazi1",
    ForeignLanguage: "Text1",
  },
];
