import CustomBtn4 from "../buttons/CustomBtn4.vue";

describe("CustomBtn4 Component Test", () => {
  it("Button name change", () => {
    const btnName = "Test Name";
    cy.mount(CustomBtn4, {
      props: {
        btnName: btnName,
      },
    });
    cy.dataCy("btn4").should("have.text", btnName);
  });
});
