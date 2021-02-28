import { RouteSwg as Route }  from '../../../../../src/route/entity';
import { assertSystemIsShown, buildRoute, fillForm } from './routes_util';
import {Namespace} from "../../../../../src/namespace/entity";

const clickNewRouteButton = () => {
  cy.getBySel('add-route-btn').click();
};

const clickSaveRouteButton = () => {
  cy.getBySel('save-route-btn').click();
};

const createRoute = (namespace: Namespace, route?: Route): Route => {
  if (!route) {
    // eslint-disable-next-line no-param-reassign
    route = buildRoute();
  }
  clickNewRouteButton();
  fillForm(route);
  clickSaveRouteButton();
  cy.get('.v-dialog').should('not.be.visible');
  // @ts-ignore
  return route;
};

// eslint-disable-next-line import/prefer-default-export
export const createValidRoute = (namespace: Namespace, route?: Route): Route => {
  const created = createRoute(route);
  assertSystemIsShown(created);
  return created;
};
