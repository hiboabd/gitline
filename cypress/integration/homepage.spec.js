describe("Homepage", () => {
    beforeEach(() => {
        cy.setCookie("Other", "other");
        cy.setCookie("XSRF-TOKEN", "abcde");
        cy.visit("/");
    });

    it("displays hello world message", () => {
        cy.get("h1").should("contain", "Hello World!");
    });
});