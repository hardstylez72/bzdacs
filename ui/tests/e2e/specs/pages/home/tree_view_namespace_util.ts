import faker from 'faker';
import { Namespace } from '../../../../../src/views/namespace/service';
import { System } from '../../../../../src/views/system/service';

export const buildNamespace = (): Namespace => ({
  name: faker.lorem.word(),
});

export const fillForm = (namespace: Namespace) => {
  cy.getBySel('name-input').clear();
  cy.getBySel('name-input').type(namespace.name);
};

export const assertNamespaceIsShown = (namespace: Namespace) => {
  cy.getBySel('tree-view').should('contain.text', namespace.name);
};

export const assertNamespaceIsNotShown = (namespace: Namespace) => {
  cy.getBySel('tree-view').should('not.contain.text', namespace.name);
};

export const clickNamespaceOptions = (system: System, namespace: Namespace) => {
  cy.getBySel(`${system.name}_${namespace.name}_namespace_options`).click();
};
