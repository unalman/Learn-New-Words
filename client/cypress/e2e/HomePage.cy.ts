describe("Home Page Content Tests", () => {
  beforeEach(() => {
    cy.visit("/");
  });
  it("Open the home page", () => {});
  it("check the buttons existence and links", () => {
    cy.dataCy("addWord")
      .should("exist")
      .and("have.prop", "href")
      .and("contain", "/addword");
    cy.dataCy("wordTest")
      .should("exist")
      .and("have.prop", "href")
      .and("contain", "/wordtest");
  });
  context("Content Side Direct Tests", () => {
    it("When clicking the addWord button, it should direct to Add Word page", () => {
      cy.dataCy("addWord").click();
      cy.location("pathname").should("equal", "/addword");
    });
    it("When clicking the wordTest button, it should direct to Word Test page", () => {
      cy.dataCy("wordTest").click();
      cy.location("pathname").should("equal", "/wordtest");
    });
  })
});
describe("Header Tests", () => {
  beforeEach(() => {
    cy.visit("/");
  });
  it("check the buttons existence and links", () => {
    cy.dataCy("addWord")
      .should("exist")
      .and("have.prop", "href")
      .and("contain", "/addword");
    cy.dataCy("wordTest")
      .should("exist")
      .and("have.prop", "href")
      .and("contain", "/wordtest");
  });
  context("Header Direct Tests", () => {
    it("When clicking the homeBtn button, it should direct to Home page page", () => {
      cy.dataCy("homeBtn").click();
      cy.location("pathname").should("equal", "/");
    });
    it("When clicking the wordTest button, it should direct to WordTest page", () => {
      cy.dataCy("contactBtn").click();
      cy.location("pathname").should("equal", "/contact");
    });
  })
  context("Localization Tests On Header", () => {
    it("Localization selectbox should exist", () => {
      cy.dataCy("locales-select").should("exist");
    });
    it("When select TR language", () => {
      cy.dataCy("locales-select").select("TR");
      cy.dataCy("homeBtn").should("contain","Anasayfa")
    });
    it("When select EN language", () => {
      cy.dataCy("locales-select").select("EN");
      cy.dataCy("homeBtn").should("contain","Home")
    });
  });
});