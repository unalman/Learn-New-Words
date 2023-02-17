import CustomBtn2 from "../buttons/CustomBtn2.vue";

describe("CustomBtn2 Component Test", () => {
  it("Button name change", () => {
    const btnName = "Test Name";
    cy.mount(CustomBtn2, {
      props: {
        btnName: btnName,
      },
    });
    cy.dataCy("btn2").should("have.text", btnName);
  });
});
