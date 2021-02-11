import faker from 'faker';
import { System } from '../../../../../src/views/system/service';

export const buildSystem = (): System => ({
  id: -1,
  name: faker.lorem.word(),
});

export const fillForm = (system: System) => {
  cy.getBySel('name-input').clear();
  cy.getBySel('name-input').type(system.name);
};

export const assertSystemIsShown = (system: System) => {
  cy.getBySel('tree-view').should('contain.text', system.name);
};

export const assertSystemIsNotShown = (system: System) => {
  cy.getBySel('tree-view').should('not.contain.text', system.name);
};
