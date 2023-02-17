import CustomBtn5 from "../buttons/CustomBtn5.vue";
import BaseHeader from "../BaseHeader.vue";

describe("BaseHeader Component Tests", () => {
  it("Home Button exist.", () => {
    cy.mount(BaseHeader, {
      components: { CustomBtn5 },
    });
    cy.dataCy("homeBtn");
  });
});
