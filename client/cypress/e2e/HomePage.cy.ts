describe("Home Page Tests", () => {
  it("Open the home page and ", () => {
    cy.visit("/");
  });
  it("check the button links", () => {
    cy.get(".customBtn1[href='/addword']");
    cy.get(".customBtn1[href='/wordtest']");
  });
});
