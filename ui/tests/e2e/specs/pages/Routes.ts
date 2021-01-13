/// <reference types="Cypress" />

// @ts-ignore
const description = 'test description';
const route = '/test/ggg213322334s';
const method = 'PUT';

describe('Routes', () => {
  it('create filled', () => {
    cy.clearCookies()
      .setCookie(Cypress.env().cookieName, Cypress.env().adminSession)
      .visit('/?lang=en');

    cy.get('[data-test=add-route-btn]').click();
    cy.get('[data-test=form-route-input]').type(route, { force: true });

    cy.get('.col-sm-2 > .v-input').click();
    cy.get('.v-list-item__content > .v-list-item__title').contains(method).click();

    cy.get('[data-test=form-description-input]').type(description, { force: true });
    cy.get('[data-test=save-route-btn]').click();
    cy.get('tbody > tr').contains(description);
    cy.get('tbody > tr').contains(route);
    cy.get('tbody > tr').contains(method);
  });
});
