import CustomBtn1 from "../buttons/CustomBtn1.vue";

describe("CustomBtn1 Component Test", () => {
  it("Button name change", () => {
    const btnName = "Test Name";
    cy.mount(CustomBtn1, {
      props: {
        btnName: btnName,
      },
    });
    cy.dataCy("btn1").should("have.text", btnName);
  });
});
