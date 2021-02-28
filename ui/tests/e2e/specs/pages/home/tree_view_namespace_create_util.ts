import { assertNamespaceIsShown, buildNamespace, fillForm } from './tree_view_namespace_util';
import { Namespace } from '../../../../../src/namespace/entity';
import { System } from '../../../../../src/system/entity';

const openSystem = (system: System) => {
  cy.getBySel(system.name).click();
};

const clickCreateNamespaceButton = (system: System) => {
  cy.getBySel(`${system.name}_createNewNamespaceBtn`).click();
};

const clickSaveSystemButton = () => {
  cy.getBySel('save-namespace-btn').click();
};

export const createNamespace = (system: System, namespace?: Namespace): Namespace => {
  if (!namespace) {
    // eslint-disable-next-line no-param-reassign
    namespace = buildNamespace();
  }

  openSystem(system);
  clickCreateNamespaceButton(system);
  fillForm(namespace);
  clickSaveSystemButton();

  cy.getBySel('create-system-dialog-card').should('not.exist');
  // @ts-ignore
  return namespace;
};

export const createValidNamespace = (system: System, namespace?: Namespace): Namespace => {
  const ns = createNamespace(system, namespace);
  assertNamespaceIsShown(ns);
  return ns;
};
