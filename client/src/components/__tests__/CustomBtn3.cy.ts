import CustomBtn3 from "../buttons/CustomBtn3.vue";

describe("CustomBtn3 Component Test", () => {
  it("Button name change", () => {
    const btnName = "Test Name";
    cy.mount(CustomBtn3, {
      props: {
        btnName: btnName,
      },
    });
    cy.dataCy("btn3").should("have.text", btnName);
  });
});
