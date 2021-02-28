import faker from 'faker';
import { Namespace } from '../../../../../src/namespace/entity';
import { System } from '../../../../../src/system/entity';

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

export const expandNamespaceNode = (system: System, namespace: Namespace) => {
  cy.getBySel(`${system.name}_${namespace.name}`).click();
}

export const clickRoutesNode = (system: System, namespace: Namespace) => {
  cy.getBySel(`${system.name}_${namespace.name}_routes`).click();
}
