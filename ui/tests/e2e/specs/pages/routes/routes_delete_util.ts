import { RouteSwg as Route } from '../../../../../src/route/entity';

const clickRouteDelete = (route: Route) => {
  cy.contains(route.route)
    .parents()
    .find('[data-test=route-delete-action]')
    .click()
};

const clickSystemDeleteOption = () => {
  cy.getBySel('system-delete-option').click();
};

const clickDeleteButton = () => {
  cy.getBySel('delete-btn').click();
};

export const deleteRoute = (route: Route) => {
  clickRouteDelete(route);
  clickSystemDeleteOption();
  clickDeleteButton();
};
