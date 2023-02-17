import ErrorBlock from "../ErrorBlock.vue";
import "../../../cypress/support/component";

describe("<ErrorBlock />", () => {
  it("renders", () => {
    cy.mount(ErrorBlock, {
      props: {
        errors: errors,
      },
    });
  });
});

const errors = ["Error 1", "Error 2"];
