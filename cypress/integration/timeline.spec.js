describe("Timeline", () => {
    beforeEach(() => {
        cy.setCookie("Other", "other");
        cy.setCookie("XSRF-TOKEN", "abcde");
        cy.visit("/timeline");
    });

    it("displays timeline heading", () => {
        cy.get("h1").should("contain", "Timeline");
    });

    it("displays timeline content", () => {
        cy.get(".title").should("contain", "airportJavaScript");
    });
});