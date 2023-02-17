import CustomBtn5 from "../buttons/CustomBtn5.vue";

describe("CustomBtn5 Component Test", () => {
  it("Button name change", () => {
    const btnName = "Test Name";
    cy.mount(CustomBtn5, {
      props: {
        btnName: btnName,
      },
    });
    cy.dataCy("btn5").should("have.text", btnName);
  });
});
