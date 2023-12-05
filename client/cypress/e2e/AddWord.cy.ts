describe("AddWord Page Tests", () => {
  beforeEach(() => {
    cy.visit("/addword");
  });
  it("Open the Add Word page", () => {});
  it("check the existence of component", () => {
    cy.dataTest("WordsTable-TestId").should("exist");
  });
  it("Add button should exist", () => {
    cy.dataCy("addbtn").should("exist");
  });
  it("When clicking the Add button,create a new word field in the grid", () => {
    cy.dataCy("addbtn").click();
    cy.dataCy("mainLangEdit").should("be.visible");
    cy.dataCy("foreignLangEdit").should("be.visible");
  });
  it("When clicking the Edit button, edit mode should be turn on", () => {
    cy.dataCy("editbtn").eq(0).click();
    cy.dataCy("mainLangEdit").should("be.visible");
    cy.dataCy("foreignLangEdit").should("be.visible");
    cy.dataCy("okbtn").should("be.visible");
    cy.dataCy("cancelbtn").should("be.visible");
  });
  it("When clicking the Ok button in Edit mode, edit mode should be turn off", () => {
    cy.dataCy("editbtn").eq(0).click();
    cy.dataCy("okbtn").eq(0).click();
    cy.dataCy("mainLangEdit").should("not.exist");
    cy.dataCy("foreignLangEdit").should("not.exist");
    cy.dataCy("okbtn").should("not.exist");
    cy.dataCy("cancelbtn").should("not.exist");
  });
  it("When clicking the Cancel button in Edit mode, edit mode should be turn off", () => {
    cy.dataCy("editbtn").eq(0).click();
    cy.dataCy("cancelbtn").eq(0).click();
    cy.dataCy("mainLangEdit").should("not.exist");
    cy.dataCy("foreignLangEdit").should("not.exist");
    cy.dataCy("okbtn").should("not.exist");
    cy.dataCy("cancelbtn").should("not.exist");
  });
});
