/// <reference types="Cypress" />

// @ts-ignore
const user = 'test-user';

describe('Guest page', () => {
  it('login success', () => {
    cy.clearCookies()
      .visit('/guest/?lang=en')
      .wait(100);
    cy.get('[data-test=login-input]')
      .type(user);
    cy.get('[data-test=login-btn]')
      .click();
    cy.get('[data-test=avatar]')
      .should('contain', user);
  });

  it('login empty', () => {
    cy.clearCookies()
      .visit('/guest/?lang=en')
      .wait(100);
    cy.get('[data-test=login-btn]')
      .click();
    cy.get('.v-form')
      .should('contain', 'required');
  });
});
