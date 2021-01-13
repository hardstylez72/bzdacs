/// <reference types="Cypress" />
// @ts-ignore
const user = 'test-user';

describe('Login - logout as Guest', () => {
  it('login-logout', () => {
    cy.clearCookies()
      .visit('/guest/?lang=en')
      .wait(100);
    cy.get('[data-test=login-input]')
      .type(user);
    cy.get('[data-test=login-btn]')
      .click();
    cy.get('[data-test=avatar]')
      .should('contain', user);
    cy.get('[data-test=logout]')
      .click();
    cy.get('[data-test=login-input]');
  });
});
