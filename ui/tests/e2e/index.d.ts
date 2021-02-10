// in cypress/support/index.d.ts
// load type definitions that come with Cypress module
/// <reference types="cypress" />

declare namespace Cypress {
  interface Chainable {
    /**
     * Custom command to select DOM element by data-test attribute.
     * @example cy.dataCy('greeting')
     */
    getBySel(value: string): Chainable<Element>;
    login(): Chainable<Element>;
  }
}
