describe("Timeline", () => {
    beforeEach(() => {
        cy.setCookie("Other", "other");
        cy.setCookie("XSRF-TOKEN", "abcde");
    });

    it("displays timeline content when username submitted", () => {
        cy.visit("/");
        cy.get("#username").type("Github Username");
        cy.get("#submit-username-form").submit();
        cy.get("h1").should("contain", "Timeline");
        cy.get(".title").should("contain", "gitline");
    });
});