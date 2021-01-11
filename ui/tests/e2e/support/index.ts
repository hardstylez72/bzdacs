// cypress/support/index.ts
Cypress.Commands.add('dataCy', (value) => cy.get(`[data-cy=${value}]`));
