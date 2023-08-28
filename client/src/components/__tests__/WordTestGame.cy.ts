import WordTestGame from "../WordTestGame.vue";
import "../../../cypress/support/component";
import CustomBtn4 from "../buttons/CustomBtn4.vue";
import type { ILanguageWord } from "../../../typings/interface/ILanguageWord";

beforeEach(() => {
  cy.mount(WordTestGame, {
    props: {
      languageData: languageWords,
    },
    components: { CustomBtn4 },
  });
});

describe("Expected Inputs", () => {
  it("'I Don't Know' button should exist", () => {
    cy.dataCy("dontKnowBtn").should("exist");
  });
  it("'Check' button should exist", () => {
    cy.dataCy("checkBtn").should("exist");
  });
  it("'Next' button shouldn't exist", () => {
    cy.dataCy("nextBtn").should("not.exist");
  });
  it("Is the question in the right place", () => {
    cy.dataCy("wordTestQuestion");
    cy.dataCy("wordTestQuestion").contains(languageWords[0].ForeignLanguage);
  });
});
describe("'I Don't Know' Button Tests", () => {
  it("When clicking the button, the button will be disabled", () => {
    cy.dataCy("dontKnowBtn").click();
    cy.dataCy("dontKnowBtn").should("be.disabled");
  });
  it("When clicking the button, we will see answer", () => {
    cy.dataCy("dontKnowBtn").click();
    cy.dataCy("wordTestResult").contains(languageWords[0].MainLanguage);
  });
  it("When clicking the button, we will be see next button", () => {
    cy.dataCy("dontKnowBtn").click();
    cy.dataCy("nextBtn").should("exist");
  });
});
describe("'Check' Button Tests", () => {
  it("If answer input is empty and clicking the 'Check' button, doesn't change button visibility", () => {
    cy.dataCy("checkBtn").should("be.visible");
  });
  it("If answer input is filled and clicking the 'Check' button, does change button visibility", () => {
    cy.dataCy("answerText").type(
      `{selectall}{backspace}${languageWords[0].MainLanguage}`
    );
    cy.dataCy("checkBtn").should("exist");
    cy.dataCy("checkBtn").click().should("not.exist");
  });
  it("When write answer and click the 'Check' button, we should see 'Next' button", () => {
    cy.dataCy("answerText").type(
      `{selectall}{backspace}${languageWords[0].MainLanguage}`
    );
    cy.dataCy("checkBtn").click();
    cy.dataCy("nextBtn").should("exist");
  });
});
describe("'Next' Button Tests",()=>{
  beforeEach(() => {
    cy.dataCy("answerText").type(
      `{selectall}{backspace}${languageWords[0].MainLanguage}`
    );
    cy.dataCy("checkBtn").click();
    cy.dataCy("nextBtn").should("exist");
  });
  it("When clicking the 'Next' button, 'Check' and 'I Don't Know' buttons should be see",()=>{
    cy.dataCy("nextBtn").click();
    cy.dataCy("checkBtn").should("exist");
    cy.dataCy("dontKnowBtn").should("exist");
  })
  it("When clicking the 'Next' button, 'Next' buttons shouldn't be see",()=>{
    cy.dataCy("nextBtn").click();
    cy.dataCy("nextBtn").should("not.exist");
  })
  it("When clicking the 'Next' button, 'I Don't Know' buttons shouldn't be disabled",()=>{
    cy.dataCy("nextBtn").click();
    cy.dataCy("dontKnowBtn").should("not.be.disabled");
  })
  it("When clicking the 'Next' button, textbox should be empty",()=>{
    cy.dataCy("nextBtn").click();
    cy.dataCy("answerText").should("have.value","");
  })
  it("When clicking the 'Next' button, Word Test Result Part shouldn't be see",()=>{
    cy.dataCy("nextBtn").click();
    cy.dataCy("wordTestResult").should("not.be.visible");
  })
})
describe("Word Test Result Part", () => {
  it("Word test result is visible", () => {
    cy.dataCy("wordTestResult").should("not.be.visible");
  });
  it("When clicking 'I don't know' button, word test result part is visible", () => {
    cy.dataCy("dontKnowBtn").click();
    cy.dataCy("wordTestResult").should("be.visible");
  });
  it("If the correct word is written , word test result part will show correct answer class", () => {
    cy.dataCy("answerText").type(
      `{selectall}{backspace}${languageWords[0].MainLanguage}`
    );
    cy.dataCy("checkBtn").click();
    cy.dataCy("wordTestResult").should("have.class", "wordTestResultCorrect");
  });
});

const languageWords = [
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
  {
    id: "3",
    MainLanguage: "Yazi3",
    ForeignLanguage: "Text3",
  },
  {
    id: "4",
    MainLanguage: "Yazi4",
    ForeignLanguage: "Text4",
  },
  {
    id: "5",
    MainLanguage: "Yazi5",
    ForeignLanguage: "Text5",
  },
] as ILanguageWord[];