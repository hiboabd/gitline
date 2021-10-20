describe("Homepage", () => {
    beforeEach(() => {
        cy.setCookie("Other", "other");
        cy.setCookie("XSRF-TOKEN", "abcde");
        cy.visit("/");
    });

    it("displays gitline heading", () => {
        cy.get("h1").should("contain", "Gitline");
    });
});