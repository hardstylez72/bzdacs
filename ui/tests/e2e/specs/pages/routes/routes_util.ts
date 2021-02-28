import faker from 'faker';
import { RouteSwg as Route } from '../../../../../src/route/entity';

export const buildRoute = (): Route => ({
  route: '/' + faker.lorem.word(),
  description: faker.lorem.word(),
  method: 'GET'
});

export const fillForm = (route: Route) => {
  cy.getBySel('form-route-input').clear();
  cy.getBySel('form-route-input').type(route.route);

  cy.getBySel('form-description-input').clear();
  cy.getBySel('form-description-input').type(route.description);

  cy.getBySel('form-method-input').click();
  cy.get('.v-list-item__content > .v-list-item__title').contains(route.method).click();
};

export const assertSystemIsShown = (route: Route) => {
  cy.get('tbody > tr').should('contain.text', route.route);
};

export const assertSystemIsNotShown = (route: Route) => {
  cy.get('tree-view').should('not.contain.text', route.route);
};

