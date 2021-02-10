Cypress.Commands.add('login', () => {
  cy.setCookie(Cypress.env().cookieName, Cypress.env().adminSession);
});

Cypress.Commands.overwrite('visit', (originalFn, url, options) => {
  cy.login();
  originalFn(url, options);
});

Cypress.Commands.add('getBySel', (sel) => {
  cy.get(`[data-test=${sel}]`);
});
