describe("AddWord Page Tests", () => {
  it("Open the AddWord page", () => {
    cy.visit("/addword");
  });
  it("check the existence of component", () => {
    cy.dataCy("addbtn");
    cy.get(".word-container");
  });
});
