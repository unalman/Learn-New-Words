import WordsTable from "../WordsTable.vue";
import "../../../cypress/support/component";
import CustomBtn2 from "../buttons/CustomBtn2.vue";
import CustomBtn3 from "../buttons/CustomBtn3.vue";
import ErrorBlock from "../ErrorBlock.vue";
import type { ILanguageWord } from "../../../typings/interface/ILanguageWord";


beforeEach(() => {
  cy.mount(WordsTable, {
    props: {
      languageData: [
        {
          id: "1",
          MainLanguage: "Yazi1",
          ForeignLanguage: "Text1",
        },
        {
          id: "2",
          MainLanguage: "Yazi2",
          ForeignLanguage: "Text2",
        },
      ] as ILanguageWord[],
    },
    components: { CustomBtn2, CustomBtn3, ErrorBlock },
  });
});

describe("Empty Data Tests", () => {
  it("If send empty array, it will shows about empty array text", () => {
    cy.mount(WordsTable, {
      props: {
        languageData: [],
      },
      components: { CustomBtn2, CustomBtn3, ErrorBlock },
    });
    cy.get(".word-list-noresult").should("exist");
  });
});
describe("Add Button Tests", () => {
  it("'Add' button should exist", () => {
    cy.dataCy("addbtn").should("exist");
  });
  it("A new object will be created, when the Add button is clicked and second click will not effect li.", () => {
    cy.dataCy("languageItem").should("have.length", 2);
    cy.dataCy("addbtn").click();
    cy.dataCy("languageItem").should("have.length", 3);
    cy.dataCy("addbtn").click();
    cy.dataCy("languageItem").should("have.length", 3);
  });
});
describe("Remove Button Tests", () => {
  it("Clicking the Delete button will delete li", () => {
    cy.dataCy("languageItem").should("have.length", 2);
    cy.dataCy("deletebtn").first().click();
    cy.dataCy("languageItem").should("have.length", 1);
  });
});
describe("Error Component Tests", () => {
  it("Error component is visible", () => {
    cy.dataCy("addbtn").click();
    cy.dataCy("okbtn").first().click();
    cy.get(".errorBlock-container").should("exist");
  });
});
describe("Cancel Button Tests", () => {
  it("When you click the cancel button in the newly added field, the field will be deleted", () => {
    cy.dataCy("addbtn").click();
    cy.dataCy("cancelbtn").first().click();
  });
  it("When you click the cancel button in the uptaded field, the field won't be update", () => {
    const datas = [
      {
        id: "1",
        MainLanguage: "Yazi1",
        ForeignLanguage: "Text1",
      },
    ] as ILanguageWord[];
    cy.mount(WordsTable, {
      props: {
        languageData: datas,
      },
      components: { CustomBtn2, CustomBtn3, ErrorBlock },
    });
    cy.dataCy("editbtn").first().click();
    cy.dataCy("mainLangEdit").type("test");
    cy.dataCy("foreignLangEdit").type("test");
    cy.dataCy("cancelbtn").first().click();
    cy.dataCy("mainLang").should("have.text", datas[0].MainLanguage);
    cy.dataCy("foreignLang").should("have.text", datas[0].ForeignLanguage);
  });
});
describe("Edit Button Tests", () => {
  it("Clicking the Edit button will open the Edit mode", () => {
    cy.dataCy("editbtn").first().click();
    cy.dataCy("mainLangEdit").should("exist");
    cy.dataCy("foreignLangEdit").should("exist");
  });
});
describe("Update Button Tests", () => {
  it("When Edit mode is on, data values can update", () => {
    const datas = [
      {
        id: "1",
        MainLanguage: "Yazi1",
        ForeignLanguage: "Text1",
      },
    ] as ILanguageWord[];
    cy.mount(WordsTable, {
      props: {
        languageData: datas,
      },
      components: { CustomBtn2, CustomBtn3, ErrorBlock },
    });
    const mainValue = "test1";
    const foreignValue = "test2";

    cy.dataCy("editbtn").click();
    cy.dataCy("mainLangEdit").type(`{selectall}{backspace}${mainValue}`);
    cy.dataCy("foreignLangEdit").type(`{selectall}{backspace}${foreignValue}`);
    cy.dataCy("okbtn").click();
    cy.dataCy("mainLang").should("have.text", mainValue);
    cy.dataCy("foreignLang").should("have.text", foreignValue);
  });
  it("If the text is empty in edit mode, it won't be updated", () => {
    const datas = [
      {
        id: "1",
        MainLanguage: "Yazi1",
        ForeignLanguage: "Text1",
      },
    ] as ILanguageWord[];
    cy.mount(WordsTable, {
      props: {
        languageData: datas,
      },
      components: { CustomBtn2, CustomBtn3, ErrorBlock },
    });
    cy.dataCy("editbtn").click();
    cy.dataCy("mainLangEdit").type(`{selectall}{backspace}`);
    cy.dataCy("foreignLangEdit").type(`{selectall}{backspace}`);
    cy.dataCy("okbtn").click();
    cy.dataCy("mainLangEdit").should("exist");
    cy.dataCy("foreignLangEdit").should("exist");
  });
});
